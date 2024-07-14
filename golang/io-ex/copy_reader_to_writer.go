package main

import (
	"io"
	"strings"
	"os"
)

func main() {
	reader := strings.NewReader("Hello, Go!")
	io.Copy(os.Stdout, reader)
}