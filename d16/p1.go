package d16

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Set[T comparable] struct {
	data map[T]bool
}

func (s *Set[T]) Add(val T) {
	if !s.Exists(val) {
		s.data[val] = true
	}
}

func (s *Set[T]) Exists(val T) bool {
	_, exists := s.data[val]
	return exists
}

type Location struct {
	row        int
	col        int
	dirRow     int
	dirCol     int
	prevRow    int
	prevCol    int
	prevDirRow int
	prevDirCol int
}

type Item struct {
	priority int
	index    int
	value    Location
}

func NewItem(priority, row, col, dirRow, dirCol int) *Item {
	return &Item{
		priority: priority,
		value: Location{
			row:    row,
			col:    col,
			dirRow: dirRow,
			dirCol: dirCol,
		},
	}
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value Location, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func P1() {
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
	seen := Set[Location]{data: make(map[Location]bool)}

	heap.Push(&pq, NewItem(0, start.row, start.col, start.dirRow, start.dirCol))
	seen.Add(start)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		seen.Add(item.value)

		if board[item.value.row][item.value.col] == 'E' {
			total = item.priority
			break
		}

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

			if seen.Exists(Location{row: newRow, col: newCol, dirRow: newDirRow, dirCol: newDirCol}) {
				continue
			}

			heap.Push(&pq, NewItem(newCost, newRow, newCol, newDirRow, newDirCol))
		}
	}

	fmt.Println("D16 P1: ", total)
}
