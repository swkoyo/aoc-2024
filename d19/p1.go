package d19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidDesign(design string, towels []string, cache *map[string]bool) bool {
	if len(design) == 0 {
		return true
	}

	if res, exists := (*cache)[design]; exists {
		return res
	}

	for i := 1; i < len(design)+1; i++ {
		exists := false
		for _, towel := range towels {
			if towel == design[:i] {
				exists = true
				break
			}
		}
		if exists && isValidDesign(design[i:], towels, cache) {
			(*cache)[design] = true
			return true
		}
	}

	(*cache)[design] = false
	return false
}

func P1() {
	file, err := os.Open("d19/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	towels := make([]string, 0)
	designs := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if strings.Contains(line, ",") {
				t := strings.Split(line, ", ")
				towels = append(towels, t...)
			} else {
				designs = append(designs, line)
			}
		}
	}

	cache := make(map[string]bool)

	for _, design := range designs {
		if isValidDesign(design, towels, &cache) {
			total += 1
		}
	}

	fmt.Println("D19 P1: ", total)
}
