package day4

import (
	"fmt"
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

func createMatrix(input string) [][]rune {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	rowCount := len(lines)
	colCount := len(lines[0])
	result := make([][]rune, rowCount)
	for i := range result {
		result[i] = make([]rune, colCount)
	}
	for i, line := range lines {
		for j, char := range line {
			result[i][j] = char
		}
	}
	return result
}

var directions = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var order = map[rune]rune{
	'X': 'M',
	'M': 'A',
	'A': 'S',
}

func countXmas(matrix [][]rune, i, j int, letter rune, direction [2]int) int {
	if letter == 'S' {
		return 1
	}

	x, y := i+direction[0], j+direction[1]
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] != order[letter] {
		return 0
	}

	return countXmas(matrix, x, y, order[letter], direction)
}

func Part1(input string) int {
	matrix := createMatrix(input)
	xmas := 0
	for i, line := range matrix {
		for j, letter := range line {
			if letter == 'X' {
				for _, dir := range directions {
					xmas += countXmas(matrix, i, j, letter, dir)
				}
			}
		}
	}

	return xmas
}

var crosses = [2][2]int{
	{-1, -1},
	{-1, 1},
}

func foundCrossmas(matrix [][]rune, i, j int) bool {
	found := 0
	for _, dir := range crosses {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] == 'A' || matrix[x][y] == 'X' {
			return false
		}
		if matrix[x][y] == 'M' {
			m, n := i-dir[0], j-dir[1]
			if m < 0 || m >= len(matrix) || n < 0 || n >= len(matrix[0]) || matrix[m][n] != 'S' {
				return false
			}
			found++
		}
		if matrix[x][y] == 'S' {
			m, n := i-dir[0], j-dir[1]
			if m < 0 || m >= len(matrix) || n < 0 || n >= len(matrix[0]) || matrix[m][n] != 'M' {
				return false
			}
			found++
		}
	}

	return found == 2
}

func Part2(input string) int {
	matrix := createMatrix(input)
	xmas := 0
	for i, line := range matrix {
		for j, letter := range line {
			if letter == 'A' {
				if foundCrossmas(matrix, i, j) {
					xmas++
				}
			}
		}
	}
	return xmas
}
