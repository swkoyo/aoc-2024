package d10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func (c *Coord) ToString() string {
	return fmt.Sprintf("(%d,%d)", c.row, c.col)
}

func UniqueTraverse(data [][]int, prev *Coord, curr Coord, path []Coord, set map[string]bool) {
	if curr.row < 0 || curr.row >= len(data) {
		return
	}

	if curr.col < 0 || curr.col >= len(data[0]) {
		return
	}

	currLevel := data[curr.row][curr.col]

	if prev != nil {
		prevLevel := data[prev.row][prev.col]

		if currLevel-1 != prevLevel {
			return
		}

	}

	path = append(path, curr)

	if currLevel == 9 {
		pathStr := ""
		for _, coord := range path {
			pathStr += coord.ToString()
		}
		set[pathStr] = true
		path = path[:len(path)-1]
		return
	}

	// UP
	UniqueTraverse(data, &curr, Coord{row: curr.row - 1, col: curr.col}, path, set)
	// DOWN
	UniqueTraverse(data, &curr, Coord{row: curr.row + 1, col: curr.col}, path, set)
	// LEFT
	UniqueTraverse(data, &curr, Coord{row: curr.row, col: curr.col - 1}, path, set)
	// RIGHT
	UniqueTraverse(data, &curr, Coord{row: curr.row, col: curr.col + 1}, path, set)

	path = path[:len(path)-1]
}

func P2() {
	file, err := os.Open("d10/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	data := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currRow := make([]int, 0)
		for _, char := range scanner.Text() {
			if unicode.IsNumber(char) {
				num, err := strconv.Atoi(string(char))
				if err != nil {
					log.Fatal(err)
				}
				currRow = append(currRow, num)
			} else {
				currRow = append(currRow, -1)
			}
		}
		data = append(data, currRow)
	}

	set := make(map[string]bool)
	for r, row := range data {
		for c, val := range row {
			if val == 0 {
				path := make([]Coord, 0)
				UniqueTraverse(data, nil, Coord{row: r, col: c}, path, set)
			}
		}
	}
	total += len(set)

	fmt.Println("D10 P2: ", total)
}
