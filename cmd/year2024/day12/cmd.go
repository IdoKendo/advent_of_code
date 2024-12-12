package day12

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day12",
	Short: "day12",
	Long:  "day12",
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

func dfs(farm [][]rune, visited [][]bool, y, x int, sides *[][3]int) (int, int) {
	visited[y][x] = true
	neighbors := getNeighbors(y, x)
	perimeter := 0
	area := 1
	for _, n := range neighbors {
		i, j := n[0], n[1]
		if i < 0 || i >= len(farm) || j < 0 || j >= len(farm[i]) || farm[i][j] != farm[y][x] {
			if sides != nil {
				*sides = append(*sides, [3]int{i, j, n[2]})
			}
			perimeter++
		} else if !visited[i][j] {
			a, p := dfs(farm, visited, i, j, sides)
			area += a
			perimeter += p
		}
	}
	return area, perimeter
}

func getNeighbors(y, x int) [][3]int {
	return [][3]int{
		{y - 1, x, 0},
		{y + 1, x, 1},
		{y, x - 1, 2},
		{y, x + 1, 3},
	}
}

func getSideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	sideCount := 0
	for _, side := range sides {
		combinations := getNeighbors(side[0], side[1])
		combinationFound := false

		for _, combination := range combinations {
			combination[2] = side[2]
			if _, found := sideMap[combination]; found {
				combinationFound = true
			}
		}
		if !combinationFound {
			sideCount++
		}

		sideMap[side] = true
	}

	return sideCount
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	farm := make([][]rune, len(lines))
	for i, line := range lines {
		farm[i] = []rune(line)
	}
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for y := 0; y < len(farm); y++ {
		for x := 0; x < len(farm[y]); x++ {
			if !visited[y][x] {
				area, perimeter := dfs(farm, visited, y, x, nil)
				sum += area * perimeter
			}
		}
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	farm := make([][]rune, len(lines))
	for i, line := range lines {
		farm[i] = []rune(line)
	}
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for y := 0; y < len(farm); y++ {
		for x := 0; x < len(farm[y]); x++ {
			if !visited[y][x] {
				sides := make([][3]int, 0)
				area, _ := dfs(farm, visited, y, x, &sides)
				sideCount := getSideCount(sides)
				sum += area * sideCount
			}
		}
	}
	return sum
}
