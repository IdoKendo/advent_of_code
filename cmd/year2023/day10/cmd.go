package day10

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day10",
	Short: "day10",
	Long:  "day10",
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

func createMap(lines []string) [][]string {
	height := len(lines)
	width := len(lines[0])
	m := make([][]string, height)

	for i, line := range lines {
		m[i] = make([]string, width)
		tiles := strings.Split(line, "")
		copy(m[i], tiles)
	}
	return m
}

func findStart(m [][]string) [2]int {
	if m[0][0] == "S" {
		return [2]int{0, 0}
	}
	var startPos [2]int
	for i, tiles := range m {
		for j, tile := range tiles {
			if tile == "S" {
				startPos = [2]int{j, i}
				break
			}
		}
		if startPos != [2]int{0, 0} {
			break
		}
	}
	return startPos
}

type direction [2]int

var (
	north     direction = [2]int{0, -1}
	east      direction = [2]int{1, 0}
	south     direction = [2]int{0, 1}
	west      direction = [2]int{-1, 0}
	northEast direction = [2]int{1, -1}
	northWest direction = [2]int{-1, -1}
	southEast direction = [2]int{1, 1}
	southWest direction = [2]int{-1, 1}
)

var pipes = map[string][]direction{
	"|": {north, south},
	"-": {east, west},
	"L": {north, east},
	"J": {north, west},
	"7": {south, west},
	"F": {south, east},
	"S": {north, west, south, east},
}

func traverseMap(m [][]string, start [2]int) map[[2]int]bool {
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		v := m[node[1]][node[0]]
		for _, dir := range pipes[v] {
			next := [2]int{node[0] + dir[0], node[1] + dir[1]}

			if next[1] < 0 || next[1] >= len(m) || next[0] < 0 || next[0] >= len(m[0]) ||
				visited[next] {
				continue
			}

			nextTile := m[next[1]][next[0]]
			if nextTile == "." {
				continue
			}

			queue = append(queue, next)
			visited[next] = true
		}
	}

	return visited
}

func printMap(m [][]string, borders map[[2]int]bool, inside map[[2]int]bool) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if borders[[2]int{x, y}] {
				fmt.Print(m[y][x])
			} else if inside[[2]int{x, y}] {
				fmt.Print("I")
			} else {
				fmt.Print("O")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	m := createMap(lines)
	start := findStart(m)
	visited := traverseMap(m, start)
	return len(visited) / 2
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	m := createMap(lines)
	start := findStart(m)
	borders := traverseMap(m, start)
	inside := make(map[[2]int]bool)
	for y, line := range m {
		for x := range line {
			if borders[[2]int{x, y}] {
				continue
			}

			crosses := 0
			x2, y2 := x, y

			for x2 < len(m[0]) && y2 < len(m) {
				c2 := m[y2][x2]
				if borders[[2]int{x2, y2}] && c2 != "L" && c2 != "7" {
					crosses++
				}
				x2++
				y2++
			}

			if crosses%2 == 1 {
				inside[[2]int{x, y}] = true
			}
		}
	}

	return len(inside) // 511 < x < 1000
}
