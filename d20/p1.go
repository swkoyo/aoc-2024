package d20

import (
	"bufio"
	"fmt"
	"os"
)

type Set struct {
	data map[string]bool
}

func (s *Set) Add(l string) {
	if !s.Exists(l) {
		s.data[l] = true
	}
}

func (s *Set) Exists(l string) bool {
	_, exists := s.data[l]
	return exists
}

func NewSet() Set {
	return Set{map[string]bool{}}
}

type Location struct {
	row int
	col int
}

type Queue struct {
	data []Location
}

func NewQueue() Queue {
	return Queue{
		data: make([]Location, 0),
	}
}

func (q *Queue) Push(c Location) {
	q.data = append(q.data, c)
}

func (q *Queue) Pop() Location {
	c := q.data[0]
	q.data = q.data[1:]
	return c
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func P1() {
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

	// for _, line := range dists {
	// 	for _, dist := range line {
	// 		fmt.Printf("%d\t", dist)
	// 	}
	// 	fmt.Println()
	// }

	jumpRowDir := []int{2, 1, 0, -1, -2, -1, 0, 1}
	jumpColDir := []int{0, 1, 2, 1, 0, -1, -2, -1}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '#' {
				continue
			}
			for i := 0; i < len(jumpRowDir); i++ {
				nr := r + jumpRowDir[i]
				nc := c + jumpColDir[i]

				if nr < 0 || nc < 0 || nr >= len(grid) || nc >= len(grid[0]) {
					continue
				}

				if grid[nr][nc] == '#' {
					continue
				}

				if dists[r][c]-dists[nr][nc] >= 102 {
					total += 1
				}
			}
		}
	}

	fmt.Println("D20 P1: ", total)
}
