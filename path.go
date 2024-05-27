package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	for _, e := range strings.Split(os.Getenv("PATH"), ":") {
		entries, err := os.ReadDir(e)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, entrie := range entries {
			fmt.Printf("entrie.Name(): %v\n", entrie.Name())
		}
	}
}
