package d13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isWhole(f float64) bool {
	return math.Floor(f) == f
}

func P2() {
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
				currMachine.PX = x + 10000000000000
				currMachine.PY = y + 10000000000000
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
		ca := (machine.PX*machine.BY - machine.PY*machine.BX) / (machine.AX*machine.BY - machine.AY*machine.BX)
		cb := (machine.PX - machine.AX*ca) / machine.BX
		if isWhole(ca) && isWhole(cb) {
			total += (int(ca) * 3) + (int(cb) * 1)
		}
	}

	fmt.Println("D13 P2: ", total)
}
