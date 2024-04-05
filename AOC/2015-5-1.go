package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // Make sure to close the file at the end
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text() // Read the current line

		if strings.Contains(line, "ab") ||
			strings.Contains(line, "cd") ||
			strings.Contains(line, "pq") ||
			strings.Contains(line, "xy") {
			continue
		} else {
			if CountContaining(line, "aeiou") >= 3 && DupInRow(line) {
				count++
			}
		}

	}
	fmt.Println(count)
}

// Count how many times runes in sub appear in str
func CountContaining(str string, sub string) int {

	count := 0
	for _, stack := range str {
		for _, target := range sub {
			if stack == target {
				count++
			}
		}
	}

	return count
}

// True if any rune in a string is a duplicate of the next rune
func DupInRow(str string) bool {

	for i, _ := range str {
		if len(str) > 1+i && str[i] == str[i+1] {
			return true
		}
	}
	return false
}
