package d05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swapInvalidPair(pages []int, rules [][]int) bool {
	idx := makePagesIdxMap(pages, rules)

	for _, rule := range rules {
		aIdx, aExists := idx[rule[0]]
		bIdx, bExists := idx[rule[1]]

		if aExists && bExists && aIdx > bIdx {
			pages[aIdx], pages[bIdx] = pages[bIdx], pages[aIdx]
			return true
		}
	}

	return false
}

func P2() {
	file, err := os.Open("d05/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	var rules [][]int
	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			slices := strings.Split(line, "|")
			x, err := strconv.Atoi(slices[0])
			if err != nil {
				log.Fatal("Failed to convert string to int")
			}

			y, err := strconv.Atoi(slices[1])
			if err != nil {
				log.Fatal("Failed to convert string to int")
			}

			rules = append(rules, []int{x, y})
		} else if strings.Contains(line, ",") {
			slices := strings.Split(line, ",")
			var nums []int
			for _, val := range slices {
				num, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal("Failed to convert string to int")
				}
				nums = append(nums, num)
			}
			data = append(data, nums)
		}
	}

	for _, pages := range data {
		if !isValidOrder(pages, rules) {
			isInvalid := true
			for isInvalid {
				isInvalid = swapInvalidPair(pages, rules)
			}
			total += getMidNum(pages)
		}
	}

	fmt.Println("D05 P2: ", total)
}
