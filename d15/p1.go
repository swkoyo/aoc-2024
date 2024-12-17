package d15

import (
	"bufio"
	"fmt"
	"os"
)

type Location struct {
	row int
	col int
}

type Warehouse struct {
	grid  [][]rune
	robot Location
}

func (w *Warehouse) GetGPS() int {
	total := 0
	for row, line := range w.grid {
		for col, char := range line {
			if char == 'O' {
				total += (100*row + col)
			}
		}
	}
	return total
}

func (w *Warehouse) Print() {
	for _, line := range w.grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (w *Warehouse) Swap(xRow, xCol, yRow, yCol int) {
	w.grid[xRow][xCol], w.grid[yRow][yCol] = w.grid[yRow][yCol], w.grid[xRow][xCol]
}

func (w *Warehouse) MoveBot(dir rune) {
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
		} else if next == 'O' {
			currC := nextC - 1
			isValid := true
			for currC >= 0 {
				if w.grid[nextR][currC] == '#' {
					isValid = false
					break
				}
				if w.grid[nextR][currC] == '.' {
					break
				}
				currC -= 1
			}
			if isValid && currC >= 0 {
				w.Swap(nextR, nextC, nextR, currC)
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
		} else if next == 'O' {
			currR := nextR - 1
			isValid := true
			for currR >= 0 {
				if w.grid[currR][nextC] == '#' {
					isValid = false
					break
				}
				if w.grid[currR][nextC] == '.' {
					break
				}
				currR -= 1
			}
			if isValid && currR >= 0 {
				w.Swap(nextR, nextC, currR, nextC)
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
		} else if next == 'O' {
			currC := nextC + 1
			isValid := true
			for currC < len(w.grid[0]) {
				if w.grid[nextR][currC] == '#' {
					isValid = false
					break
				}
				if w.grid[nextR][currC] == '.' {
					break
				}
				currC += 1
			}
			if currC < len(w.grid[0]) && isValid {
				w.Swap(nextR, nextC, nextR, currC)
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
		} else if next == 'O' {
			currR := nextR + 1
			isValid := true
			for currR < len(w.grid) {
				if w.grid[currR][nextC] == '#' {
					isValid = false
					break
				}
				if w.grid[currR][nextC] == '.' {
					break
				}
				currR += 1
			}
			if currR < len(w.grid) && isValid {
				w.Swap(nextR, nextC, currR, nextC)
				w.Swap(nextR, nextC, w.robot.row, w.robot.col)
				w.robot.row = nextR
				w.robot.col = nextC
			}
			return
		}
	}
}

func P1() {
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
			warehouse.grid = append(warehouse.grid, []rune(line))
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

	for _, move := range moves {
		warehouse.MoveBot(move)
	}

	total = warehouse.GetGPS()

	fmt.Println("D15 P1: ", total)
}
