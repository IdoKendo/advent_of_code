package day11

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day11",
	Short: "day11",
	Long:  "day11",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())

	},
}

func execute(parent, command string) {
	content, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input1.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result := Part1(string(content))
	fmt.Println("Part 1 result: ", result)

	content, err = os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input2.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result = Part2(string(content))
	fmt.Println("Part 2 result: ", result)
}

type galaxy struct {
	x, y int
}

type GalaxyMap struct {
	expandedRows        []int
	expandedCols        []int
	expansionMultiplier int
	galaxies            []galaxy
}

func newGalaxyMap(input string, expansionMultiplier int) GalaxyMap {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	gm := GalaxyMap{
		expandedRows:        []int{},
		expandedCols:        []int{},
		expansionMultiplier: expansionMultiplier,
		galaxies:            []galaxy{},
	}

	for y, line := range lines {
		if !strings.Contains(line, "#") {
			gm.expandedRows = append(gm.expandedRows, y)
		}
	}

	for x := 0; x < len(lines[0]); x++ {
		isEmpty := true
		for _, line := range lines {
			if line[x:x+1] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			gm.expandedCols = append(gm.expandedCols, x)
		}
	}

	for y, line := range lines {
		for x, p := range line {
			if p == '#' {
				gm.galaxies = append(gm.galaxies, galaxy{x, y})
			}
		}
	}

	return gm
}

func (gm *GalaxyMap) distanceBetween(a, b galaxy) int {
	minX, maxX := min(a.x, b.x), max(a.x, b.x)
	minY, maxY := min(a.y, b.y), max(a.y, b.y)
	expandedRowCount := 0
	for _, row := range gm.expandedRows {
		if row > minY && row < maxY {
			expandedRowCount++
		}
	}

	expandedColCount := 0
	for _, col := range gm.expandedCols {
		if col > minX && col < maxX {
			expandedColCount++
		}
	}

	expansionDistance := (expandedRowCount + expandedColCount) * (gm.expansionMultiplier - 1)
	return (maxX - minX) + (maxY - minY) + expansionDistance
}

func (gm *GalaxyMap) shortestPaths() int {
	sum := 0
	for i, a := range gm.galaxies {
		for _, b := range gm.galaxies[i+1:] {
			length := gm.distanceBetween(a, b)
			sum += length
		}
	}

	return sum
}

func Part1(input string) int {
	galaxyMap := newGalaxyMap(input, 2)
	return galaxyMap.shortestPaths()
}

func Example10(input string) int {
	galaxyMap := newGalaxyMap(input, 10)
	return galaxyMap.shortestPaths()
}

func Example100(input string) int {
	galaxyMap := newGalaxyMap(input, 100)
	return galaxyMap.shortestPaths()
}

func Part2(input string) int {
	galaxyMap := newGalaxyMap(input, 1000000)
	return galaxyMap.shortestPaths()
}
