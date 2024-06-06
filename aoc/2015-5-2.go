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

	defer file.Close() 
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text() // Read the current line

		if FindPair(line) &&
			OneLetterRepeat(line) {

			count++
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

// True if any rune in a string is a duplicate of the next rune
func FindPair(str string) bool {

	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		if len(str) > 1+i {

			key := string(str[i]) + string(str[i+1])
			m[key]++
			if m[key] == 2 && strings.Count(str, key) == 2 {

				fmt.Println(str, key)
				return true
			}

		}
	}
	return false
}

// True if any rune in a string is a duplicate of the next rune
func OneLetterRepeat(str string) bool {

	for i := 0; i < len(str); i++ {
		if len(str) > 2+i && str[i] == str[i+2] {
			return true
		}
	}
	return false
}
