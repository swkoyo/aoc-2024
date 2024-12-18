package d13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_USES = 100

type Machine struct {
	PX float64
	PY float64
	AX float64
	AY float64
	BX float64
	BY float64
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
				currMachine = &Machine{}
			}
			parts := strings.Fields(line)
			if parts[0] == "Prize:" {
				xStr := strings.TrimSuffix(strings.TrimPrefix(parts[1], "X="), ",")
				yStr := strings.TrimSuffix(strings.TrimPrefix(parts[2], "Y="), ",")
				x, _ := strconv.ParseFloat(xStr, 64)
				y, _ := strconv.ParseFloat(yStr, 64)
				currMachine.PX = x
				currMachine.PY = y
				machines = append(machines, currMachine)
				currMachine = nil
			} else {
				xStr := strings.TrimSuffix(strings.TrimPrefix(parts[2], "X+"), ",")
				yStr := strings.TrimSuffix(strings.TrimPrefix(parts[3], "Y+"), ",")
				x, _ := strconv.ParseFloat(xStr, 64)
				y, _ := strconv.ParseFloat(yStr, 64)
				if strings.Contains(parts[1], "A") {
					currMachine.AX = x
					currMachine.AY = y
				} else {
					currMachine.BX = x
					currMachine.BY = y
				}
			}
		}
	}

	for _, machine := range machines {
		minTokens := 0
		for i := 0; i <= MAX_USES; i++ {
			for j := 0; j < MAX_USES; j++ {
				if ((machine.AX*float64(i))+(machine.BX*float64(j)) == machine.PX) && ((machine.AY*float64(i))+(machine.BY*float64(j)) == machine.PY) {
					tokens := (i * 3) + (j * 1)
					if minTokens == 0 || tokens < minTokens {
						minTokens = tokens
					}
				}
			}
		}
		total += minTokens
	}

	fmt.Println("D13 P1: ", total)
}
