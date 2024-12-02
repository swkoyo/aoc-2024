package d01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func P1() {
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

	sort.Ints(xarr[:])
	sort.Ints(yarr[:])

	for i := 0; i < len(xarr); i++ {
		x := xarr[i]
		y := yarr[i]
		diff := 0

		if x > y {
			diff = x - y
		} else {
			diff = y - x
		}

		total += diff
	}

	fmt.Println("D01 P1: ", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
