package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	data := []byte("Data to write to file")
	buffer := bytes.NewBuffer(data)
	file, _ := os.Create("buffer.txt")
	defer file.Close()
	io.Copy(file, buffer)
}
