package day21

import (
	"fmt"
	"image"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day21",
	Short: "day21",
	Long:  "day21",
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

var numpad = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{" ", "0", "A"},
}

var dirpad = [][]string{
	{" ", "^", "A"},
	{"<", "v", ">"},
}

var numpad_positions = buildPositions(numpad)

var dirpad_positions = buildPositions(dirpad)

func buildPositions(keypad [][]string) map[string]image.Point {
	positions := make(map[string]image.Point)
	for y, row := range keypad {
		for x, char := range row {
			positions[char] = image.Point{X: x, Y: y}
		}
	}
	return positions
}

type keypad int

const (
	numpad_type keypad = iota
	dirpad_type
)

func (kp keypad) positions() map[string]image.Point {
	if kp == numpad_type {
		return numpad_positions
	}
	return dirpad_positions
}

func (kp keypad) sequences(code string, start string) [][]string {
	if code == "" {
		return [][]string{{}}
	}

	keypadPositions := kp.positions()
	position := keypadPositions[start]
	nextPosition := keypadPositions[string(code[0])]
	gap := keypadPositions[" "]

	sequences := [][]string{}
	for _, sequence := range kp.sequences(code[1:], string(code[0])) {
		for _, move := range moves(position, nextPosition, gap) {
			sequences = append(sequences, append([]string{move}, sequence...))
		}
	}
	return sequences
}

var memo = make(map[string]int)

func (kp keypad) shortestSequence(code string, proxies int) int {
	key := fmt.Sprintf("%s:%d:%d", code, proxies, kp)
	if result, exists := memo[key]; exists {
		return result
	}

	sequences := kp.sequences(code, "A")
	if proxies == 0 {
		minLength := math.MaxInt
		for _, sequence := range sequences {
			totalLength := 0
			for _, part := range sequence {
				totalLength += len(part)
			}
			if totalLength < minLength {
				minLength = totalLength
			}
		}
		memo[key] = minLength
		return minLength
	}

	minLength := math.MaxInt
	for _, sequence := range sequences {
		totalLength := 0
		for _, dircode := range sequence {
			totalLength += dirpad_type.shortestSequence(dircode, proxies-1)
		}
		if totalLength < minLength {
			minLength = totalLength
		}
	}
	memo[key] = minLength

	return memo[key]
}

func moves(position, nextPosition, gap image.Point) []string {
	hArrow := "<"
	if nextPosition.X > position.X {
		hArrow = ">"
	}

	vArrow := "^"
	if nextPosition.Y > position.Y {
		vArrow = "v"
	}

	hDist := int(math.Abs(float64(position.X - nextPosition.X)))
	vDist := int(math.Abs(float64(position.Y - nextPosition.Y)))

	var options []string
	if position == nextPosition {
		options = append(options, "A")
	} else if position.X == nextPosition.X {
		options = append(options, strings.Repeat(vArrow, vDist)+"A")
	} else if position.Y == nextPosition.Y {
		options = append(options, strings.Repeat(hArrow, hDist)+"A")
	} else {
		if !gapInPath(gap.X, gap.Y, nextPosition.X, position.Y, position.X, nextPosition.Y) {
			options = append(options, strings.Repeat(hArrow, hDist)+strings.Repeat(vArrow, vDist)+"A")
		}
		if !gapInPath(gap.X, gap.Y, position.X, nextPosition.Y, nextPosition.X, position.Y) {
			options = append(options, strings.Repeat(vArrow, vDist)+strings.Repeat(hArrow, hDist)+"A")
		}
	}
	return options
}

func gapInPath(gapX, gapY, startX, startY, endX, endY int) bool {
	if gapX == startX {
		step := 1
		if startY > endY {
			step = -1
		}
		for y := startY; y != endY+step; y += step {
			if y == gapY {
				return true
			}
		}
	} else if gapY == startY {
		step := 1
		if startX > endX {
			step = -1
		}
		for x := startX; x != endX+step; x += step {
			if x == gapX {
				return true
			}
		}
	}

	return false
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	re := regexp.MustCompile(`\d+`)
	fmt.Println()
	for _, line := range lines {
		numeric, _ := strconv.Atoi(re.FindString(line))
		result += numeric * numpad_type.shortestSequence(line, 2)
	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	re := regexp.MustCompile(`\d+`)
	fmt.Println()
	for _, line := range lines {
		numeric, _ := strconv.Atoi(re.FindString(line))
		result += numeric * numpad_type.shortestSequence(line, 25)
	}

	return result
}
