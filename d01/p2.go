package d01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func P2() {
	file, err := os.Open("d01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	xarr := make([]int, 0)
	yarr := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		x, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		xarr = append(xarr, x)

		y, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}

		yarr = append(yarr, y)
	}

	m := make(map[int]int)

	for _, xVal := range xarr {
		count := 0
		for _, yVal := range yarr {
			if xVal == yVal {
				count += 1
			}
		}
		m[xVal] = count
	}

	for key, val := range m {
		res := key * val
		total += res
	}

	fmt.Println("D01 P2: ", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
