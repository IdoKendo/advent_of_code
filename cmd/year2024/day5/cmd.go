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

func generateRules(lines []string) map[int]map[int]bool {
	rules := make(map[int]map[int]bool)
	for _, line := range lines {
		n := strings.Split(line, "|")
		x, _ := strconv.Atoi(n[0])
		y, _ := strconv.Atoi(n[1])
		if rules[x] == nil {
			rules[x] = make(map[int]bool)
		}
		rules[x][y] = true
	}

	return rules
}

func generatePages(update string) []int {
	nums := strings.Split(update, ",")
	pages := make([]int, len(nums))
	for i, num := range nums {
		pages[i], _ = strconv.Atoi(num)
	}

	return pages
}

func isValid(pages []int, rules map[int]map[int]bool) bool {
	for i, page := range pages {
		if len(rules[page]) < len(pages[i+1:]) {
			return false
		}
		for _, c := range pages[i+1:] {
			if !rules[page][c] {
				return false
			}
		}
	}

	return true
}

func Part1(input string) int {
	lines := strings.Split(input, "\n\n")
	rules := generateRules(strings.Split(lines[0], "\n"))
	updates := strings.Split(lines[1], "\n")
	updates = updates[:len(updates)-1]
	sum := 0
	for _, update := range updates {
		pages := generatePages(update)
		if isValid(pages, rules) {
			sum += pages[len(pages)/2]
		}
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	rules := generateRules(strings.Split(lines[0], "\n"))
	updates := strings.Split(lines[1], "\n")
	updates = updates[:len(updates)-1]
	sum := 0
	for _, update := range updates {
		pages := generatePages(update)
		if !isValid(pages, rules) {
			sort.Slice(pages, func(i, j int) bool {
				a, b := pages[i], pages[j]
				if rules[a] != nil && rules[a][b] {
					return true
				}
				if rules[b] != nil && rules[b][a] {
					return false
				}
				return a < b
			})
			sum += pages[len(pages)/2]
		}
	}
	return sum
}
