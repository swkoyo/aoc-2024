package d06

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	Row int
	Col int
}

type Guard struct {
	Position Coordinate
	Val      rune
	NextMove Coordinate
}

type Board struct {
	Grid         []string
	CurrentGuard Guard
	Set          map[Coordinate]bool
	Total        int
}

func (b *Board) PositionIsInBounds(c Coordinate) bool {
	if c.Row < 0 || c.Row >= len(b.Grid) || c.Col < 0 || c.Col >= len(b.Grid[0]) {
		return false
	}
	return true
}

func (b *Board) MoveGuard() {
	nextRune, isIn := b.PeekGuard()
	for isIn && nextRune == '#' {
		if b.CurrentGuard.Val == '<' {
			b.CurrentGuard.Val = '^'
			b.CurrentGuard.NextMove = Coordinate{Row: -1, Col: 0}
		} else if b.CurrentGuard.Val == '^' {
			b.CurrentGuard.Val = '>'
			b.CurrentGuard.NextMove = Coordinate{Row: 0, Col: 1}
		} else if b.CurrentGuard.Val == '>' {
			b.CurrentGuard.Val = 'v'
			b.CurrentGuard.NextMove = Coordinate{Row: 1, Col: 0}
		} else {
			b.CurrentGuard.Val = '<'
			b.CurrentGuard.NextMove = Coordinate{Row: 0, Col: -1}
		}
		nextRune, isIn = b.PeekGuard()
	}

	b.CurrentGuard.Position = Coordinate{
		Row: b.CurrentGuard.Position.Row + b.CurrentGuard.NextMove.Row,
		Col: b.CurrentGuard.Position.Col + b.CurrentGuard.NextMove.Col,
	}

	if !b.PositionIsInBounds(b.CurrentGuard.Position) {
		return
	}

	if _, exists := b.Set[b.CurrentGuard.Position]; !exists {
		b.Set[b.CurrentGuard.Position] = true
		b.Total += 1
	}
}

func (b *Board) PeekGuard() (rune, bool) {
	nextRow := b.CurrentGuard.Position.Row + b.CurrentGuard.NextMove.Row
	nextCol := b.CurrentGuard.Position.Col + b.CurrentGuard.NextMove.Col
	if !b.PositionIsInBounds(Coordinate{Row: nextRow, Col: nextCol}) {
		return '_', false
	}
	return rune(b.Grid[nextRow][nextCol]), true
}

func GetGuardPosition(grid []string) Guard {
	res := Guard{}
	for r, row := range grid {
		for c, char := range row {
			if char == '<' || char == '^' || char == '>' || char == 'v' {
				res.Position.Row = r
				res.Position.Col = c
				res.Val = char

				if char == '<' {
					res.NextMove = Coordinate{Row: 0, Col: -1}
				} else if char == '^' {
					res.NextMove = Coordinate{Row: -1, Col: 0}
				} else if char == '>' {
					res.NextMove = Coordinate{Row: 0, Col: 1}
				} else {
					res.NextMove = Coordinate{Row: 1, Col: 0}
				}
				return res
			}
		}
	}
	return res
}

func P1() {
	file, err := os.Open("d06/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	board := Board{
		Grid:         make([]string, 0),
		Set:          make(map[Coordinate]bool),
		Total:        0,
		CurrentGuard: Guard{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		board.Grid = append(board.Grid, line)
	}

	board.CurrentGuard = GetGuardPosition(board.Grid)
	board.Set[board.CurrentGuard.Position] = true
	board.Total = 1

	for {
		board.MoveGuard()
		if !board.PositionIsInBounds(board.CurrentGuard.Position) {
			break
		}
	}

	fmt.Println("D06 P1: ", board.Total)
}
