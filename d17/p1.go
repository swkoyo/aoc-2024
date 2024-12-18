package d17

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	A       int
	B       int
	C       int
	Output  []string
	Pointer int
	Program []int
}

func (c Computer) OutputToString() string {
	return strings.Join(c.Output, ",")
}

func (c Computer) Print() {
	fmt.Println("#####################################")
	fmt.Printf("PROGRAM: %v\n", c.Program)
	fmt.Println()
	fmt.Println("REGISTERS")
	fmt.Printf("A: %d\n", c.A)
	fmt.Printf("B: %d\n", c.B)
	fmt.Printf("C: %d\n", c.C)
	fmt.Println()
	fmt.Printf("POINTER: %d\n", c.Pointer)
	fmt.Println()
	fmt.Printf("OUTPUT: %v\n", c.Output)
	fmt.Println("#####################################")
}

func (c *Computer) IsDone() bool {
	return c.Pointer >= len(c.Program)
}

func (c *Computer) Run() {
	for !c.IsDone() {
		c.RunInstruction()
	}
}

func (c *Computer) RunInstruction() {
	opcode := c.Program[c.Pointer]
	operand := c.Program[c.Pointer+1]

	if opcode == 0 {
		c.Adv(c.GetComboOperand(operand))
	} else if opcode == 1 {
		c.Bxl(operand)
	} else if opcode == 2 {
		c.Bst(c.GetComboOperand(operand))
	} else if opcode == 3 {
		c.Jnz(operand)
	} else if opcode == 4 {
		c.Bxc()
	} else if opcode == 5 {
		c.Out(c.GetComboOperand(operand))
	} else if opcode == 6 {
		c.Bdv(c.GetComboOperand(operand))
	} else if opcode == 7 {
		c.Cdv(c.GetComboOperand(operand))
	}
}

func (c *Computer) GetComboOperand(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}
	if operand == 4 {
		return c.A
	}
	if operand == 5 {
		return c.B
	}
	if operand == 6 {
		return c.C
	}
	return -1
}

func (c *Computer) Adv(operand int) {
	c.A >>= operand
	c.Pointer += 2
}

func (c *Computer) Bxl(operand int) {
	c.B ^= operand
	c.Pointer += 2
}

func (c *Computer) Bst(operand int) {
	c.B = operand % 8
	c.Pointer += 2
}

func (c *Computer) Jnz(operand int) {
	if c.A == 0 {
		c.Pointer += 2
		return
	}
	c.Pointer = operand
}

func (c *Computer) Bxc() {
	c.B ^= c.C
	c.Pointer += 2
}

func (c *Computer) Out(operand int) {
	res := operand % 8
	c.Output = append(c.Output, strconv.Itoa(res))
	c.Pointer += 2
}

func (c *Computer) Bdv(operand int) {
	c.B = c.A >> operand
	c.Pointer += 2
}

func (c *Computer) Cdv(operand int) {
	c.C = c.A >> operand
	c.Pointer += 2
}

func P1() {
	file, err := os.Open("d17/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	computer := Computer{Output: make([]string, 0), Pointer: 0, Program: make([]int, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			parts := strings.Fields(line)
			if parts[0] == "Register" {
				reg := parts[1][0]
				val, _ := strconv.Atoi(parts[2])
				if reg == 'A' {
					computer.A = val
				} else if reg == 'B' {
					computer.B = val
				} else {
					computer.C = val
				}
			} else if parts[0] == "Program:" {
				for _, char := range strings.Split(parts[1], ",") {
					val, _ := strconv.Atoi(char)
					computer.Program = append(computer.Program, val)
				}
			}
		}
	}

	computer.Run()
	computer.Print()

	fmt.Println("D17 P1: ", computer.OutputToString())
}
