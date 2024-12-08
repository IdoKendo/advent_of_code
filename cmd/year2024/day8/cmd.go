package day8

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day8",
	Short: "day8",
	Long:  "day8",
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

func getAntennas(rows []string) [][][2]int {
	antennaByFrequency := make(map[rune][][2]int)
	for x, row := range rows {
		for y, ch := range row {
			if ch == '.' {
				continue
			}
			antennaByFrequency[ch] = append(antennaByFrequency[ch], [2]int{x, y})
		}
	}

	antennas := make([][][2]int, 0, len(antennaByFrequency))
	for _, positions := range antennaByFrequency {
		antennas = append(antennas, positions)
	}
	return antennas
}

func findAntinodes(antenna1, antenna2 [2]int, width, height int) [][2]int {
	x1, y1 := antenna1[0], antenna1[1]
	x2, y2 := antenna2[0], antenna2[1]
	antinodes := [][2]int{
		{2*x2 - x1, 2*y2 - y1},
		{2*x1 - x2, 2*y1 - y2},
	}

	validAntinodes := make([][2]int, 0, len(antinodes))
	for _, antinode := range antinodes {
		if antinode[0] >= 0 && antinode[0] < width && antinode[1] >= 0 && antinode[1] < height {
			validAntinodes = append(validAntinodes, antinode)
		}
	}
	return validAntinodes
}

func printGrid(grid []string, antinodes map[[2]int]bool) {
	fmt.Println()
	for x, row := range grid {
		for y, val := range row {
			if antinodes[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(string(val))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	antennas := getAntennas(lines)
	antinodes := make(map[[2]int]bool)
	for _, a := range antennas {
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				for _, antinode := range findAntinodes(a[i], a[j], len(lines[0]), len(lines)) {
					antinodes[antinode] = true
				}
			}
		}
	}

	return len(antinodes)
}

func findReoccurringAntinodes(antenna1, antenna2 [2]int, width, height int) [][2]int {
	x1, y1 := antenna1[0], antenna1[1]
	x2, y2 := antenna2[0], antenna2[1]
	dx := x2 - x1
	dy := y2 - y1
	antinodes := make([][2]int, 0)

	for multiplier := 0; multiplier < max(width, height); multiplier++ {
		antinode := [2]int{x1 + multiplier*dx, y1 + multiplier*dy}
		if antinode[0] < 0 || antinode[0] >= width || antinode[1] < 0 || antinode[1] >= height {
			break
		}
		antinodes = append(antinodes, antinode)
	}

	for multiplier := -1; multiplier > -max(width, height); multiplier-- {
		antinode := [2]int{x1 + multiplier*dx, y1 + multiplier*dy}
		if antinode[0] < 0 || antinode[0] >= width || antinode[1] < 0 || antinode[1] >= height {
			break
		}
		antinodes = append(antinodes, antinode)
	}

	return antinodes
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	antennas := getAntennas(lines)
	antinodes := make(map[[2]int]bool)
	for _, group := range antennas {
		for i := 0; i < len(group); i++ {
			for j := i + 1; j < len(group); j++ {
				for _, antinode := range findReoccurringAntinodes(group[i], group[j], len(lines[0]), len(lines)) {
					antinodes[antinode] = true
				}
			}
		}
	}

	return len(antinodes)
}
