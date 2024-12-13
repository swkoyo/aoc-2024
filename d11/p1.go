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

type Stone struct {
	val *big.Int
}

func (s *Stone) Length() int {
	return len(s.ToString())
}

func (s *Stone) ToString() string {
	return s.val.String()
}

func (s *Stone) IsEven() bool {
	return s.Length()%2 == 0
}

func (s *Stone) Split() (Stone, Stone) {
	sliceLen := s.Length() / 2
	stoneStr := s.ToString()
	x := stoneStr[:sliceLen]
	y := stoneStr[sliceLen:]
	xNum := new(big.Int)
	yNum := new(big.Int)
	_, success := xNum.SetString(x, 10)
	if !success {
		log.Fatal("Failed to convert")
	}
	_, success = yNum.SetString(y, 10)
	if !success {
		log.Fatal("Failed to convert")
	}
	return Stone{val: xNum}, Stone{val: yNum}
}

func Blink(stones []Stone) []Stone {
	i := 0
	stonesLen := len(stones)

	for i < stonesLen {
		currStone := stones[i]
		if currStone.val.Cmp(big.NewInt(int64(0))) == 0 {
			stones[i].val = big.NewInt(int64(1))
		} else if !currStone.IsEven() {
			stones[i].val.Mul(stones[i].val, big.NewInt(int64(2024)))
		} else {
			lStone, rStone := currStone.Split()
			newStones := make([]Stone, len(stones)+1)
			copy(newStones, stones[:i])
			newStones[i] = lStone
			newStones[i+1] = rStone
			copy(newStones[i+2:], stones[i+1:])
			stones = newStones
			stonesLen = len(stones)
			i += 1
		}
		i += 1
	}

	return stones
}

func P1() {
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

	for i := 25; i > 0; i-- {
		data = Blink(data)
	}

	fmt.Println("D11 P1: ", len(data))
}
