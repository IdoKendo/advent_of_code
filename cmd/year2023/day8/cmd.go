package day8

import (
	"fmt"
	"math"
	"os"
	"regexp"
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

func parse(input string) (map[string][2]string, []int) {
	lines := strings.Split(input, "\n")
	chars := strings.Split(lines[0], "")
	nodes := lines[2 : len(lines)-1]
	instructions := make([]int, len(chars))
	for i, r := range chars {
		switch r {
		case "L":
			instructions[i] = 0
		case "R":
			instructions[i] = 1
		}
	}
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	network := make(map[string][2]string)
	for _, node := range nodes {
		n := re.FindStringSubmatch(node)
		network[n[1]] = [2]string{n[2], n[3]}
	}

	return network, instructions
}

func Part1(input string) int {
	network, instructions := parse(input)
	stepsTaken := 0
	position := "AAA"
	for position != "ZZZ" {
		i := instructions[stepsTaken%len(instructions)]
		stepsTaken++
		position = network[position][i]
	}

	return stepsTaken
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	result := slice[0]
	for _, num := range slice[1:] {
		result = int(math.Abs(float64(result*num))) / gcd(result, num)
	}
	return result
}

func Part2(input string) int {
	network, instructions := parse(input)
	stepsTaken := 0
	var ghosts []string
	for position := range network {
		if strings.HasSuffix(position, "A") {
			ghosts = append(ghosts, position)
		}
	}

	ghostPaths := make([]int, len(ghosts))
	for {
		product := 1
		for _, num := range ghostPaths {
			product *= num
		}
		if product > 0 {
			break
		}

		i := instructions[stepsTaken%len(instructions)]
		stepsTaken++

		for j := range ghosts {
			ghosts[j] = network[ghosts[j]][i]
			if strings.HasSuffix(ghosts[j], "Z") {
				ghostPaths[j] = stepsTaken
			}
		}
	}

	return lcm(ghostPaths)
}
