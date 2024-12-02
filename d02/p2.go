package d02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func P2() {
	file, err := os.Open("d02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := make([]int, 0)
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				log.Fatal("Failed to convert word to int", word)
			}
			arr = append(arr, num)
		}
		idx := 0
		safe := isSafe(arr)

		for !safe && idx < len(arr) {
			newArr := make([]int, 0, len(arr)-1)
			newArr = append(newArr, arr[:idx]...)
			newArr = append(newArr, arr[idx+1:]...)
			safe = isSafe(newArr)
			idx += 1
		}

		if safe {
			total += 1
		}
	}

	fmt.Println("D02 P2: ", total)
}
