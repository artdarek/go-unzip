package main

import (
	"fmt"

	"github.com/artdarek/go-unzip/pkg/unzip"
)

func main() {
	uz := unzip.New()

	files, err := uz.Extract("./data/file.zip", "./data/directory")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("extracted files count: %d", len(files))
	fmt.Printf("files list: %v", files)
}
