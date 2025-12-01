package day1

import (
	"fmt"
	"math"
	"os"
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

const (
	dialSize int = 100
	startPos int = 50
)

func floorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func countZeroes(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	zeroPointers := 0
	zeroPasses := 0
	pos := startPos
	for _, line := range lines {
		prev := pos
		val, _ := strconv.Atoi(line[1:])

		if line[0] == 'R' {
			pos = prev + val
			zeroPasses += floorDiv(pos, dialSize) - floorDiv(prev, dialSize)
		} else {
			pos = prev - val
			zeroPasses += floorDiv(prev-1, dialSize) - floorDiv(pos-1, dialSize)
		}

		if pos%dialSize == 0 {
			zeroPointers++
		}
	}
	return zeroPointers, zeroPasses
}

func Part1(input string) int {
	pointers, _ := countZeroes(input)
	return pointers
}

func Part2(input string) int {
	_, passes := countZeroes(input)
	return passes
}
