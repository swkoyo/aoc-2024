package d10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Coord struct {
	row int
	col int
}

func Traverse(data [][]int, prev *Coord, curr Coord, set map[Coord]bool) {
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

	if currLevel == 9 {
		set[curr] = true
		return
	}

	// UP
	Traverse(data, &curr, Coord{row: curr.row - 1, col: curr.col}, set)
	// DOWN
	Traverse(data, &curr, Coord{row: curr.row + 1, col: curr.col}, set)
	// LEFT
	Traverse(data, &curr, Coord{row: curr.row, col: curr.col - 1}, set)
	// RIGHT
	Traverse(data, &curr, Coord{row: curr.row, col: curr.col + 1}, set)
}

func P1() {
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

	for r, row := range data {
		for c, val := range row {
			if val == 0 {
				set := make(map[Coord]bool)
				Traverse(data, nil, Coord{row: r, col: c}, set)
				total += len(set)
			}
		}
	}

	fmt.Println("D10 P1: ", total)
}
