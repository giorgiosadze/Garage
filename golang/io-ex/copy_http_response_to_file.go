package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("https://osadze.com")
	defer resp.Body.Close()
	file, _ := os.Create("response.html")
	defer file.Close()
	io.Copy(file, resp.Body)
}