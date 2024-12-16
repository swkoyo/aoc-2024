package d13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_USES = 100

type Location struct {
	X int
	Y int
}

func NewLocation(x, y int) Location {
	return Location{
		x,
		y,
	}
}

type Button struct {
	X    int
	Y    int
	Cost int
}

func NewButton(x, y, cost int) Button {
	return Button{
		x,
		y,
		cost,
	}
}

type Machine struct {
	CurrentLoc Location
	PrizeLoc   Location
	ABtn       Button
	BBtn       Button
}

func NewMachine() *Machine {
	return &Machine{
		CurrentLoc: NewLocation(0, 0),
	}
}

func FindMinTokensHelper(xLeft, yLeft, aUses, bUses int, ABtn, BBtn Button, checked, successful map[string]bool) {
	if _, exists := checked[fmt.Sprintf("%d:%d", xLeft, yLeft)]; exists {
		return
	}

	checked[fmt.Sprintf("%d:%d", xLeft, yLeft)] = true

	if xLeft == 0 && yLeft == 0 {
		successful[fmt.Sprintf("%d:%d", aUses, bUses)] = true
		return
	}

	if aUses >= MAX_USES || bUses >= MAX_USES {
		return
	}

	if xLeft < 0 || yLeft < 0 {
		return
	}

	FindMinTokensHelper(xLeft-ABtn.X, yLeft-ABtn.Y, aUses+1, bUses, ABtn, BBtn, checked, successful)
	FindMinTokensHelper(xLeft-BBtn.X, yLeft-BBtn.Y, aUses, bUses+1, ABtn, BBtn, checked, successful)

	return
}

func FindMinTokens(machine Machine) int {
	successful := make(map[string]bool)
	checked := make(map[string]bool)
	FindMinTokensHelper(machine.PrizeLoc.X, machine.PrizeLoc.Y, 0, 0, machine.ABtn, machine.BBtn, checked, successful)
	total := 0
	for uses, _ := range successful {
		parts := strings.Split(uses, ":")
		aUses, _ := strconv.Atoi(parts[0])
		bUses, _ := strconv.Atoi(parts[1])
		cost := (machine.ABtn.Cost * aUses) + (machine.BBtn.Cost * bUses)
		if total == 0 || cost < total {
			total = cost
		}

	}
	return total
}

func P1() {
	file, err := os.Open("d13/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	machines := make([]*Machine, 0)

	var currMachine *Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if currMachine == nil {
				currMachine = NewMachine()
			}
			parts := strings.Fields(line)
			if parts[0] == "Prize:" {
				xStr := strings.TrimSuffix(strings.TrimPrefix(parts[1], "X="), ",")
				yStr := strings.TrimSuffix(strings.TrimPrefix(parts[2], "Y="), ",")
				x, _ := strconv.Atoi(xStr)
				y, _ := strconv.Atoi(yStr)
				currMachine.PrizeLoc = NewLocation(x, y)
				machines = append(machines, currMachine)
				currMachine = nil
			} else {
				xStr := strings.TrimSuffix(strings.TrimPrefix(parts[2], "X+"), ",")
				yStr := strings.TrimSuffix(strings.TrimPrefix(parts[3], "Y+"), ",")
				x, _ := strconv.Atoi(xStr)
				y, _ := strconv.Atoi(yStr)
				if strings.Contains(parts[1], "A") {
					currMachine.ABtn = NewButton(x, y, 3)
				} else {
					currMachine.BBtn = NewButton(x, y, 1)
				}
			}
		}
	}

	for _, machine := range machines {
		total += FindMinTokens(*machine)
	}

	fmt.Println("D13 P1: ", total)
}
