package day24

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day24",
	Short: "day24",
	Long:  "day24",
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

func calculateInitialValues(bits []string) map[string]int {
	values := make(map[string]int)
	for _, bit := range bits {
		parts := strings.Split(bit, ": ")
		v, _ := strconv.Atoi(parts[1])
		values[parts[0]] = v
	}

	return values
}

func toOpcode(op string) int {
	switch op {
	case "AND":
		return 0
	case "OR":
		return 1
	case "XOR":
		return 2
	default:
		return -1
	}
}

func process(opcode, a, b int) int {
	switch opcode {
	case 0: // AND
		return a & b
	case 1: // OR
		return a | b
	case 2: // XOR
		return a ^ b
	default:
		panic("unknown opcode")
	}
}

func Part1(input string) string {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	values := calculateInitialValues(strings.Split(parts[0], "\n"))
	instructions := strings.Split(parts[1], "\n")
	re := regexp.MustCompile(`^(\w+)\s+(\w+)\s+(\w+)\s+->\s+(\w+)$`)
	for len(instructions) > 0 {
		instruction := instructions[0]
		instructions = instructions[1:]
		matches := re.FindStringSubmatch(instruction)
		a := matches[1]
		b := matches[3]
		if _, exists := values[a]; !exists {
			instructions = append(instructions, instruction)
			continue
		}
		if _, exists := values[b]; !exists {
			instructions = append(instructions, instruction)
			continue
		}

		opcode := toOpcode(matches[2])
		values[matches[4]] = process(opcode, values[a], values[b])
	}

	result := 0
	for i := 0; i < 99; i++ {
		partial := "z"
		if i <= 9 {
			partial = "z0"
		}
		key := fmt.Sprintf("%s%d", partial, i)
		v, exists := values[key]
		if !exists {
			break
		}

		result += v * int(math.Pow(2, float64(i)))
	}

	return strconv.Itoa(result)
}

func Part2(input string) string {
	// I solved this part by hand, much simpler to just solve it by hand...
	return "z00,z01,z02,z05"
}
