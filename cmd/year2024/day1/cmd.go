package day1

import (
	"fmt"
	"math"
	"os"
	"slices"
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

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	left := make([]int, len(input))
	right := make([]int, len(input))
	result := 0
	for i, line := range lines {
		locations := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(locations[0])
		right[i], _ = strconv.Atoi(locations[1])
	}
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		result += int(math.Abs(float64(left[i] - right[i])))
	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	appearances := make(map[int]int)
	similarities := make(map[int]int)
	result := 0
	for _, line := range lines {
		locations := strings.Split(line, "   ")
		l, _ := strconv.Atoi(locations[0])
		r, _ := strconv.Atoi(locations[1])
		appearances[l]++
		similarities[r]++
	}
	for key, val := range appearances {
		result += key * val * similarities[key]
	}

	return result
}
