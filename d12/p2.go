package d12

import (
	"bufio"
	"fmt"
	"os"
)

type Set struct {
	data map[Coordinate]bool
}

func NewSet() Set {
	return Set{
		data: make(map[Coordinate]bool),
	}
}

func (s *Set) Length() int {
	return len(s.data)
}

func (s *Set) Exists(c Coordinate) bool {
	_, exists := s.data[c]
	return exists
}

func (s *Set) Add(c Coordinate) {
	if !s.Exists(c) {
		s.data[c] = true
	}
}

func (s *Set) Union(other Set) {
	for c, _ := range other.data {
		s.Add(c)
	}
}

type Queue struct {
	data []Coordinate
}

func NewQueue() Queue {
	return Queue{
		data: make([]Coordinate, 0),
	}
}

func (q *Queue) Push(c Coordinate) {
	q.data = append(q.data, c)
}

func (q *Queue) Pop() Coordinate {
	c := q.data[0]
	q.data = q.data[1:]
	return c
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func Permimeter(region Set) int {
	total := region.Length() * 4
	for coord, _ := range region.data {
		nextCoords := []Coordinate{
			Coordinate{coord.row - 1, coord.col},
			Coordinate{coord.row, coord.col + 1},
			Coordinate{coord.row + 1, coord.col},
			Coordinate{coord.row, coord.col - 1},
		}
		for _, nextCoord := range nextCoords {
			if region.Exists(nextCoord) {
				total -= 1
			}
		}
	}
	return total
}

func IsInBounds(data []string, coord Coordinate) bool {
	return coord.row >= 0 && coord.row < len(data) && coord.col >= 0 && coord.col < len(data[0])
}

func GetPlant(data []string, coord Coordinate) rune {
	if !IsInBounds(data, coord) {
		return '_'
	}
	return rune(data[coord.row][coord.col])
}

func NumCorners(data []string, coord Coordinate) int {
	total := 0

	up := GetPlant(data, Coordinate{coord.row - 1, coord.col})
	rightUp := GetPlant(data, Coordinate{coord.row - 1, coord.col + 1})
	right := GetPlant(data, Coordinate{coord.row, coord.col + 1})
	rightDown := GetPlant(data, Coordinate{coord.row + 1, coord.col + 1})
	down := GetPlant(data, Coordinate{coord.row + 1, coord.col})
	leftDown := GetPlant(data, Coordinate{coord.row + 1, coord.col - 1})
	left := GetPlant(data, Coordinate{coord.row, coord.col - 1})
	leftUp := GetPlant(data, Coordinate{coord.row - 1, coord.col - 1})

	currPlant := GetPlant(data, coord)

	// Diagonal
	// UP /
	if up != currPlant && left != currPlant {
		total += 1
	}
	// UP \
	if up != currPlant && right != currPlant {
		total += 1
	}
	// DOWN \
	if down != currPlant && left != currPlant {
		total += 1
	}
	// DOWN /
	if down != currPlant && right != currPlant {
		total += 1
	}

	// Inner
	// LEFTUP
	if left == currPlant && up == currPlant && leftUp != currPlant {
		total += 1
	}
	// LEFTDOWN
	if left == currPlant && down == currPlant && leftDown != currPlant {
		total += 1
	}
	// RIGHTUP
	if right == currPlant && up == currPlant && rightUp != currPlant {
		total += 1
	}
	// RIGHTDOWN
	if right == currPlant && down == currPlant && rightDown != currPlant {
		total += 1
	}

	return total
}

func Sides(data []string, region Set) int {
	total := 0
	for coord, _ := range region.data {
		total += NumCorners(data, coord)
	}
	return total
}

func P2() {
	file, err := os.Open("d12/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	seen := NewSet()
	regions := make([]Set, 0)
	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for row, currRow := range data {
		for col, currChar := range currRow {
			currCoord := Coordinate{row, col}
			if seen.Exists(currCoord) {
				continue
			}
			region := NewSet()
			region.Add(currCoord)
			queue := NewQueue()
			queue.Push(currCoord)
			for !queue.IsEmpty() {
				coord := queue.Pop()
				nextCoords := []Coordinate{
					Coordinate{coord.row - 1, coord.col},
					Coordinate{coord.row, coord.col + 1},
					Coordinate{coord.row + 1, coord.col},
					Coordinate{coord.row, coord.col - 1},
				}
				for _, nextCoord := range nextCoords {
					if region.Exists(nextCoord) {
						continue
					}
					if nextCoord.row < 0 || nextCoord.row >= len(data) || nextCoord.col < 0 || nextCoord.col >= len(data[0]) {
						continue
					}
					if data[nextCoord.row][nextCoord.col] != byte(currChar) {
						continue
					}
					region.Add(nextCoord)
					queue.Push(nextCoord)
				}
			}
			regions = append(regions, region)
			seen.Union(region)
		}
	}

	for _, region := range regions {
		total += (region.Length() * Sides(data, region))
	}

	fmt.Println("D12 P2: ", total)
}
