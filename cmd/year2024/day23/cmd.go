package day23

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day23",
	Short: "day23",
	Long:  "day23",
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

func getPairs(input string) map[string][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	pairs := make(map[string][]string)
	for _, line := range lines {
		p := strings.Split(line, "-")
		a, b := p[0], p[1]
		pairs[a] = append(pairs[a], b)
		pairs[b] = append(pairs[b], a)
	}

	return pairs
}

func Part1(input string) string {
	pairs := getPairs(input)
	sets := make(map[[3]string]bool)
	for a, neighbors := range pairs {
		for i, b := range neighbors {
			for _, c := range neighbors[i+1:] {
				for _, d := range pairs[b] {
					if c == d {
						s := []string{a, b, c}
						sort.Strings(s)
						sets[[3]string(s)] = true
					}
				}
			}
		}
	}

	result := 0

	for set := range sets {
		if strings.HasPrefix(set[0], "t") || strings.HasPrefix(set[1], "t") || strings.HasPrefix(set[2], "t") {
			result++
		}
	}

	return strconv.Itoa(result)
}

func bronKerbosch(graph map[string][]string, r, p, x map[string]bool, maxClique *[]string) {
	if len(p) == 0 && len(x) == 0 {
		if len(r) > len(*maxClique) {
			*maxClique = make([]string, 0, len(r))
			for node := range r {
				*maxClique = append(*maxClique, node)
			}
		}
		return
	}

	pCopy := make(map[string]bool)
	for node := range p {
		pCopy[node] = true
	}

	for v := range pCopy {
		newR := make(map[string]bool)
		for key := range r {
			newR[key] = true
		}
		newP := make(map[string]bool)
		newX := make(map[string]bool)

		newR[v] = true

		for _, neighbor := range graph[v] {
			if p[neighbor] {
				newP[neighbor] = true
			}
			if x[neighbor] {
				newX[neighbor] = true
			}
		}

		bronKerbosch(graph, newR, newP, newX, maxClique)

		delete(p, v)
		x[v] = true
	}
}

func Part2(input string) string {
	pairs := getPairs(input)

	r := make(map[string]bool)
	p := make(map[string]bool)
	x := make(map[string]bool)

	for pc := range pairs {
		p[pc] = true
	}

	largestSet := []string{}
	bronKerbosch(pairs, r, p, x, &largestSet)
	sort.Strings(largestSet)

	return strings.Join(largestSet, ",")
}
