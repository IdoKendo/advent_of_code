package day18

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day18",
	Short: "day18",
	Long:  "day18",
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

	result := Part1(string(content), 71, 1024)
	fmt.Println("Part 1 result: ", result)

	content, err = os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input2.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result = Part2(string(content), 71, 1024)
	fmt.Println("Part 2 result: ", result)
}

var directions = []image.Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type memoryMaze struct {
	start       image.Point
	corruptions []image.Point
	end         image.Point
}

func newMemoryMaze(lines []string, gridSize int) memoryMaze {
	corruptions := make([]image.Point, len(lines))
	for i, line := range lines {
		s := strings.Split(line, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		corruptions[i] = image.Point{x, y}
	}

	return memoryMaze{
		start:       image.Point{0, 0},
		corruptions: corruptions,
		end:         image.Point{gridSize - 1, gridSize - 1},
	}
}

func (m memoryMaze) print(gridSize int, curr image.Point) {
	walls := make(map[image.Point]bool)
	for _, corruption := range m.corruptions {
		walls[corruption] = true
	}

	for y := range gridSize {
		for x := range gridSize {
			p := image.Point{x, y}
			if curr.Eq(p) {
				fmt.Print("@")
			} else if m.start.Eq(p) {
				fmt.Print("S")
			} else if m.end.Eq(p) {
				fmt.Print("E")
			} else if walls[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *memoryMaze) countSteps(gridSize int) (int, bool) {
	queue := []image.Point{m.start}
	visited := make(map[image.Point]bool)
	visited[m.start] = true

	for _, corruption := range m.corruptions {
		visited[corruption] = true
	}

	parent := make(map[image.Point]image.Point)
	distances := make(map[image.Point]int)
	distances[m.start] = 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == m.end {
			path := []image.Point{}
			current := m.end
			for current != m.start {
				path = append([]image.Point{current}, path...)
				current = parent[current]
			}
			path = append([]image.Point{m.start}, path...)
			return len(path) - 1, true
		}

		for _, dir := range directions {
			next := p.Add(dir)

			if next.X >= 0 && next.X < gridSize && next.Y >= 0 && next.Y < gridSize && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				parent[next] = p
				distances[next] = distances[p] + 1
			}
		}
	}

	return -1, false
}

func Part1(input string, gridSize int, bytesFallen int) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	lines = lines[:bytesFallen]
	memoryMaze := newMemoryMaze(lines, gridSize)
	result, _ := memoryMaze.countSteps(gridSize)
	return strconv.Itoa(result)
}

func Part2(input string, gridSize int, bytesFallen int) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	result := ""
	for i := bytesFallen; i < len(lines); i++ {
		memoryMaze := newMemoryMaze(lines[:i], gridSize)
		_, ok := memoryMaze.countSteps(gridSize)
		if !ok {
			result = lines[i-1]
			break
		}

	}
	return result
}
