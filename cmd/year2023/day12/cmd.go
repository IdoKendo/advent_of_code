package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day12",
	Short: "day12",
	Long:  "day12",
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

type state [3]int

func countArrangements(line string) int {
	parts := strings.Split(line, " ")
	desired := strings.Split(parts[1], ",")
	states := parts[0]
	arrangements := make([]int, len(desired))
	for i, d := range desired {
		arrangements[i], _ = strconv.Atoi(d)
	}

	pos := 0
	curr := map[state]int{{0, 0, 0}: 1}
	next := make(map[state]int)
	for _, ch := range states {
		for state, num := range curr {
			idx, count, expectingDot := state[0], state[1], state[2]
			switch {
			case (ch == '#' || ch == '?') && idx < len(arrangements) && expectingDot == 0:
				if ch == '?' && count == 0 {
					next[state] += num
				}
				count++
				if count == arrangements[idx] {
					idx, count, expectingDot = idx+1, 0, 1
				}
				next[[3]int{idx, count, expectingDot}] += num
			case (ch == '.' || ch == '?') && count == 0:
				expectingDot = 0
				next[[3]int{idx, count, expectingDot}] += num
			}
		}
		curr, next = next, curr
		for k := range next {
			delete(next, k)
		}
	}
	for s, v := range curr {
		if s[0] == len(arrangements) {
			pos += v
		}
	}
	return pos
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	total := 0

	for _, line := range lines {
		total += countArrangements(strings.TrimRight(line, "\n"))
	}

	return total
}

func Part2(input string) int {
	return 0
}
