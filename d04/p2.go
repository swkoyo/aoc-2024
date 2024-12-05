package d04

import (
	"bufio"
	"fmt"
	"os"
)

func P2() {
	file, err := os.Open("d04/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	REF := "MAS"
	REV_REF := "SAM"

	total := 0

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for row, _ := range data {
		for col, _ := range data[row] {
			if data[row][col] == 'A' && row >= 1 && row < len(data)-1 && col >= 1 && col < len(data[0])-1 {
				// downright
				dr := string([]byte{data[row-1][col-1], data[row][col], data[row+1][col+1]})

				// downleft
				dl := string([]byte{data[row-1][col+1], data[row][col], data[row+1][col-1]})

				if (dr == REF || dr == REV_REF) && (dl == REF || dl == REV_REF) {
					total += 1
				}
			}
		}
	}

	fmt.Println("D04 P2: ", total)
}
