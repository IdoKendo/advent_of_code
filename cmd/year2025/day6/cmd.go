package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day6",
	Short: "day6",
	Long:  "day6",
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

type Problem struct {
	numbers []int
	op      string
}

func (p *Problem) Result() int {
	var sum int
	switch p.op {
	case "*":
		sum = 1
	case "+":
		sum = 0
	}

	for _, num := range p.numbers {
		switch p.op {
		case "*":
			sum *= num
		case "+":
			sum += num
		}
	}

	return sum
}

func transpose(matrix [][]string) [][]string {
	if len(matrix) == 0 {
		return nil
	}
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]string, cols)
	for i := range transposed {
		transposed[i] = make([]string, rows)
	}

	for r := range rows {
		if len(matrix[r]) == 0 {
			break
		}
		for c := range cols {
			transposed[c][r] = matrix[r][c]
		}
	}

	return transposed
}

func transformToCephalopodProblems(matrix [][]string) map[int]*Problem {
	idx := 0
	problems := make(map[int]*Problem)
	problems[idx] = &Problem{}
	for _, line := range matrix {
		text := strings.TrimSpace(strings.Join(line, ""))
		if text == "" {
			idx++
			problems[idx] = &Problem{}
			continue
		}
		pos := strings.IndexAny(text, "*+")
		if pos != -1 {
			problems[idx].op = string(text[pos])
		}
		text = strings.TrimSpace(strings.Map(func(r rune) rune {
			switch r {
			case '*', '+':
				return -1
			default:
				return r
			}
		}, text))
		number, _ := strconv.Atoi(text)
		problems[idx].numbers = append(problems[idx].numbers, number)
	}

	return problems
}

func transformToProblems(input string) map[int]*Problem {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	problems := make(map[int]*Problem)

	for _, line := range lines {
		for i, field := range strings.Fields(line) {
			if problems[i] == nil {
				problems[i] = &Problem{}
			}

			number, err := strconv.Atoi(field)
			if err != nil {
				problems[i].op = field
				continue
			}

			problems[i].numbers = append(problems[i].numbers, number)
		}
	}

	return problems
}

func solve(problems map[int]*Problem) int {
	sum := 0
	for _, problem := range problems {
		sum += problem.Result()
	}

	return sum
}

func Part1(input string) int {
	problems := transformToProblems(input)
	return solve(problems)
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	original := make([][]string, len(lines))
	for i, line := range lines {
		original[i] = strings.Split(line, "")
	}
	transposed := transpose(original)
	problems := transformToCephalopodProblems(transposed)

	return solve(problems)
}
