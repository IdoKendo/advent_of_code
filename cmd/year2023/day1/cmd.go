package day1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  "day1",
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

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	sum := 0
	re := regexp.MustCompile("[0-9]")
	for _, line := range lines {
		numbers := re.FindAllString(line, -1)
		i, _ := strconv.Atoi(numbers[0])
		j, _ := strconv.Atoi(numbers[len(numbers)-1])
		sum += i*10 + j
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	sum := 0
	re := regexp.MustCompile("[0-9]")
	words := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for _, line := range lines {
		for k, v := range words {
			line = strings.ReplaceAll(line, k, v)
		}
		numbers := re.FindAllString(line, -1)
		i, _ := strconv.Atoi(numbers[0])
		j, _ := strconv.Atoi(numbers[len(numbers)-1])
		sum += i*10 + j
	}
	return sum
}
