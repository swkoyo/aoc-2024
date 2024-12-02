package d02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isInc(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func isDec(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			return false
		}
	}
	return true
}

func isSafe(arr []int) bool {
	if !isInc(arr) && !isDec(arr) {
		return false
	}

	for i := 1; i < len(arr); i++ {
		diff := 0
		if arr[i-1] < arr[i] {
			diff = arr[i] - arr[i-1]
		} else {
			diff = arr[i-1] - arr[i]
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func P1() {
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
		if isSafe(arr) {
			total += 1
		}
	}

	fmt.Println("D02 P1: ", total)
}
