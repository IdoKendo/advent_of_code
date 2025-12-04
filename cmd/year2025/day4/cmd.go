package day4

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  "day4",
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

var directions = [8]image.Point{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Grid struct {
	grid           map[image.Point]bool
	height, length int
}

func newGrid(input string) Grid {
	grid := make(map[image.Point]bool)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, line := range lines {
		for x, val := range line {
			point := image.Point{x, y}
			grid[point] = val == '@'
		}
	}

	return Grid{
		grid:   grid,
		height: len(lines),
		length: len(lines[0]),
	}
}

func (g Grid) AccessibleRolls() []image.Point {
	accessibleRolls := []image.Point{}
	for p, v := range g.grid {
		if !v {
			continue
		}
		surroundingRolls := 0
		for _, direction := range directions {
			position := p.Add(direction)
			if position.X > g.length || position.Y > g.height || !g.grid[position] {
				continue
			}
			surroundingRolls++
		}
		if surroundingRolls < 4 {
			accessibleRolls = append(accessibleRolls, p)
		}
	}
	return accessibleRolls
}

func (g *Grid) RemoveRolls(rolls []image.Point) {
	for _, roll := range rolls {
		g.grid[roll] = false
	}
}

func Part1(input string) int {
	grid := newGrid(input)

	return len(grid.AccessibleRolls())
}

func Part2(input string) int {
	grid := newGrid(input)
	removedRolls := 0
	accessibleRolls := grid.AccessibleRolls()
	for len(accessibleRolls) > 0 {
		removedRolls += len(accessibleRolls)
		grid.RemoveRolls(accessibleRolls)
		accessibleRolls = grid.AccessibleRolls()
	}
	return removedRolls
}
