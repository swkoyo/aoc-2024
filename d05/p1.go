package d05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func P1() {
	file, err := os.Open("d05/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	rules := make(map[int][]int)
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

			if _, exists := rules[x]; !exists {
				rules[x] = make([]int, 0)
			}
			rules[x] = append(rules[x], y)
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
		isValid := true
		for prevPage, nextPages := range rules {
			if !isValid {
				break
			}
			prevPageIdx := -1
			for idx, currPage := range pages {
				if currPage == prevPage {
					prevPageIdx = idx
					break
				}
			}
			if prevPageIdx > -1 {
				for _, nextPage := range nextPages {
					nextPageIdx := -1
					for idx, currPage := range pages {
						if currPage == nextPage {
							nextPageIdx = idx
							break
						}
					}
					if nextPageIdx > -1 && prevPageIdx > nextPageIdx {
						isValid = false
						break
					}
				}
			}
		}
		if isValid {
			midIdx := len(pages) / 2
			mid := pages[midIdx]
			total += mid
		}
	}

	fmt.Println("D05 P1: ", total)
}
