package d18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Location struct {
	row  int
	col  int
	dist int
}

func (l *Location) ToString() string {
	return fmt.Sprintf("%d,%d", l.row, l.col)
}

type Queue struct {
	data []*Location
}

func (q *Queue) Push(v *Location) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() *Location {
	c := q.data[0]
	q.data = q.data[1:]
	return c
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func PrintGrid(grid [][]rune) {
	fmt.Println("----------------------------------------------------")
	for _, line := range grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------------")
}

func P1() {
	file, err := os.Open("d18/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	gridLen := 71
	grid := make([][]rune, 0)

	for i := 0; i < gridLen; i++ {
		line := make([]rune, 0)
		for j := 0; j < gridLen; j++ {
			line = append(line, '.')
		}
		grid = append(grid, line)
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i >= 1024 {
			break
		}
		line := scanner.Text()
		parts := strings.Split(line, ",")
		col, _ := strconv.Atoi(parts[0])
		row, _ := strconv.Atoi(parts[1])
		grid[row][col] = '#'
		i += 1
	}

	PrintGrid(grid)

	er := gridLen - 1
	ec := gridLen - 1

	rowDir := []int{1, 0, -1, 0}
	colDir := []int{0, 1, 0, -1}

	queue := Queue{data: make([]*Location, 0)}
	queue.Push(&Location{0, 0, 0})
	seen := Set{data: make(map[string]bool)}
	seen.Add("0,0")
	isFound := false

	for !queue.IsEmpty() && !isFound {
		loc := queue.Pop()
		for i := 0; i < 4; i++ {
			nr := loc.row + rowDir[i]
			nc := loc.col + colDir[i]

			if nr < 0 || nr >= gridLen {
				continue
			}

			if nc < 0 || nc >= gridLen {
				continue
			}

			if grid[nr][nc] == '#' {
				continue
			}

			newLoc := Location{nr, nc, loc.dist + 1}
			if seen.Exists(newLoc.ToString()) {
				continue
			}

			if nr == er && nc == ec {
				total = newLoc.dist
				isFound = true
				break
			}

			seen.Add(newLoc.ToString())

			queue.Push(&newLoc)
		}
	}

	fmt.Println("D18 P1: ", total)
}

//
// import (
// 	"bufio"
// 	"container/heap"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )
//
// type Set struct {
// 	data map[string]bool
// }
//
// func (s *Set) Add(l string) {
// 	if !s.Exists(l) {
// 		s.data[l] = true
// 	}
// }
//
// func (s *Set) Exists(l string) bool {
// 	_, exists := s.data[l]
// 	return exists
// }
//
// type Location struct {
// 	row int
// 	col int
// }
//
// func (l *Location) ToString() string {
// 	return fmt.Sprintf("%d,%d", l.row, l.col)
// }
//
// type Item struct {
// 	priority int
// 	index    int
// 	value    Location
// }
//
// func NewItem(priority, row, col int) *Item {
// 	return &Item{
// 		priority: priority,
// 		value: Location{
// 			row: row,
// 			col: col,
// 		},
// 	}
// }
//
// type PriorityQueue []*Item
//
// func (pq PriorityQueue) Len() int { return len(pq) }
//
// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].priority < pq[j].priority
// }
//
// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }
//
// func (pq *PriorityQueue) Push(x any) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }
//
// func (pq *PriorityQueue) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil
// 	item.index = -1
// 	*pq = old[0 : n-1]
// 	return item
// }
//
// func (pq *PriorityQueue) update(item *Item, value Location, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }
//
// func PrintGrid(grid [][]rune) {
// 	fmt.Println("----------------------------------------------------")
// 	for _, line := range grid {
// 		for _, char := range line {
// 			fmt.Print(string(char))
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("----------------------------------------------------")
// }
//
// func P1() {
// 	file, err := os.Open("d18/input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
//
// 	total := 0
//
// 	gridLen := 71
// 	grid := make([][]rune, 0)
//
// 	for i := 0; i < gridLen; i++ {
// 		line := make([]rune, 0)
// 		for j := 0; j < gridLen; j++ {
// 			line = append(line, '.')
// 		}
// 		grid = append(grid, line)
// 	}
//
// 	i := 0
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		if i >= 1024 {
// 			break
// 		}
// 		line := scanner.Text()
// 		parts := strings.Split(line, ",")
// 		col, _ := strconv.Atoi(parts[0])
// 		row, _ := strconv.Atoi(parts[1])
// 		grid[row][col] = '#'
// 		i += 1
// 	}
//
// 	PrintGrid(grid)
//
// 	er := gridLen - 1
// 	ec := gridLen - 1
//
// 	rowDir := []int{1, 0, -1, 0}
// 	colDir := []int{0, 1, 0, -1}
//
// 	pq := make(PriorityQueue, 0)
// 	seen := Set{data: make(map[string]bool)}
//
// 	heap.Push(&pq, NewItem(0, 0, 0))
// 	l := Location{0, 0}
// 	seen.Add(l.ToString())
//
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		seen.Add(item.value.ToString())
//
// 		// fmt.Println(item.value.row, item.value.col)
// 		if item.value.row == er && item.value.col == ec {
// 			total = item.priority
// 			break
// 		}
//
// 		for i := 0; i < 4; i++ {
// 			newCost := item.priority + 1
// 			nr := item.value.row + rowDir[i]
// 			nc := item.value.col + colDir[i]
//
// 			if nr < 0 || nr >= gridLen {
// 				continue
// 			}
//
// 			if nc < 0 || nc >= gridLen {
// 				continue
// 			}
//
// 			if grid[nr][nc] == '#' {
// 				continue
// 			}
//
// 			newLoc := Location{nr, nc}
// 			if seen.Exists(newLoc.ToString()) {
// 				continue
// 			}
//
// 			heap.Push(&pq, NewItem(newCost, nr, nc))
// 		}
// 	}
//
// 	fmt.Println("D18 P1: ", total)
// }
