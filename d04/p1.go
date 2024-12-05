package d04

import (
	"bufio"
	"fmt"
	"os"
)

func P1() {
	file, err := os.Open("d04/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	REF := "XMAS"

	total := 0

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for row, _ := range data {
		for col, _ := range data[row] {
			if data[row][col] == 'X' {
				// left
				if col >= 3 {
					leftRunes := []byte{data[row][col], data[row][col-1], data[row][col-2], data[row][col-3]}
					if string(leftRunes) == REF {
						total += 1
					}
				}

				// leftup
				if row >= 3 && col >= 3 {
					leftUpRunes := []byte{data[row][col], data[row-1][col-1], data[row-2][col-2], data[row-3][col-3]}
					if string(leftUpRunes) == REF {
						total += 1
					}
				}

				// up
				if row >= 3 {
					upRunes := []byte{data[row][col], data[row-1][col], data[row-2][col], data[row-3][col]}
					if string(upRunes) == REF {
						total += 1
					}
				}

				// upright
				if row >= 3 && col < len(data[0])-3 {
					upRightRunes := []byte{data[row][col], data[row-1][col+1], data[row-2][col+2], data[row-3][col+3]}
					if string(upRightRunes) == REF {
						total += 1
					}
				}

				// right
				if col < len(data[0])-3 {
					rightRunes := []byte{data[row][col], data[row][col+1], data[row][col+2], data[row][col+3]}
					if string(rightRunes) == REF {
						total += 1
					}
				}

				// downright
				if row < len(data)-3 && col < len(data[0])-3 {
					downRightRunes := []byte{data[row][col], data[row+1][col+1], data[row+2][col+2], data[row+3][col+3]}
					if string(downRightRunes) == REF {
						total += 1
					}
				}

				// down
				if row < len(data)-3 {
					downRunes := []byte{data[row][col], data[row+1][col], data[row+2][col], data[row+3][col]}
					if string(downRunes) == REF {
						total += 1
					}
				}

				// downleft
				if row < len(data)-3 && col >= 3 {
					downLeftRunes := []byte{data[row][col], data[row+1][col-1], data[row+2][col-2], data[row+3][col-3]}
					if string(downLeftRunes) == REF {
						total += 1
					}
				}
			}
		}
	}

	fmt.Println("D04 P1: ", total)
}
