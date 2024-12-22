package day22

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day22",
	Short: "day22",
	Long:  "day22",
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

func step(n int) int {
	n = ((n * 64) ^ n) % 16777216
	n = ((n / 32) ^ n) % 16777216
	n = ((n * 2048) ^ n) % 16777216
	return n
}

type simulator struct {
	memo map[int]int
}

func newSimulator() simulator {
	memo := make(map[int]int)
	return simulator{memo}
}

func (s simulator) run(secret, generations int) int {
	if generations == 0 {
		return secret
	}

	generations--
	if val, exists := s.memo[secret]; exists {
		return s.run(val, generations)
	}

	s.memo[secret] = step(secret)

	return s.run(s.memo[secret], generations)
}

func (s simulator) prices(secret int) []int {
	s.run(secret, 2000)
	result := make([]int, 2000)
	result[0] = secret % 10
	prev := secret
	for i := 1; i < 2000; i++ {
		result[i] = s.memo[prev] % 10
		prev = s.memo[prev]
	}

	return result
}

func diffs(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}

	result := make([]int, len(slice))
	result[0] = 0

	for i := 1; i < len(slice); i++ {
		result[i] = slice[i] - slice[i-1]
	}

	return result
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	sim := newSimulator()
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		result += sim.run(secret, 2000)
	}

	return result
}

type monkey struct {
	secret  int
	prices  []int
	changes []int
}

func (m monkey) sequences() map[[4]int]int {
	s := make(map[[4]int]int)
	for i := 3; i < len(m.prices); i++ {
		key := [4]int{m.changes[i-3], m.changes[i-2], m.changes[i-1], m.changes[i]}
		if _, exists := s[key]; exists {
			continue
		}
		s[key] = m.prices[i]
	}
	return s
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sim := newSimulator()
	monkeys := make([]monkey, len(lines))
	for i, line := range lines {
		secret, _ := strconv.Atoi(line)
		p := sim.prices(secret)
		monkeys[i] = monkey{
			secret:  secret,
			prices:  p,
			changes: diffs(p),
		}
	}

	sequences := make(map[[4]int]int)
	for _, m := range monkeys {
		for seq, bananas := range m.sequences() {
			sequences[seq] += bananas
		}
	}

	result := 0
	for _, bananas := range sequences {
		result = max(result, bananas)
	}

	return result
}
