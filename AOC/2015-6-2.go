package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // Make sure to close the file at the end
	scanner := bufio.NewScanner(file)

	var array [1000][1000]int

	//count := 0
	for scanner.Scan() {
		line := scanner.Text() // Read the current line
		sline := strings.Split(line, " ")
		stype := parse_type(sline)
		sbeg := parse_point_begin(sline)
		send := parse_point_end(sline)
		//fmt.Println(sline)
		rw(sbeg, send, &array, stype)
	}
	fmt.Println(rwCount(&array))
}

func parse_type(line []string) string {
	if line[0] == "turn" {
		//fmt.Println(line[0] + " " + line[1])
		return line[0] + " " + line[1]
	}
	if line[0] == "toggle" {
		return line[0]
	}
	return "Error"
}

func parse_point_begin(line []string) Point {
	if parse_type(line) == "turn on" || parse_type(line) == "turn off" {
		locs := strings.Split(line[2], ",")
		value1, _ := strconv.Atoi(locs[0])
		value2, _ := strconv.Atoi(locs[1])
		return Point{value1, value2}
	} else {
		locs := strings.Split(line[1], ",")
		value1, _ := strconv.Atoi(locs[0])
		value2, _ := strconv.Atoi(locs[1])
		return Point{value1, value2}
	}
}

func parse_point_end(line []string) Point {
	if parse_type(line) == "turn on" || parse_type(line) == "turn off" {
		locs := strings.Split(line[4], ",")
		value1, _ := strconv.Atoi(locs[0])
		value2, _ := strconv.Atoi(locs[1])
		return Point{value1, value2}
	} else if parse_type(line) == "toggle" {
		locs := strings.Split(line[3], ",")
		value1, _ := strconv.Atoi(locs[0])
		value2, _ := strconv.Atoi(locs[1])
		return Point{value1, value2}
	}
	fmt.Println("ERROR")
	return Point{0, 0}
}
func rw(begin Point, end Point, array *[1000][1000]int, t string) {

	for i := begin.X; i <= end.X; i++ {
		for j := begin.Y; j <= end.Y; j++ {

			if t == "turn on" {
				array[i][j]++
			} else if t == "turn off" && array[i][j] != 0 {
				array[i][j]--
			} else if t == "toggle" {

				array[i][j] += 2

			}
		}
	}
}

func rwCount(array *[1000][1000]int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			count += array[i][j]
		}
	}
	return count
}
