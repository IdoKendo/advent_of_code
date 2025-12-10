package day10

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day10",
	Short: "day10",
	Long:  "day10",
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

type LightState struct {
	status     string
	buttons    []int
	pressCount int
}

type Machine struct {
	indicationLights   string
	buttons            [][]int
	joltageRequirement []int
}

func newMachine(s string) Machine {
	m := Machine{}

	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(s)
	m.indicationLights = matches[1]

	re = regexp.MustCompile(`\((.*?)\)`)
	allMatches := re.FindAllStringSubmatch(s, -1)
	for _, match := range allMatches {
		content := match[1]
		if content == "" {
			m.buttons = append(m.buttons, nil)
			continue
		}
		parts := strings.Split(content, ",")
		var nums []int
		for _, p := range parts {
			n, _ := strconv.Atoi(strings.TrimSpace(p))
			nums = append(nums, n)
		}
		m.buttons = append(m.buttons, nums)
	}

	re = regexp.MustCompile(`\{(.*?)\}`)
	matches = re.FindStringSubmatch(s)
	for _, p := range strings.Split(matches[1], ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(p))
		m.joltageRequirement = append(m.joltageRequirement, n)
	}

	return m
}

func (m Machine) ConfigureIndicationLights() int {
	var status string
	for range len(m.indicationLights) {
		status += "."
	}
	pressCount := math.MaxInt
	visited := make(map[string]int)
	q := []LightState{}
	for _, buttons := range m.buttons {
		q = append(q, LightState{status, buttons, 0})
		visited[status] = 0
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.pressCount > visited[curr.status] {
			continue
		}
		if curr.status == m.indicationLights {
			pressCount = min(pressCount, curr.pressCount)
			continue
		}
		for _, buttons := range m.buttons {
			if slices.Equal(buttons, curr.buttons) {
				continue
			}
			next := LightState{
				pressButton(curr.status, buttons),
				buttons,
				curr.pressCount + 1,
			}
			if v, ok := visited[next.status]; !ok || next.pressCount < v {
				visited[next.status] = next.pressCount
				q = append(q, next)
			}
		}
	}

	return pressCount
}

func pressButton(status string, button []int) string {
	for _, idx := range button {
		b := []byte(status)

		if b[idx] == '.' {
			b[idx] = '#'
		} else {
			b[idx] = '.'
		}

		status = string(b)
	}

	return status
}

func (m Machine) ConfigureJoltage() int {
	return 11 // TODO: implement
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pressCount := 0
	for _, line := range lines {
		m := newMachine(line)
		pressCount += m.ConfigureIndicationLights()
	}
	return pressCount
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pressCount := 0
	for _, line := range lines {
		m := newMachine(line)
		pressCount += m.ConfigureJoltage()
	}
	return pressCount
}
