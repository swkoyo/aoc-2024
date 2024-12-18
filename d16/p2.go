package d16

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Queue[T comparable] struct {
	data []T
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Pop() T {
	c := q.data[0]
	q.data = q.data[1:]
	return c
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func NewItemWithPrev(priority, row, col, dirRow, dirCol, prevRow, prevCol, prevDirRow, prevDirCol int) *Item {
	return &Item{
		priority: priority,
		value: Location{
			row:        row,
			col:        col,
			dirRow:     dirRow,
			dirCol:     dirCol,
			prevRow:    prevRow,
			prevCol:    prevCol,
			prevDirRow: prevDirRow,
			prevDirCol: prevDirCol,
		},
	}
}

func P2() {
	file, err := os.Open("d16/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	board := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, []rune(line))
	}

	var start Location

	for row, line := range board {
		for col, char := range line {
			if char == 'S' {
				start = Location{row: row, col: col, dirRow: 0, dirCol: 1}
			}
		}
	}

	pq := make(PriorityQueue, 0)
	lowestCost := make(map[Location]int)
	lowestCost[start] = 0
	backtrack := make(map[Location]*Set[Location])
	bestCost := int(math.MaxInt)
	endStates := Set[Location]{data: make(map[Location]bool)}

	heap.Push(&pq, NewItemWithPrev(0, start.row, start.col, start.dirRow, start.dirCol, -1, -1, -1, -1))

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if cost, exists := lowestCost[item.value]; exists && item.priority > cost {
			continue
		}
		lowestCost[item.value] = item.priority

		if board[item.value.row][item.value.col] == 'E' {
			if item.priority > bestCost {
				break
			}
			bestCost = item.priority
			endStates.Add(Location{row: item.value.row, col: item.value.col, dirRow: item.value.dirRow, dirCol: item.value.dirCol})
		}

		currLoc := Location{row: item.value.row, col: item.value.col, dirRow: item.value.dirRow, dirCol: item.value.dirCol}
		if _, exists := backtrack[currLoc]; !exists {
			backtrack[currLoc] = &Set[Location]{data: make(map[Location]bool)}
		}
		backtrack[currLoc].Add(Location{row: item.value.prevRow, col: item.value.prevCol, dirRow: item.value.prevDirRow, dirCol: item.value.prevDirCol})

		moves := make([]*Item, 0)
		moves = append(moves, NewItem(item.priority+1, item.value.row+item.value.dirRow, item.value.col+item.value.dirCol, item.value.dirRow, item.value.dirCol))
		moves = append(moves, NewItem(item.priority+1000, item.value.row, item.value.col, item.value.dirCol, item.value.dirRow*-1))
		moves = append(moves, NewItem(item.priority+1000, item.value.row, item.value.col, item.value.dirCol*-1, item.value.dirRow))

		for _, move := range moves {
			newCost := move.priority
			newRow := move.value.row
			newCol := move.value.col
			newDirRow := move.value.dirRow
			newDirCol := move.value.dirCol

			if board[newRow][newCol] == '#' {
				continue
			}

			if cost, exists := lowestCost[Location{row: newRow, col: newCol, dirRow: newDirRow, dirCol: newDirCol}]; exists && item.priority > cost {
				continue
			}

			heap.Push(&pq, NewItemWithPrev(newCost, newRow, newCol, newDirRow, newDirCol, item.value.row, item.value.col, item.value.dirRow, item.value.dirCol))
		}
	}

	states := Queue[Location]{data: make([]Location, 0)}
	seen := Set[Location]{data: make(map[Location]bool)}

	for state, _ := range endStates.data {
		states.Push(state)
		seen.Add(state)
	}

	for !states.IsEmpty() {
		loc := states.Pop()
		if _, exists := backtrack[loc]; !exists {
			continue
		}
		for last, _ := range backtrack[loc].data {
			if seen.Exists(last) {
				continue
			}
			seen.Add(last)
			states.Push(last)
		}
	}

	// for loc, _ := range seen.data {
	// 	fmt.Printf("(%d, %d)\n", loc.row, loc.col)
	// }

	total = len(seen.data)

	fmt.Println("D16 P2: ", total)
}
