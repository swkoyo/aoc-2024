package d15

import (
	"bufio"
	"fmt"
	"os"
)

type Set struct {
	data map[Location]bool
}

func NewSet() Set {
	return Set{
		data: make(map[Location]bool),
	}
}

func (s *Set) Exists(l Location) bool {
	_, exists := s.data[l]
	return exists
}

func (s *Set) Add(l Location) {
	if !s.Exists(l) {
		s.data[l] = true
	}
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

func (w *Warehouse) GetGPSHard() int {
	total := 0
	for row, line := range w.grid {
		for col, char := range line {
			if char == '[' {
				total += (100*row + col)
			}
		}
	}
	return total
}

func (w *Warehouse) MoveBotHard(dir rune) {
	nextR := w.robot.row
	nextC := w.robot.col
	if dir == '<' {
		nextC -= 1
		next := w.grid[nextR][nextC]
		if next == '.' {
			w.Swap(nextR, nextC, w.robot.row, w.robot.col)
			w.robot.row = nextR
			w.robot.col = nextC
			return
		} else if next == '#' {
			return
		} else if next == ']' {
			currC := nextC - 2
			isValid := true
			for currC >= 0 {
				if w.grid[nextR][currC] == '#' {
					isValid = false
					break
				}
				if w.grid[nextR][currC] == '.' {
					break
				}
				currC -= 2
			}
			if isValid && currC >= 0 {
				for i := currC; i < nextC; i++ {
					w.Swap(nextR, i, nextR, i+1)
				}
				w.Swap(nextR, nextC, w.robot.row, w.robot.col)
				w.robot.row = nextR
				w.robot.col = nextC
			}
			return
		}
	} else if dir == '>' {
		nextC += 1
		next := w.grid[nextR][nextC]
		if next == '.' {
			w.Swap(nextR, nextC, w.robot.row, w.robot.col)
			w.robot.row = nextR
			w.robot.col = nextC
			return
		} else if next == '#' {
			return
		} else if next == '[' {
			currC := nextC + 2
			isValid := true
			for currC < len(w.grid[0]) {
				if w.grid[nextR][currC] == '#' {
					isValid = false
					break
				}
				if w.grid[nextR][currC] == '.' {
					break
				}
				currC += 2
			}
			if currC < len(w.grid[0]) && isValid {
				for i := currC; i > nextC; i-- {
					w.Swap(nextR, i, nextR, i-1)
				}
				w.Swap(nextR, nextC, w.robot.row, w.robot.col)
				w.robot.row = nextR
				w.robot.col = nextC
			}
			return
		}
	} else if dir == '^' {
		nextR -= 1
		next := w.grid[nextR][nextC]
		if next == '.' {
			w.Swap(nextR, nextC, w.robot.row, w.robot.col)
			w.robot.row = nextR
			w.robot.col = nextC
			return
		} else if next == '#' {
			return
		} else if next == '[' || next == ']' {
			blocks := make([]Location, 0)
			visited := NewSet()
			queue := NewQueue()
			queue.Push(Location{nextR, nextC})
			if w.grid[nextR][nextC] == '[' {
				queue.Push(Location{nextR, nextC + 1})
			} else {
				queue.Push(Location{nextR, nextC - 1})
			}
			isValid := true
			for !queue.IsEmpty() && isValid {
				loc := queue.Pop()
				if visited.Exists(loc) {
					continue
				}
				blocks = append(blocks, loc)
				visited.Add(loc)
				next := w.grid[loc.row-1][loc.col]
				if next == '#' {
					isValid = false
				} else if next == '[' || next == ']' {
					queue.Push(Location{loc.row - 1, loc.col})
					if next == '[' {
						queue.Push(Location{loc.row - 1, loc.col + 1})
					} else {
						queue.Push(Location{loc.row - 1, loc.col - 1})
					}
				}
			}

			if isValid {
				for i := len(blocks) - 1; i >= 0; i-- {
					block := blocks[i]
					w.Swap(block.row, block.col, block.row-1, block.col)
				}
				w.Swap(nextR, nextC, w.robot.row, w.robot.col)
				w.robot.row = nextR
				w.robot.col = nextC
			}

			return
		}
	} else if dir == 'v' {
		nextR += 1
		next := w.grid[nextR][nextC]
		if next == '.' {
			w.Swap(nextR, nextC, w.robot.row, w.robot.col)
			w.robot.row = nextR
			w.robot.col = nextC
			return
		} else if next == '#' {
			return
		} else if next == '[' || next == ']' {
			blocks := make([]Location, 0)
			visited := NewSet()
			queue := NewQueue()
			queue.Push(Location{nextR, nextC})
			if w.grid[nextR][nextC] == '[' {
				queue.Push(Location{nextR, nextC + 1})
			} else {
				queue.Push(Location{nextR, nextC - 1})
			}
			isValid := true
			for !queue.IsEmpty() && isValid {
				loc := queue.Pop()
				if visited.Exists(loc) {
					continue
				}
				blocks = append(blocks, loc)
				visited.Add(loc)
				next := w.grid[loc.row+1][loc.col]
				if next == '#' {
					isValid = false
				} else if next == '[' || next == ']' {
					queue.Push(Location{loc.row + 1, loc.col})
					if next == '[' {
						queue.Push(Location{loc.row + 1, loc.col + 1})
					} else {
						queue.Push(Location{loc.row + 1, loc.col - 1})
					}
				}
			}

			if isValid {
				for i := len(blocks) - 1; i >= 0; i-- {
					block := blocks[i]
					w.Swap(block.row, block.col, block.row+1, block.col)
				}
				w.Swap(nextR, nextC, w.robot.row, w.robot.col)
				w.robot.row = nextR
				w.robot.col = nextC
			}

			return
		}
	}
}

func P2() {
	file, err := os.Open("d15/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	warehouse := Warehouse{
		grid: make([][]rune, 0),
	}
	moves := make([]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			currLine := make([]rune, 0)
			for _, char := range line {
				if char == '#' {
					currLine = append(currLine, []rune{'#', '#'}...)
				} else if char == 'O' {
					currLine = append(currLine, []rune{'[', ']'}...)
				} else if char == '.' {
					currLine = append(currLine, []rune{'.', '.'}...)
				} else if char == '@' {
					currLine = append(currLine, []rune{'@', '.'}...)
				}
			}
			warehouse.grid = append(warehouse.grid, currLine)
		} else {
			moves = append(moves, []rune(line)...)
		}
	}

	for row, line := range warehouse.grid {
		for col, char := range line {
			if char == '@' {
				warehouse.robot = Location{row, col}
			}
		}
	}

	fmt.Println("INIT:")
	warehouse.Print()
	for _, move := range moves {
		// fmt.Printf("MOVE: %s\n", string(move))
		warehouse.MoveBotHard(move)
		// warehouse.Print()
	}
	fmt.Println("FINAL:")
	warehouse.Print()

	total = warehouse.GetGPSHard()

	fmt.Println("D15 P2: ", total)
}
