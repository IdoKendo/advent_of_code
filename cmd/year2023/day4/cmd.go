package day4

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  "day4",
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

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	sum := 0
	for _, line := range lines {
		title := strings.Split(line, ":")
		numbers := strings.Split(title[1], "|")
		winningNumbers := numbers[0]
		hand := numbers[1]
		numbersList := strings.Split(winningNumbers, " ")
		handList := strings.Split(hand, " ")
		matches := -1
		for _, h := range handList {
			if h == "" {
				continue
			}
			if contains(numbersList, h) {
				matches += 1
			}
		}
		sum += int(math.Pow(2, float64(matches)))
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	pileSize := 0
	multipliers := make(map[int]int)
	for i := range lines {
		multipliers[i] = 1
	}
	for i, line := range lines {
		pileSize += multipliers[i]
		title := strings.Split(line, ":")
		numbers := strings.Split(title[1], "|")
		winningNumbers := numbers[0]
		hand := numbers[1]
		numbersList := strings.Split(winningNumbers, " ")
		handList := strings.Split(hand, " ")
		matches := -1
		for _, h := range handList {
			if h == "" {
				continue
			}
			if contains(numbersList, h) {
				matches += 1
			}
		}
		if matches >= 0 {
			for j := 1; j < matches+2; j++ {
				multipliers[i+j] += multipliers[i]
			}
		}
	}
	return pileSize
}
