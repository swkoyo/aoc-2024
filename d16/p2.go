package d16

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type SSet[T comparable] struct {
	data map[T]bool
}

func (s SSet[T]) Add(val T) {
	if !s.Exists(val) {
		s.data[val] = true
	}
}

func (s SSet[T]) Exists(val T) bool {
	_, exists := s.data[val]
	return exists
}

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
	lowestCost := make(map[string]int)
	lowestCost[start.ToString()] = 0
	backtrack := make(map[string]SSet[string])
	bestCost := int(math.MaxInt)
	endStates := SSet[string]{data: make(map[string]bool)}

	heap.Push(&pq, NewItem(0, start.row, start.col, start.dirRow, start.dirCol))

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		cost := item.priority
		r := item.value.row
		c := item.value.col
		dr := item.value.dirRow
		dc := item.value.dirCol

		lowest := int(math.MaxInt)
		if lc, exists := lowestCost[item.value.ToString()]; exists {
			lowest = lc
		}
		if cost > lowest {
			continue
		}

		if board[r][c] == 'E' {
			if cost > bestCost {
				break
			}
			bestCost = cost
			endStates.Add(item.value.ToString())
		}

		moves := make([]*Item, 0)
		moves = append(moves, NewItem(item.priority+1, r + dr, c + dc, dr, dc))
		moves = append(moves, NewItem(item.priority+1000, r, c, dc, dr * -1))
		moves = append(moves, NewItem(item.priority+1000, r, c, dc * -1, dr))

		for _, move := range moves {
			newCost := move.priority
			nr := move.value.row
			nc := move.value.col
			ndr := move.value.dirRow
			ndc := move.value.dirCol

			if board[nr][nc] == '#' {
				continue
			}

			lowest := int(math.MaxInt)
			if lc, exists := lowestCost[move.value.ToString()]; exists {
				lowest = lc
			}

			if  newCost > lowest {
				continue
			}

			if newCost < lowest {
				backtrack[move.value.ToString()] = SSet[string]{data: make(map[string]bool)}
				lowestCost[move.value.ToString()] = newCost
			}
			backtrack[move.value.ToString()].Add(item.value.ToString())
			heap.Push(&pq, NewItem(newCost, nr, nc, ndr, ndc))
		}
	}

	states := Queue[string]{data: make([]string, 0)}
	seen := SSet[string]{data: make(map[string]bool)}

	for state, _ := range endStates.data {
		states.Push(state)
		seen.Add(state)
	}

	for !states.IsEmpty() {
		key := states.Pop()
		arr := SSet[string]{data: make(map[string]bool)}

		if a, exists := backtrack[key]; exists {
			arr = a
		}

		for last, _ := range arr.data {
			if seen.Exists(last) {
				continue
			}
			seen.Add(last)
			states.Push(last)
		}
	}

	res := SSet[string]{data: make(map[string]bool)}

	for loc, _ := range seen.data {
		parts := strings.Split(loc, ",")	
		str := fmt.Sprintf("%s,%s", parts[0], parts[1])
		res.Add(str)
	}

	total = len(res.data)

	fmt.Println("D16 P2: ", total)
}
