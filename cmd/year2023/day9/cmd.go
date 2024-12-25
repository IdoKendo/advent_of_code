package day9

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "day9",
	Long:  "day9",
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

func genDiffs(sequence []int) []int {
	differences := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
	}
	return differences
}

func allZeroes(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}
	return true
}

func futureOf(seq []int) int {
	seqs := [][]int{seq}
	for {
		seq = genDiffs(seq)
		if allZeroes(seq) {
			break
		}
		seqs = append(seqs, seq)
	}

	for i := len(seqs) - 2; i >= 0; i-- {
		lastValue := seqs[i][len(seqs[i])-1]
		lastDifference := seqs[i+1][len(seqs[i+1])-1]
		seqs[i] = append(seqs[i], lastValue+lastDifference)
	}

	return seqs[0][len(seqs[0])-1]
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	result := 0
	for _, history := range lines {
		values := strings.Split(history, " ")
		seq := make([]int, len(values))
		for i, value := range values {
			seq[i], _ = strconv.Atoi(value)
		}
		result += futureOf(seq)
	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	result := 0
	for _, history := range lines {
		values := strings.Split(history, " ")
		seq := make([]int, len(values))
		for i, value := range values {
			seq[i], _ = strconv.Atoi(value)
		}
		slices.Reverse(seq)
		result += futureOf(seq)
	}

	return result
}
