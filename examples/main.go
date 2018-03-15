package main

import (
	"github.com/artdarek/go-unzip"
	"fmt"
)

func main() {
	uz := unzip.New("file.zip", "directory/")
	err := uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}