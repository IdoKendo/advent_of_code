package day7

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "day7",
	Long:  "day7",
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

type TachyonManifold struct {
	start   int
	diagram [][]bool
}

func newTachyonManifold(input string) TachyonManifold {
	t := TachyonManifold{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	t.diagram = make([][]bool, len(lines))
	for y, line := range lines {
		t.diagram[y] = make([]bool, len(line))
		for x, ch := range line {
			switch ch {
			case 'S':
				t.start = x
			case '^':
				t.diagram[y][x] = true
			}
		}
	}

	return t
}

func (t *TachyonManifold) ShootBeam() (int, int) {
	splits := make([]int, len(t.diagram[0]))
	splits[t.start] = 1
	count := 0
	for _, line := range t.diagram {
		for idx, r := range line {
			if r && splits[idx] != 0 {
				beams := splits[idx]

				splits[idx+1] += beams
				splits[idx-1] += beams

				splits[idx] = 0

				count++
			}
		}
	}

	total := 0
	for _, n := range splits {
		total += n
	}

	return count, total
}

func Part1(input string) int {
	tachyonManifold := newTachyonManifold(input)
	count, _ := tachyonManifold.ShootBeam()
	return count
}

func Part2(input string) int {
	tachyonManifold := newTachyonManifold(input)
	_, total := tachyonManifold.ShootBeam()
	return total
}
