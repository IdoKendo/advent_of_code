package day13

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day13",
	Short: "day13",
	Long:  "day13",
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

func minTokens(lines []string, prizeDiff int) int {
	a := point{}
	b := point{}
	prize := point{}

	buttonARegex := regexp.MustCompile(`Button A: X([+=-]?\d+), Y([+=-]?\d+)`)
	buttonBRegex := regexp.MustCompile(`Button B: X([+=-]?\d+), Y([+=-]?\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=?(\d+), Y=?(\d+)`)

	if matches := buttonARegex.FindStringSubmatch(strings.TrimSpace(lines[0])); matches != nil {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		a = point{x: x, y: y}
	}

	if matches := buttonBRegex.FindStringSubmatch(strings.TrimSpace(lines[1])); matches != nil {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		b = point{x: x, y: y}
	}

	if matches := prizeRegex.FindStringSubmatch(strings.TrimSpace(lines[2])); matches != nil {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		prize = point{x: x + prizeDiff, y: y + prizeDiff}
	}

	numeratorA := prize.x*b.y - prize.y*b.x
	numeratorB := prize.y*a.x - prize.x*a.y
	denominator := a.x*b.y - a.y*b.x

	if denominator != 0 {
		a := float64(numeratorA) / float64(denominator)
		b := float64(numeratorB) / float64(denominator)
		if a == float64(int(a)) && b == float64(int(b)) {
			return 3*int(a) + int(b)
		}
	}

	return 0

}

func Part1(input string) int {
	lines := strings.Split(input, "\n\n")
	tokens := 0
	for _, l := range lines {
		tokens += minTokens(strings.Split(l, "\n"), 0)
	}
	return tokens
}

func Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	tokens := 0
	for _, l := range lines {
		tokens += minTokens(strings.Split(l, "\n"), 10000000000000)
	}
	return tokens
}
