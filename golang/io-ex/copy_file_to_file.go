package main

import (
	"io"
	"os"
)

func main() {
	src, _ := os.Open("source.txt")
	dst, _ := os.Create("destination.txt")
	defer src.Close()
	defer dst.Close()
	io.Copy(dst, src)
}
