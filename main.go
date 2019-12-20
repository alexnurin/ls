package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	d    = flag.String("d", ".", "Directory to process")
	a    = flag.Bool("a", false, "Print all info")
	h    = flag.Bool("h", false, "Visual size info")
	sort = flag.Bool("sort", false, "Sorted by date")
)

func bsort(arr *[]os.FileInfo) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j].Size() > (*arr)[j+1].Size() {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
func hrSize(fsize int64) string {
	if *h {
		szs := [6]string{"B", "KB", "MB", "GB", "TB", "PB"}
		id := 0
		for fsize >= 1024 {
			id++
			fsize = (fsize-1)/1024 + 1
		}
		return fmt.Sprintf("%d%s", fsize, szs[id])
	} else {
		return fmt.Sprintf("%d", fsize)
	}
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 01 15:4")
	fmt.Printf("%s %s %s \n", hrSize(file.Size()), time, file.Name())
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)
	if *sort {
		bsort(&files)
	}
	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
