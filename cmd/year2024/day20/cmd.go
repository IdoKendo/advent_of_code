package day20

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day20",
	Short: "day20",
	Long:  "day20",
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

func isPossible(combination string, towels map[string]bool) bool {
	if towels[combination] || combination == "" {
		return true
	}
	if len(combination) == 1 && !towels[combination] {
		return false
	}

	possible := false
	for i := range len(combination) {
		if towels[combination[:i]] {
			possible = possible || isPossible(combination[i+1:], towels)
			if possible {
				return true
			}
		}
	}

	return false
}

func Part1(input string) int {
	parts := strings.Split(input, "\n\n")
	towelList := strings.Split(parts[0], ", ")
	combinations := strings.Split(parts[1], "\n")
	combinations = combinations[:len(combinations)-1]
	possibleCombos := 0
	towels := make(map[string]bool)

	fmt.Println("Towels we have:")
	for _, t := range towelList {
		fmt.Println(t)
		towels[t] = true
	}

	fmt.Println()
	for _, combination := range combinations {
		if isPossible(combination, towels) {
			fmt.Println("Possible: ", combination)
			possibleCombos++
		}
	}

	return possibleCombos // < 249
}

func Part2(input string) int {
	return 0
}
