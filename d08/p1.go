package d08

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Location struct {
	row int
	col int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}

func GetDistance(x Location, y Location) (int, int) {
	rowDiff := x.row - y.row
	colDiff := x.col - y.col
	return Abs(rowDiff), Abs(colDiff)
}

func GetAntinodes(x Location, y Location) (Location, Location) {
	rowDiff, colDiff := GetDistance(x, y)
	var leftAntinode Location
	var rightAntinode Location
	if x.col <= y.col {
		leftAntinode = Location{row: x.row - rowDiff, col: x.col - colDiff}
		rightAntinode = Location{row: y.row + rowDiff, col: y.col + colDiff}
	} else {
		leftAntinode = Location{row: x.row - rowDiff, col: x.col + colDiff}
		rightAntinode = Location{row: y.row + rowDiff, col: y.col - colDiff}
	}
	return leftAntinode, rightAntinode
}

func IsInBounds(rowLen int, colLen int, location Location) bool {
	return location.row >= 0 && location.row < rowLen && location.col >= 0 && location.col < colLen
}

func P1() {
	file, err := os.Open("d08/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	rowLen := 0
	colLen := 0
	nodeMap := make(map[rune][]Location)
	antinodeSet := make(map[Location]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colLen = len(line)
		for col, char := range line {
			if IsAlphaNumeric(char) {
				if _, exists := nodeMap[char]; !exists {
					nodeMap[char] = make([]Location, 0)
				}
				nodeMap[char] = append(nodeMap[char], Location{row: rowLen, col: col})
			}
		}
		rowLen += 1
	}

	for _, nodes := range nodeMap {
		for i := 0; i < len(nodes)-1; i++ {
			currNode := nodes[i]
			for j := i + 1; j < len(nodes); j++ {
				nextNode := nodes[j]
				lAntinode, rAntinode := GetAntinodes(currNode, nextNode)
				if IsInBounds(rowLen, colLen, lAntinode) {
					antinodeSet[lAntinode] = true
				}
				if IsInBounds(rowLen, colLen, rAntinode) {
					antinodeSet[rAntinode] = true
				}
			}
		}
	}

	total = len(antinodeSet)

	fmt.Println("D08 P1: ", total)
}
