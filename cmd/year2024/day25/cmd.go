package day25

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day25",
	Short: "day25",
	Long:  "day25",
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

type mechanismList struct {
	keys  [][5]int
	locks [][5]int
}

func newMechanismList(input string) mechanismList {
	m := mechanismList{
		keys:  [][5]int{},
		locks: [][5]int{},
	}

	combinations := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, combination := range combinations {
		lines := strings.Split(combination, "\n")

		values := [5]int{0, 0, 0, 0, 0}
		for _, line := range lines {
			symbols := strings.Split(line, "")
			for x, symbol := range symbols {
				if symbol == "#" {
					values[x]++
				}
			}
		}

		if lines[0] == "#####" {
			m.locks = append(m.locks, values)
		} else if lines[len(lines)-1] == "#####" {
			m.keys = append(m.keys, values)
		} else {
			panic("this is not a key and not a lock")
		}
	}

	return m
}

func (m mechanismList) matches() int {
	res := 0
	for _, key := range m.keys {
		for _, lock := range m.locks {
			fits := true

			for i := 0; i < len(key); i++ {
				if key[i]+lock[i] > 7 {
					fits = false
				}
			}

			if fits {
				res++
			}
		}
	}

	return res
}

func Part1(input string) int {
	m := newMechanismList(input)
	return m.matches()
}

func Part2(input string) int {
	return 0
}
