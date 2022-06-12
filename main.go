package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"

	"github.com/batt0s/go-ascii/asciicon"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Specify a file.")
		os.Exit(1)
	}
	img := loadFile(args[0])
	asciicon.Convert(img)
}

func loadFile(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ERROR : %v", err)
		os.Exit(1)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("ERROR : %v", err)
		os.Exit(1)
	}
	return img
}
