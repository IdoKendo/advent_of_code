package day19

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day19",
	Short: "day19",
	Long:  "day19",
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

func isPossible(design string, towels []string) bool {
	memo := make(map[string]bool)
	return isPossibleHelper(design, towels, memo)
}

func isPossibleHelper(design string, towels []string, memo map[string]bool) bool {
	if val, found := memo[design]; found {
		return val
	}

	if len(design) == 0 {
		return true
	}

	for _, pattern := range towels {
		trimmed := strings.TrimSuffix(design, pattern)
		if trimmed == design {
			continue
		}
		if isPossibleHelper(trimmed, towels, memo) {
			memo[design] = true
			return true
		}
	}

	memo[design] = false
	return false
}

func Part1(input string) int {
	parts := strings.Split(input, "\n\n")
	designs := strings.Split(parts[1], "\n")
	possibleDesigns := 0
	towels := strings.Split(parts[0], ", ")

	for _, design := range designs[:len(designs)-1] {
		if isPossible(design, towels) {
			possibleDesigns++
		}
	}

	return possibleDesigns
}

func countPossibleDesigns(design string, towels []string, cache map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	if possibilities, found := cache[design]; found {
		return possibilities
	}

	possibilities := 0
	for _, towel := range towels {
		trimmed := strings.TrimSuffix(design, towel)
		if trimmed != design {
			possibilities += countPossibleDesigns(trimmed, towels, cache)
		}
	}

	cache[design] = possibilities

	return possibilities
}

func Part2(input string) int {
	parts := strings.Split(input, "\n\n")
	designs := strings.Split(parts[1], "\n")
	possibleDesigns := 0
	towels := strings.Split(parts[0], ", ")
	cache := make(map[string]int)
	for _, design := range designs[:len(designs)-1] {
		possibleDesigns += countPossibleDesigns(design, towels, cache)
	}

	return possibleDesigns
}
