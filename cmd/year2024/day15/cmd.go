package day15

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day15",
	Short: "day15",
	Long:  "day15",
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
	x, y int
}

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

type house struct {
	grid  map[point]string
	robot point
}

func newHouse(input string) house {
	grid, robot := map[point]string{}, point{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			if r == '@' {
				robot = point{x, y}
				r = '.'
			}
			grid[point{x, y}] = string(r)
		}
	}

	return house{grid, robot}
}

func (h *house) moveRobot(move string) {
	delta := map[string]point{
		"^": {0, -1}, ">": {1, 0}, "v": {0, 1}, "<": {-1, 0},
		"[": {1, 0}, "]": {-1, 0},
	}

	queue := []point{h.robot}
	boxes := map[point]string{}
	blocked := false
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if _, ok := boxes[p]; ok {
			continue
		}
		boxes[p] = h.grid[p]

		switch n := p.add(delta[move]); h.grid[n] {
		case "#":
			blocked = true
		case "[", "]":
			queue = append(queue, n.add(delta[h.grid[n]]))
			fallthrough
		case "O":
			queue = append(queue, n)
		}
	}
	if blocked {
		return
	}

	for b := range boxes {
		h.grid[b] = "."
	}
	for b := range boxes {
		h.grid[b.add(delta[move])] = boxes[b]
	}
	h.robot = h.robot.add(delta[move])

}

func (h house) gps() int {
	gps := 0
	for p, r := range h.grid {
		if r == "O" || r == "[" {
			gps += 100*p.y + p.x
		}
	}
	return gps
}

func Part1(input string) int {
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	h := newHouse(parts[0])

	for _, move := range strings.ReplaceAll(parts[1], "\n", "") {
		h.moveRobot(string(move))
	}
	return h.gps()
}

func Part2(input string) int {
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")
	h := newHouse(r.Replace(parts[0]))
	for _, move := range strings.ReplaceAll(parts[1], "\n", "") {
		h.moveRobot(string(move))
	}
	return h.gps()
}
