package d14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	row int
	col int
}

func PrintGrid(rowLen, colLen int, robots []*Robot) {
	grid := make([][]rune, 0)
	for i := 0; i < rowLen; i++ {
		line := make([]rune, 0)
		for j := 0; j < colLen; j++ {
			line = append(line, '.')
		}
		grid = append(grid, line)
	}

	for _, robot := range robots {
		grid[robot.row][robot.col] = '#'
	}

	for _, line := range grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func HasTree(rowLen int, set map[Coord]bool) bool {
	ROWS_COUNT := 31

	for coord, _ := range set {
		if coord.row+ROWS_COUNT >= rowLen {
			continue
		}
		found := true
		for i := 1; i <= ROWS_COUNT; i++ {
			if _, exists := set[Coord{coord.row + i, coord.col}]; !exists {
				found = false
			}
		}
		if found {
			return true
		}
	}

	return false
}

func P2() {
	file, err := os.Open("d14/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	robots := make([]*Robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		robotPos := strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		robotV := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")

		pCol, _ := strconv.Atoi(robotPos[0])
		pRow, _ := strconv.Atoi(robotPos[1])
		vCol, _ := strconv.Atoi(robotV[0])
		vRow, _ := strconv.Atoi(robotV[1])

		robots = append(robots, &Robot{pRow, pCol, vRow, vCol})
	}

	ROW_LEN := 103
	COL_LEN := 101

	i := 1
	found := false
	for !found {
		isUnique := true
		set := make(map[Coord]bool)
		for _, robot := range robots {
			robot.Move(ROW_LEN, COL_LEN)
			coord := Coord{robot.row, robot.col}
			if _, exists := set[coord]; exists {
				isUnique = false
			} else {
				set[coord] = true
			}
		}
		if isUnique && HasTree(ROW_LEN, set) {
			PrintGrid(ROW_LEN, COL_LEN, robots)
			break
		}
		i += 1
	}

	total = i

	fmt.Println("D14 P2: ", total)
}
