package d09

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func DecodeData(data string) []string {
	res := make([]string, 0)
	id := 0
	for i, char := range data {
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
			res = append(res, currChar)
		}
		if isFile {
			id += 1
		}
	}
	return res
}

func OrderBlocks(blocks []string) []string {
	lPtr := 0
	rPtr := len(blocks) - 1

	for lPtr < rPtr {
		for blocks[lPtr] != "." {
			lPtr += 1
		}

		for blocks[rPtr] == "." {
			rPtr -= 1
		}

		if lPtr >= rPtr {
			break
		}

		blocks[lPtr], blocks[rPtr] = blocks[rPtr], blocks[lPtr]

		lPtr += 1
		rPtr -= 1
	}

	return blocks
}

func GetChecksum(data []string) *big.Int {
	res := big.NewInt(0)
	for i, char := range data {
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

func P1() {
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

	decodedData := DecodeData(data)
	decodedData = OrderBlocks(decodedData)
	total = GetChecksum(decodedData)

	fmt.Println("D09 P1: ", total)
}
