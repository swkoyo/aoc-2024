package d19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validDesignCount(design string, towels []string, cache *map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	if res, exists := (*cache)[design]; exists {
		return res
	}

	count := 0

	for i := 1; i < len(design)+1; i++ {
		exists := false
		for _, towel := range towels {
			if towel == design[:i] {
				exists = true
				break
			}
		}
		if exists {
			count += validDesignCount(design[i:], towels, cache)
		}
	}

	(*cache)[design] = count
	return count
}

func P2() {
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

	cache := make(map[string]int)

	for _, design := range designs {
		total += validDesignCount(design, towels, &cache)
	}

	fmt.Println("D19 P2: ", total)
}
