package d14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	row  int
	col  int
	rowV int
	colV int
}

func (r *Robot) Move(rowLen, colLen int) {
	r.row += r.rowV
	if r.row < 0 {
		r.row += rowLen
	} else if r.row >= rowLen {
		r.row %= rowLen
	}
	r.col += r.colV
	if r.col < 0 {
		r.col += colLen
	} else if r.col >= colLen {
		r.col %= colLen
	}
}

func P1() {
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
	SECONDS := 100

	for i := 0; i < SECONDS; i++ {
		for _, robot := range robots {
			robot.Move(ROW_LEN, COL_LEN)
		}
	}

	MIDROW := ROW_LEN / 2
	MIDCOL := COL_LEN / 2

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, robot := range robots {
		// MIDDLE
		if robot.row == MIDROW || robot.col == MIDCOL {
			continue
		}

		// Q1
		if robot.row < MIDROW && robot.col < MIDCOL {
			q1 += 1
		}

		// Q2
		if robot.row < MIDROW && robot.col > MIDCOL {
			q2 += 1
		}

		// Q3
		if robot.row > MIDROW && robot.col < MIDCOL {
			q3 += 1
		}

		// Q4
		if robot.row > MIDROW && robot.col > MIDCOL {
			q4 += 1
		}
	}

	total = q1 * q2 * q3 * q4

	fmt.Println("D14 P1: ", total)
}
