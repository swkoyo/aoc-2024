package d06

import (
	"bufio"
	"fmt"
	"os"
)

type ValCoord struct {
	Row int
	Col int
	Val rune
}

type RuneBoard struct {
	Grid         [][]rune
	StartPos     Guard
	CurrentGuard Guard
	Set          map[ValCoord]bool
	Total        int
}

func (b *RuneBoard) Reset() {
	b.CurrentGuard = b.StartPos
	b.Set = make(map[ValCoord]bool)
}

func (b *RuneBoard) Move() bool {
	numTurns := 0

	if !b.AtEnd() && b.Peek() == '#' {
		for numTurns < 4 {
			b.Turn()
			if b.AtEnd() {
				return false
			}
			if b.Peek() != '#' {
				break
			}
			numTurns += 1
		}
	}

	if numTurns == 4 {
		return false
	}

	if b.AtEnd() {
		return false
	}

	b.CurrentGuard.Position = Coordinate{
		Row: b.CurrentGuard.Position.Row + b.CurrentGuard.NextMove.Row,
		Col: b.CurrentGuard.Position.Col + b.CurrentGuard.NextMove.Col,
	}

	_, exists := b.Set[ValCoord{Row: b.CurrentGuard.Position.Row, Col: b.CurrentGuard.Position.Col, Val: b.CurrentGuard.Val}]
	if exists {
		b.Total += 1
		return false
	}
	b.Set[ValCoord{Row: b.CurrentGuard.Position.Row, Col: b.CurrentGuard.Position.Col, Val: b.CurrentGuard.Val}] = true
	return true
}

func (b *RuneBoard) GetGuardPos() Guard {
	res := Guard{}
	for r, row := range b.Grid {
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

func (b *RuneBoard) Turn() {
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
}

func (b *RuneBoard) PositionIsInBounds(c Coordinate) bool {
	if c.Row < 0 || c.Row >= len(b.Grid) || c.Col < 0 || c.Col >= len(b.Grid[0]) {
		return false
	}
	return true
}

func (b *RuneBoard) AtEnd() bool {
	nextRow := b.CurrentGuard.Position.Row + b.CurrentGuard.NextMove.Row
	nextCol := b.CurrentGuard.Position.Col + b.CurrentGuard.NextMove.Col
	return !b.PositionIsInBounds(Coordinate{Row: nextRow, Col: nextCol})
}

func (b *RuneBoard) Peek() rune {
	nextRow := b.CurrentGuard.Position.Row + b.CurrentGuard.NextMove.Row
	nextCol := b.CurrentGuard.Position.Col + b.CurrentGuard.NextMove.Col
	return rune(b.Grid[nextRow][nextCol])
}

func P2() {
	file, err := os.Open("d06/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	board := RuneBoard{
		Grid:  make([][]rune, 0),
		Set:   make(map[ValCoord]bool),
		Total: 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		board.Grid = append(board.Grid, []rune(line))
	}

	board.CurrentGuard = board.GetGuardPos()
	board.StartPos = board.GetGuardPos()

	for currentRow := 0; currentRow < len(board.Grid); currentRow++ {
		for currentCol := 0; currentCol < len(board.Grid[currentRow]); currentCol++ {
			currentRune := board.Grid[currentRow][currentCol]
			if currentRune == '#' {
				continue
			}
			if currentRune == '^' || currentRune == '>' || currentRune == 'v' || currentRune == '<' {
				continue
			}
			board.Grid[currentRow][currentCol] = '#'
			for {
				moved := board.Move()
				if !moved {
					break
				}
			}
			board.Reset()
			board.Grid[currentRow][currentCol] = '.'
		}
	}

	fmt.Println("D06 P2: ", board.Total)
}
