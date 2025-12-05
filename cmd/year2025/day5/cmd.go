package day5

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  "day5",
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

func parseRangesAndIngredientsIDs(lines []string) ([]int, []int) {
	var ranges []int
	var ingredientsIDs []int

	step := 0
	for _, line := range lines {
		if line == "" {
			step++
			continue
		}

		switch step {
		case 0:
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			end++

			i := sort.SearchInts(ranges, start)
			j := sort.SearchInts(ranges, end)

			var add []int
			if i%2 == 0 {
				add = append(add, start)
			}
			if j%2 == 0 {
				add = append(add, end)
			}

			newRanges := make([]int, 0, len(ranges)-(j-i)+len(add))
			newRanges = append(newRanges, ranges[:i]...)
			newRanges = append(newRanges, add...)
			newRanges = append(newRanges, ranges[j:]...)
			ranges = newRanges

		case 1:
			id, _ := strconv.Atoi(line)
			ingredientsIDs = append(ingredientsIDs, id)
		}
	}

	return ranges, ingredientsIDs
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges, ingredientsIDs := parseRangesAndIngredientsIDs(lines)

	count := 0
	for _, ID := range ingredientsIDs {
		if sort.SearchInts(ranges, ID)%2 == 1 {
			count++
		}
	}

	return count
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges, _ := parseRangesAndIngredientsIDs(lines)

	count := 0
	for i := 0; i < len(ranges)/2; i++ {
		count += ranges[2*i+1] - ranges[2*i]
	}

	return count
}
