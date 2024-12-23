package d20

import (
	"bufio"
	"fmt"
	"os"
)

func P2() {
	file, err := os.Open("d20/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	grid := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	var start Location

	for row, line := range grid {
		for col, char := range line {
			if char == 'S' {
				start = Location{row, col}
			}
		}
	}

	dists := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		line := make([]int, 0)
		for j := 0; j < len(grid[0]); j++ {
			line = append(line, -1)
		}
		dists = append(dists, line)
	}

	dists[start.row][start.col] = 0

	q := NewQueue()
	q.Push(start)

	rowDir := []int{1, 0, -1, 0}
	colDir := []int{0, 1, 0, -1}

	for !q.IsEmpty() {
		loc := q.Pop()
		for i := 0; i < 4; i++ {
			nr := loc.row + rowDir[i]
			nc := loc.col + colDir[i]

			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
				continue
			}

			if grid[nr][nc] == '#' {
				continue
			}

			if dists[nr][nc] != -1 {
				continue
			}

			dists[nr][nc] = dists[loc.row][loc.col] + 1
			q.Push(Location{nr, nc})
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '#' {
				continue
			}
			for radius := 2; radius < 21; radius++ {
				for dr := 0; dr < radius+1; dr++ {
					dc := radius - dr
					s := NewSet()
					s.Add(Location{r + dr, c + dc})
					s.Add(Location{r + dr, c - dc})
					s.Add(Location{r - dr, c + dc})
					s.Add(Location{r - dr, c - dc})
					for loc, _ := range s.data {
						nr := loc.row
						nc := loc.col

						if nr < 0 || nc < 0 || nr >= len(grid) || nc >= len(grid[0]) {
							continue
						}

						if grid[nr][nc] == '#' {
							continue
						}

						if dists[r][c]-dists[nr][nc] >= 100+radius {
							total += 1
						}
					}
				}

			}
		}
	}

	fmt.Println("D20 P2: ", total)
}
