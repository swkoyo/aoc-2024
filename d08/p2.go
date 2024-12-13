package d08

import (
	"bufio"
	"fmt"
	"os"
)

func GetAntinodesList(rowLen, colLen int, x Location, y Location) []Location {
	res := make([]Location, 0)
	rowDiff, colDiff := GetDistance(x, y)
	if x.col <= y.col {
		// UPLEFT
		currRow := x.row - rowDiff
		currCol := x.col - colDiff
		for currRow >= 0 && currCol >= 0 {
			res = append(res, Location{row: currRow, col: currCol})
			currRow -= rowDiff
			currCol -= colDiff
		}

		// DOWNRIGHT
		currRow = y.row + rowDiff
		currCol = y.col + colDiff
		for currRow < rowLen && currCol < colLen {
			res = append(res, Location{row: currRow, col: currCol})
			currRow += rowDiff
			currCol += colDiff
		}
	} else {
		// UPRIGHT
		currRow := x.row - rowDiff
		currCol := x.col + colDiff
		for currRow >= 0 && currCol < colLen {
			res = append(res, Location{row: currRow, col: currCol})
			currRow -= rowDiff
			currCol += colDiff
		}

		// DOWNLEFT
		currRow = y.row + rowDiff
		currCol = y.col - colDiff
		for currRow < rowLen && currCol >= 0 {
			res = append(res, Location{row: currRow, col: currCol})
			currRow += rowDiff
			currCol -= colDiff
		}
	}
	return res
}

func P2() {
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
				location := Location{row: rowLen, col: col}
				nodeMap[char] = append(nodeMap[char], location)
				antinodeSet[location] = true
			}
		}
		rowLen += 1
	}

	for _, nodes := range nodeMap {
		for i := 0; i < len(nodes)-1; i++ {
			currNode := nodes[i]
			for j := i + 1; j < len(nodes); j++ {
				nextNode := nodes[j]
				antinodes := GetAntinodesList(rowLen, colLen, currNode, nextNode)
				for _, antinode := range antinodes {
					antinodeSet[antinode] = true
				}
			}
		}
	}

	total = len(antinodeSet)

	fmt.Println("D08 P2: ", total)
}
