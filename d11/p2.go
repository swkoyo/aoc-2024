package d11

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func (s *Stone) BlinkStone() []Stone {
	res := make([]Stone, 0)
	if s.val.Cmp(big.NewInt(int64(0))) == 0 {
		s.val = big.NewInt(int64(1))
		res = append(res, *s)
	} else if !s.IsEven() {
		s.val.Mul(s.val, big.NewInt(int64(2024)))
		return append(res, *s)
	} else {
		lStone, rStone := s.Split()
		res = append(res, lStone)
		res = append(res, rStone)
	}

	return res
}

func GetStoneBlinkLen(stone Stone, blinks int, cache map[string]*big.Int) *big.Int {
	if blinks == 0 {
		return big.NewInt(int64(1))
	}

	blinkedStones := stone.BlinkStone()
	total := big.NewInt(int64(0))

	for _, stone := range blinkedStones {
		cacheStr := fmt.Sprintf("%s:%d", stone.ToString(), blinks-1)
		if cachedLen, exists := cache[cacheStr]; exists {
			total.Add(total, cachedLen)
		} else {
			cachedLen := GetStoneBlinkLen(stone, blinks-1, cache)
			cache[cacheStr] = cachedLen
			total.Add(total, cachedLen)
		}
	}

	return total
}

func P2() {
	file, err := os.Open("d11/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data []Stone

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, val := range strings.Split(line, " ") {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, Stone{val: big.NewInt(int64(num))})
		}
	}

	total := big.NewInt(int64(0))
	cache := make(map[string]*big.Int)

	for _, stone := range data {
		total.Add(total, GetStoneBlinkLen(stone, 75, cache))
	}

	fmt.Println("D11 P2: ", total)
}
