package d09

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func EffDecodeData(data string) [][]string {
	var res [][]string
	id := 0
	for i, char := range data {
		var currBlock []string
		isFile := false
		currChar := "."
		if i%2 == 0 {
			isFile = true
			currChar = strconv.Itoa(id)
		}
		count, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < count; i++ {
			currBlock = append(currBlock, currChar)
		}
		if isFile {
			id += 1
		}
		res = append(res, currBlock)
	}
	return res
}

func EffOrderBlocks(blocks [][]string) [][]string {
	currBlocks := make([][]string, len(blocks))
	for i := range blocks {
		currBlocks[i] = make([]string, len(blocks[i]))
		copy(currBlocks[i], blocks[i])
	}
	currRPtr := len(currBlocks) - 1
	for currRPtr >= 0 {
		if len(currBlocks[currRPtr]) <= 0 || currBlocks[currRPtr][0] == "." {
			currRPtr -= 1
			continue
		}

		lPtr := 0
		currNumLen := len(currBlocks[currRPtr])

		for lPtr < currRPtr {
			if len(currBlocks[lPtr]) <= 0 {
				lPtr += 1
				continue
			}

			if currBlocks[lPtr][0] != "." {
				lPtr += 1
				continue
			}

			if len(currBlocks[lPtr]) < currNumLen {
				lPtr += 1
				continue
			}

			break
		}

		if lPtr >= currRPtr {
			currRPtr -= 1
			continue
		}

		// fmt.Printf("LEFT: %v\n", currBlocks[lPtr])
		// fmt.Printf("RIGHT: %v\n", currBlocks[currRPtr])
		// fmt.Printf("BEFORE: %v\n", currBlocks)
		if len(currBlocks[lPtr]) == currNumLen {
			currBlocks[lPtr], currBlocks[currRPtr] = currBlocks[currRPtr], currBlocks[lPtr]
		} else {
			currNum := currBlocks[currRPtr][0]
			var numBlock []string
			var blankBlock []string
			var replaceBlock []string
			numCount := len(currBlocks[currRPtr])
			blankCount := len(currBlocks[lPtr]) - len(currBlocks[currRPtr])
			for numCount > 0 {
				numBlock = append(numBlock, currNum)
				replaceBlock = append(replaceBlock, ".")
				numCount -= 1
			}
			for blankCount > 0 {
				blankBlock = append(blankBlock, ".")
				blankCount -= 1
			}
			currBlocks[currRPtr] = replaceBlock
			currBlocks[lPtr] = blankBlock
			newBlocks := make([][]string, len(currBlocks)+1)
			copy(newBlocks, currBlocks[:lPtr])
			newBlocks[lPtr] = numBlock
			copy(newBlocks[lPtr+1:], currBlocks[lPtr:])
			currBlocks = newBlocks
		}
		// fmt.Printf("AFTER: %v\n", currBlocks)
		// fmt.Println()

		currRPtr -= 1
	}

	return currBlocks
}

func EffGetChecksum(data [][]string) *big.Int {
	var parsedData []string
	for _, block := range data {
		for _, val := range block {
			parsedData = append(parsedData, val)
		}
	}
	res := big.NewInt(0)
	for i, char := range parsedData {
		if char == "." {
			continue
		}
		num, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		bigNum := big.NewInt(int64(num))
		bigNum.Mul(bigNum, big.NewInt(int64(i)))
		res.Add(res, bigNum)
	}
	return res
}

func P2() {
	file, err := os.Open("d09/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := big.NewInt(0)

	var data string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = scanner.Text()
	}

	decodedData := EffDecodeData(data)
	decodedData = EffOrderBlocks(decodedData)
	total = EffGetChecksum(decodedData)

	fmt.Println("D09 P2: ", total)
}
