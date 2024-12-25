package day3

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  "day3",
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

func parse(grid [][]rune, x int, y int) int {
	for x > 0 && unicode.IsDigit(grid[y][x-1]) {
		x--
	}
	n := 0
	for x < len(grid[y]) && unicode.IsDigit(grid[y][x]) {
		n = (n * 10) + int(grid[y][x]-'0')
		x++
	}
	return n
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	grid := make([][]rune, 0, 140)
	dirs := []int{-1, 0, 1}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	sum := 0
	for y := 0; y < len(grid); y++ {
		var n int
		var isValid bool

		for x := 0; x < len(grid[y]); x++ {
			if !unicode.IsDigit(grid[y][x]) {
				if isValid {
					sum += n
				}
				n = 0
				isValid = false
				continue
			}

			n = (n * 10) + int(grid[y][x]-'0')

			for _, dy := range dirs {
				for _, dx := range dirs {
					y2 := y + dy
					x2 := x + dx
					if x2 >= 0 && x2 < len(grid[y]) && y2 >= 0 && y2 < len(grid) &&
						(x2 != x || y2 != y) {
						if grid[y2][x2] != '.' && !unicode.IsDigit(grid[y2][x2]) {
							isValid = true
						}
					}
				}
			}
		}

		if isValid && n > 0 {
			sum += n
		}
	}

	return sum
}

func Part2(input string) int {
	grid := make([][]rune, 0, 140)
	dirs := []int{-1, 0, 1}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}
	sum := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '*' {
				continue
			}

			ratios := make([]int, 0)
			for _, dy := range dirs {
				for _, dx := range dirs {
					x2 := x + dx
					y2 := y + dy
					if x2 < 0 || x2 >= len(grid[y]) || y2 < 0 || y2 >= len(grid) ||
						!unicode.IsDigit(grid[y2][x2]) {
						continue
					}
					v := parse(grid, x2, y2)
					if !contains(ratios, v) {
						ratios = append(ratios, v)
					}
				}
			}

			if len(ratios) == 2 {
				sum += ratios[0] * ratios[1]
			}
		}
	}

	return sum
}
