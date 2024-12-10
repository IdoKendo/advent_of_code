package day10

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
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

type point struct {
	x, y, v int
}

type topographicMap struct {
	startingPoints []point
	topographicMap [][]int
}

func (t topographicMap) print(w io.Writer) {
	for _, l := range t.topographicMap {
		fmt.Fprintln(w, l)
	}
}

func (t topographicMap) scores() int {
	total := 0
	for _, sp := range t.startingPoints {
		total += calculateScore(t.topographicMap, sp)
	}
	return total
}

func (t topographicMap) ratings() int {
	total := 0
	for _, ep := range t.startingPoints {
		total += calculateRating(t.topographicMap, ep)
	}
	return total
}

var directions = [][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func calculateScore(m [][]int, s point) int {
	queue := []point{s}
	visited := make(map[point]bool)
	visited[s] = true
	score := 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.v == 9 && !visited[p] {
			score++
		}
		visited[p] = true
		for _, dir := range directions {
			y := p.y + dir[0]
			x := p.x + dir[1]
			if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
				continue
			}

			c := m[y][x]
			if c-p.v == 1 {
				queue = append(queue, point{x: x, y: y, v: c})
			}
		}
	}
	return score
}

func calculateRating(m [][]int, s point) int {
	queue := []point{s}
	score := 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.v == 9 {
			score++
		}
		for _, dir := range directions {
			y := p.y + dir[0]
			x := p.x + dir[1]
			if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
				continue
			}

			c := m[y][x]
			if c-p.v == 1 {
				queue = append(queue, point{x: x, y: y, v: c})
			}
		}
	}
	return score
}

func newtopographicMap(input string) topographicMap {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	tm := make([][]int, len(lines))
	var sp []point
	for i, line := range lines {
		digits := strings.Split(line, "")
		tm[i] = make([]int, len(digits))
		for j, digit := range digits {
			tm[i][j], _ = strconv.Atoi(digit)
			if tm[i][j] == 0 {
				sp = append(sp, point{x: j, y: i, v: 0})
			}
		}
	}
	return topographicMap{
		startingPoints: sp,
		topographicMap: tm,
	}
}

func Part1(input string) int {
	tm := newtopographicMap(input)
	var buf bytes.Buffer
	tm.print(&buf)
	return tm.scores()
}

func Part2(input string) int {
	tm := newtopographicMap(input)
	var buf bytes.Buffer
	tm.print(&buf)
	return tm.ratings()
}
