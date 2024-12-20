package day20

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day20",
	Short: "day20",
	Long:  "day20",
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

	result := Part1(string(content), 100)
	fmt.Println("Part 1 result: ", result)

	content, err = os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input2.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result = Part2(string(content), 100)
	fmt.Println("Part 2 result: ", result)
}

type maze struct {
	grid  map[image.Point]bool
	start image.Point
}

func newMaze(lines []string) maze {
	grid := make(map[image.Point]bool)
	var start image.Point
	for y, line := range lines {
		for x, p := range line {
			if p != '#' {
				point := image.Point{x, y}
				grid[point] = true
				if p == 'S' {
					start = point
				}
			}
		}
	}
	return maze{grid, start}
}

func (m maze) countCheats(save, duration int) int {
	dist := make(map[image.Point]int)
	dist[m.start] = 0
	queue := []image.Point{m.start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, p := range []image.Point{
			{X: pos.X - 1, Y: pos.Y},
			{X: pos.X + 1, Y: pos.Y},
			{X: pos.X, Y: pos.Y - 1},
			{X: pos.X, Y: pos.Y + 1},
		} {
			if m.grid[p] && dist[p] == 0 {
				dist[p] = dist[pos] + 1
				queue = append(queue, p)
			}
		}
	}

	var cheatCount int

	for point1, dist1 := range dist {
		for point2, dist2 := range dist {
			if point1 == point2 {
				continue
			}
			manhattan := int(math.Abs(float64(point1.X-point2.X)) + math.Abs(float64(point1.Y-point2.Y)))
			if manhattan <= duration && (dist2-dist1-manhattan) >= save {
				cheatCount++
			}
		}
	}

	return cheatCount
}

func Part1(input string, save int) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	m := newMaze(lines)
	return m.countCheats(save, 2)
}

func Part2(input string, save int) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	m := newMaze(lines)
	return m.countCheats(save, 20) + 1
}
