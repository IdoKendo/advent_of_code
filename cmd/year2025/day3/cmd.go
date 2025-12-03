package day3

import (
	"fmt"
	"os"
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

func largestJoltage(line string, batteryCount int) int {
	result := make([]byte, 0, batteryCount)

	lastIdx := 0
	for d := range batteryCount {
		end := len(line) - batteryCount + d + 1

		maxIdx := lastIdx
		maxBattery := line[lastIdx]

		for i := lastIdx; i < end; i++ {
			if line[i] > maxBattery {
				maxBattery = line[i]
				maxIdx = i
			}
		}

		result = append(result, maxBattery)
		lastIdx = maxIdx + 1
	}

	joltage := 0
	for _, c := range result {
		joltage = joltage*10 + int(c-'0')
	}
	return joltage
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	for _, line := range lines {
		result += largestJoltage(line, 2)
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	for _, line := range lines {
		result += largestJoltage(line, 12)
	}
	return result
}
