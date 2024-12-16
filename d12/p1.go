package d12

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	row int
	col int
}

func FindPlot(originPlant rune, plot Coordinate, currentRegion map[Coordinate]bool, data []string) int {
	if plot.row < 0 || plot.row >= len(data) || plot.col < 0 || plot.col >= len(data[0]) {
		return 1
	}

	if _, exists := currentRegion[plot]; exists {
		return 0
	}

	if data[plot.row][plot.col] != byte(originPlant) {
		return 1
	}

	currentRegion[plot] = true
	upPlot := Coordinate{row: plot.row - 1, col: plot.col}
	rightPlot := Coordinate{row: plot.row, col: plot.col + 1}
	downPlot := Coordinate{row: plot.row + 1, col: plot.col}
	leftPlot := Coordinate{row: plot.row, col: plot.col - 1}

	return FindPlot(originPlant, upPlot, currentRegion, data) + FindPlot(originPlant, rightPlot, currentRegion, data) + FindPlot(originPlant, downPlot, currentRegion, data) + FindPlot(originPlant, leftPlot, currentRegion, data)
}

func PlotIsUsed(plot Coordinate, plots []map[Coordinate]bool) bool {
	for _, p := range plots {
		if _, exists := p[plot]; exists {
			return true
		}
	}
	return false
}

func P1() {
	file, err := os.Open("d12/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	regions := make([]map[Coordinate]bool, 0)
	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for row, currRow := range data {
		for col, char := range currRow {
			plot := Coordinate{row: row, col: col}
			if !PlotIsUsed(plot, regions) {
				currPlotRegion := make(map[Coordinate]bool)
				perimeter := FindPlot(char, plot, currPlotRegion, data)
				area := len(currPlotRegion)
				regions = append(regions, currPlotRegion)
				total += (area * perimeter)
			}
		}
	}

	fmt.Println("D12 P1: ", total)
}
