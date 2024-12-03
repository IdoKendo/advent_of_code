package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	// "strconv"

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

func parseInstruction(re *regexp.Regexp, match string) int {
	nums := re.FindStringSubmatch(match)
	n2, _ := strconv.Atoi(nums[1])
	n1, _ := strconv.Atoi(nums[2])
	return n1 * n2
}

func Part1(input string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	instructions := re.FindAllString(input, -1)
	sum := 0
	re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, instruction := range instructions {
		sum += parseInstruction(re, instruction)
	}

	return sum
}

func Part2(input string) int {
	re := regexp.MustCompile(`(don't\(\)|do\(\)|mul\(\d+,\d+\))`)
	instructions := re.FindAllString(input, -1)
	sum := 0
	re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	enabled := true
	for _, instruction := range instructions {
		if instruction == "don't()" {
			enabled = false
			continue
		}
		if instruction == "do()" {
			enabled = true
			continue
		}
		if !enabled {
			continue
		}
		sum += parseInstruction(re, instruction)
	}

	return sum
}
