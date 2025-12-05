package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  "day1",
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

func countIncreases(measurements []string) int {
	increases := 0
	for i := 1; i < len(measurements); i++ {
		prev, _ := strconv.Atoi(measurements[i-1])
		curr, _ := strconv.Atoi(measurements[i])
		if curr > prev {
			increases++
		}
	}

	return increases
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	return countIncreases(lines)
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	measurements := make([]string, len(lines)-3)
	idx := 0

	for i := range measurements {
		a, _ := strconv.Atoi(lines[i])
		b, _ := strconv.Atoi(lines[i+1])
		c, _ := strconv.Atoi(lines[i+2])
		measurements[idx] = strconv.Itoa(a + b + c)
		idx++
	}

	return countIncreases(measurements)
}
