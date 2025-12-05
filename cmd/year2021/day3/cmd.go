package day3

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  "day3",
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

func countBits(lines []string) ([]int, []int) {
	ones := make([]int, len(lines[0]))
	zeroes := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, bit := range strings.Split(line, "") {
			if bit == "1" {
				ones[i]++
			} else {
				zeroes[i]++
			}
		}
	}

	return ones, zeroes
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ones, zeroes := countBits(lines)
	gamma := make([]int, len(ones))
	epsilon := make([]int, len(ones))
	for i := range len(ones) {
		if ones[i] > zeroes[i] {
			gamma[i] = 1
		} else {
			epsilon[i] = 1
		}
	}

	gammaValue := 0
	epsilonValue := 0
	mul := 1

	for i := len(ones) - 1; i >= 0; i-- {
		gammaValue += gamma[i] * mul
		epsilonValue += epsilon[i] * mul
		mul *= 2
	}

	return gammaValue * epsilonValue
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ones, zeroes := countBits(lines)
	oxygenLines := slices.Clone(lines)
	CO2Lines := slices.Clone(lines)
	for i := range len(ones) {
		if ones[i] > zeroes[i] {
			for j, line := range CO2Lines {
				if line == "" {
					continue
				} else if line[i] != '0' {
					CO2Lines[j] = ""
				} else {
					oxygenLines[j] = ""
				}
			}
		} else {
			for j, line := range oxygenLines {
				if line == "" {
					continue
				} else if line[i] != '1' {
					oxygenLines[j] = ""
				} else {
					CO2Lines[j] = ""
				}
			}
		}
	}

	return 0
}
