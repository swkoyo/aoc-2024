package d07

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type TNode struct {
	operator string
	left     *TNode
	center   *TNode
	right    *TNode
	value    *big.Int
}

func TValExists(root *TNode, val *big.Int) bool {
	if root == nil {
		return false
	}

	if root.value.Cmp(val) > 0 {
		return false
	}

	if root.left == nil && root.center == nil && root.right == nil {
		return root.value.Cmp(val) == 0
	}

	return TValExists(root.left, val) || TValExists(root.center, val) || TValExists(root.right, val)
}

func TNewTree(values []int) *TNode {
	if len(values) == 0 {
		return nil
	}

	return &TNode{
		operator: "",
		left:     nil,
		center:   nil,
		right:    nil,
		value:    big.NewInt(int64(values[0])),
	}
}

func TGenerateTree(root *TNode, remainingValues []int) {
	if len(remainingValues) == 0 {
		return
	}

	nextVal := big.NewInt(int64(remainingValues[0]))

	addNode := &TNode{
		operator: "+",
		left:     nil,
		center:   nil,
		right:    nil,
		value:    new(big.Int).Add(root.value, nextVal),
	}
	root.left = addNode
	TGenerateTree(addNode, remainingValues[1:])

	concatIntStr := root.value.String() + strconv.Itoa(remainingValues[0])
	num, err := strconv.Atoi(concatIntStr)
	if err != nil {
		log.Fatal(err)
	}
	bigNum := big.NewInt(int64(num))
	concatNode := &TNode{
		operator: "||",
		left:     nil,
		center:   nil,
		right:    nil,
		value:    bigNum,
	}
	root.center = concatNode
	TGenerateTree(concatNode, remainingValues[1:])

	multiNode := &TNode{
		operator: "*",
		left:     nil,
		center:   nil,
		right:    nil,
		value:    new(big.Int).Mul(root.value, nextVal),
	}
	root.right = multiNode
	TGenerateTree(multiNode, remainingValues[1:])
}

func P2() {
	file, err := os.Open("d07/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := big.NewInt(0)

	data := make(map[*big.Int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slices := strings.Split(line, ":")
		nums := strings.Split(slices[1], " ")

		res, err := strconv.Atoi(slices[0])
		if err != nil {
			log.Fatal(err)
		}

		key := big.NewInt(int64(res))
		data[key] = make([]int, 0)
		for _, num := range nums[1:] {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			data[key] = append(data[key], n)
		}
	}

	for key, nums := range data {
		root := TNewTree(nums)
		TGenerateTree(root, nums[1:])
		if TValExists(root, key) {
			total.Add(total, key)
		}
	}

	fmt.Println("D07 P2: ", total)
}
