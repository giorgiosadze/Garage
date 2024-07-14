package main

import (
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("file.txt")
	defer file.Close()
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
