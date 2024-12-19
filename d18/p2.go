package d18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P2() {
	file, err := os.Open("d18/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	gridLen := 71
	grid := make([][]rune, 0)
	remainingLoc := make([]Location, 0)

	for i := 0; i < gridLen; i++ {
		line := make([]rune, 0)
		for j := 0; j < gridLen; j++ {
			line = append(line, '.')
		}
		grid = append(grid, line)
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		col, _ := strconv.Atoi(parts[0])
		row, _ := strconv.Atoi(parts[1])
		if i < 1024 {
			grid[row][col] = '#'
		} else {
			remainingLoc = append(remainingLoc, Location{row, col, 0})
		}
		i += 1
	}

	// PrintGrid(grid)

	er := gridLen - 1
	ec := gridLen - 1

	rowDir := []int{1, 0, -1, 0}
	colDir := []int{0, 1, 0, -1}

	for _, currLoc := range remainingLoc {
		grid[currLoc.row][currLoc.col] = '#'

		queue := Queue{data: make([]*Location, 0)}
		queue.Push(&Location{0, 0, 0})
		seen := Set{data: make(map[string]bool)}
		seen.Add("0,0")
		isFound := false

		for !queue.IsEmpty() && !isFound {
			loc := queue.Pop()
			for i := 0; i < 4; i++ {
				nr := loc.row + rowDir[i]
				nc := loc.col + colDir[i]

				if nr < 0 || nr >= gridLen {
					continue
				}

				if nc < 0 || nc >= gridLen {
					continue
				}

				if grid[nr][nc] == '#' {
					continue
				}

				newLoc := Location{nr, nc, loc.dist + 1}
				if seen.Exists(newLoc.ToString()) {
					continue
				}

				if nr == er && nc == ec {
					isFound = true
					break
				}

				seen.Add(newLoc.ToString())

				queue.Push(&newLoc)
			}
		}

		if !isFound {
			fmt.Println("D18 P2: ", fmt.Sprintf("%d,%d", currLoc.col, currLoc.row))
			return
		}
	}
}
