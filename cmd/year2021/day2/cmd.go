package day2

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  "day2",
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

var directions = map[string]image.Point{
	"forward": {1, 0},
	"up":      {0, -1},
	"down":    {0, 1},
}

func pilotSubmarine(input string, withAim bool) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pos := image.Point{X: 0, Y: 0}
	aim := 0
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		direction := directions[parts[0]]
		distance, _ := strconv.Atoi(parts[1])
		if !withAim {
			pos = pos.Add(image.Point{
				X: distance * direction.X,
				Y: distance * direction.Y,
			})
		} else {
			switch direction {
			case directions["forward"]:
				pos = pos.Add(image.Point{X: distance, Y: distance * aim})
			case directions["up"]:
				aim -= distance
			case directions["down"]:
				aim += distance
			}
		}
	}

	return pos.X * pos.Y
}

func Part1(input string) int {
	return pilotSubmarine(input, false)
}

func Part2(input string) int {
	return pilotSubmarine(input, true)
}
