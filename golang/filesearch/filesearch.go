package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "./"
	search := "filename.txt"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.Name() == search {
			fmt.Println("Found:", path)
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
