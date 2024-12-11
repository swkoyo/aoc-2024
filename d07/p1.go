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

type Node struct {
	operator string
	left     *Node
	right    *Node
	value    *big.Int
}

func ValExists(root *Node, val *big.Int) bool {
	if root == nil {
		return false
	}

	if root.left == nil && root.right == nil {
		return root.value.Cmp(val) == 0
	}

	return ValExists(root.left, val) || ValExists(root.right, val)
}

func NewTree(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	return &Node{
		operator: "",
		left:     nil,
		right:    nil,
		value:    big.NewInt(int64(values[0])),
	}
}

func GenerateTree(root *Node, remainingValues []int) {
	if len(remainingValues) == 0 {
		return
	}

	nextVal := big.NewInt(int64(remainingValues[0]))

	addNode := &Node{
		operator: "+",
		left:     nil,
		right:    nil,
		value:    new(big.Int).Add(root.value, nextVal),
	}
	root.left = addNode
	GenerateTree(addNode, remainingValues[1:])

	multiNode := &Node{
		operator: "*",
		left:     nil,
		right:    nil,
		value:    new(big.Int).Mul(root.value, nextVal),
	}
	root.right = multiNode
	GenerateTree(multiNode, remainingValues[1:])
}

func P1() {
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
		root := NewTree(nums)
		GenerateTree(root, nums[1:])
		if ValExists(root, key) {
			total.Add(total, key)
		}
	}

	fmt.Println("D07 P1: ", total)
}
