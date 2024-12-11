package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day11",
	Short: "day11",
	Long:  "day11",
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

func parseStones(input string) map[string]int {
	line := strings.TrimRight(input, "\n")
	stones := make(map[string]int)
	for _, v := range strings.Split(line, " ") {
		stones[v]++
	}

	return stones
}

func blink(stones map[string]int) map[string]int {
	newStones := make(map[string]int)

	for stone, count := range stones {
		var result []string

		if stone == "0" {
			result = []string{"1"}
		} else if len(stone)%2 == 0 {
			mid := len(stone) / 2
			left, _ := strconv.Atoi(stone[:mid])
			right, _ := strconv.Atoi(stone[mid:])
			result = []string{strconv.Itoa(left), strconv.Itoa(right)}
		} else {
			v, _ := strconv.Atoi(stone)
			v *= 2024
			result = []string{strconv.Itoa(v)}
		}

		for _, newStone := range result {
			newStones[newStone] += count
		}
	}

	return newStones
}

func Part1(input string) int {
	stones := parseStones(input)
	for range 25 {
		stones = blink(stones)
	}

	result := 0
	for _, count := range stones {
		result += count
	}

	return result
}

func Part2(input string) int {
	stones := parseStones(input)
	for range 75 {
		stones = blink(stones)
	}

	result := 0
	for _, count := range stones {
		result += count
	}

	return result
}
