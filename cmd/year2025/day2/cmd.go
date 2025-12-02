package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  "day2",
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

func isInvalidByHalf(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)

	if n%2 != 0 {
		return false
	}

	half := n / 2
	return s[:half] == s[half:]
}

func isInvalidByAny(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)
	half := n / 2

	for j := 1; j <= half; j++ {
		if n%j != 0 {
			continue
		}

		repeat := n / j
		pattern := s[:j]

		match := true
		for i := range repeat {
			if s[i*j:(i+1)*j] != pattern {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

func sumInvalidIDs(input string, isInvalid func(int) bool) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges := strings.Split(lines[0], ",")
	invalidsSum := 0
	for _, r := range ranges {
		IDs := strings.Split(r, "-")
		start, _ := strconv.Atoi(IDs[0])
		end, _ := strconv.Atoi(IDs[1])
		for i := start; i <= end; i++ {
			if isInvalid(i) {
				invalidsSum += i
			}
		}
	}
	return invalidsSum
}

func Part1(input string) int {
	return sumInvalidIDs(input, isInvalidByHalf)
}

func Part2(input string) int {
	return sumInvalidIDs(input, isInvalidByAny)
}
