package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "Print all info")
)

func hrSize(fsize int64) string {
	szs := [5]string{"B", "KB", "MB", "GB", "TB"}
	id := 0
	for fsize > 1024 {
		id++
		fsize = fsize / 1024
	}
	return fmt.Sprintf("%d%s", fsize, szs[id])
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 01 15:4")
	fmt.Printf("%s %s %s \n", hrSize(file.Size()), time, file.Name())
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
