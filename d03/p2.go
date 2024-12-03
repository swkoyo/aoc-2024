package d03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func P2() {
	file, err := os.Open("d03/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)

	var txt string

	for scanner.Scan() {
		txt += scanner.Text()
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(txt, -1)

	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			total += (x * y)
		}
	}

	fmt.Println("D03 P2: ", total)
}
