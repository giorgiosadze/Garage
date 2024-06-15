package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("uploads", "upload-*.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully uploaded file\n")
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	fmt.Println("Server started on :8080")
	os.MkdirAll("uploads", os.ModePerm)
	http.ListenAndServe(":8080", nil)
}
