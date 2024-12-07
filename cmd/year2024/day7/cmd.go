package day7

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "day7",
	Long:  "day7",
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

func evaluates(nums []int, testValue int) bool {
	sums := []int{nums[0]}

	for _, num := range nums[1:] {
		nextSums := []int{}
		for _, n := range sums {
			nextSums = append(nextSums, num*n)
			nextSums = append(nextSums, num+n)
		}
		sums = nextSums
	}

	for _, sum := range sums {
		if sum == testValue {
			return true
		}
	}

	return false
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	equations := lines[:len(lines)-1]
	re := regexp.MustCompile(`\b\d+\b`)
	sum := 0
	for _, equation := range equations {
		parts := re.FindAllString(equation, -1)
		nums := make([]int, len(parts))
		for i, s := range parts {
			num, _ := strconv.Atoi(s)
			nums[i] = num
		}
		if evaluates(nums[1:], nums[0]) {
			sum += nums[0]
		}
	}

	return sum
}

func evaluatesWithConcat(nums []int, testValue int) bool {
	sums := []int{nums[0]}

	for _, num := range nums[1:] {
		nextSums := []int{}
		for _, n := range sums {
			nextSums = append(nextSums, num*n)
			nextSums = append(nextSums, num+n)
			a := strconv.Itoa(n)
			b := strconv.Itoa(num)
			c, _ := strconv.Atoi(fmt.Sprintf("%s%s", a, b))
			nextSums = append(nextSums, c)
		}
		sums = nextSums
	}

	for _, sum := range sums {
		if sum == testValue {
			return true
		}
	}

	return false
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	equations := lines[:len(lines)-1]
	re := regexp.MustCompile(`\b\d+\b`)
	sum := 0
	for _, equation := range equations {
		parts := re.FindAllString(equation, -1)
		nums := make([]int, len(parts))
		for i, s := range parts {
			num, _ := strconv.Atoi(s)
			nums[i] = num
		}
		if evaluatesWithConcat(nums[1:], nums[0]) {
			sum += nums[0]
		}
	}

	return sum
}
