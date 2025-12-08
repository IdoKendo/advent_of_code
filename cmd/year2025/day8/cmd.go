package day8

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

	result := Part1(string(content), 1000)
	fmt.Println("Part 1 result: ", result)

	content, err = os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input2.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result = Part2(string(content), 1000)
	fmt.Println("Part 2 result: ", result)
}

type JunctionBox struct {
	X, Y, Z float64
}

func (j JunctionBox) Euclidean(other JunctionBox) float64 {
	return math.Sqrt((other.X-j.X)*(other.X-j.X) + (other.Y-j.Y)*(other.Y-j.Y) + (other.Z-j.Z)*(other.Z-j.Z))
}

type JunctionsList []JunctionBox

func newJunctionsList(input string) JunctionsList {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	junctions := make([]JunctionBox, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		junctions[i] = JunctionBox{
			X: float64(x),
			Y: float64(y),
			Z: float64(z),
		}
	}

	return junctions
}

type Pair struct {
	d2 float64
	i  int
	j  int
}

type DSU struct {
	parent []int
	size   []int
}

func newDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := range p {
		p[i] = i
		s[i] = 1
	}
	return &DSU{parent: p, size: s}
}

func (d *DSU) Find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *DSU) Union(a, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

func Part1(input string, limit int) int {
	junctions := newJunctionsList(input)
	length := len(junctions)
	pairs := make([]Pair, 0, length*length/2)

	for i := range length {
		for j := i + 1; j < length; j++ {
			d2 := junctions[i].Euclidean(junctions[j])
			pairs = append(pairs, Pair{d2, i, j})
		}
	}

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].d2 < pairs[b].d2
	})

	dsu := newDSU(length)

	for n := 0; n < len(pairs) && n < limit; n++ {
		p := pairs[n]
		dsu.Union(p.i, p.j)
	}

	sizes := make(map[int]int)
	for i := range length {
		r := dsu.Find(i)
		sizes[r]++
	}

	arr := make([]int, 0, len(sizes))
	for _, s := range sizes {
		arr = append(arr, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	result := 1
	for i := 0; i < len(arr) && i < 3; i++ {
		result *= arr[i]
	}

	return result
}

func Part2(input string, limit int) int {
	junctions := newJunctionsList(input)
	length := len(junctions)
	pairs := make([]Pair, 0, length*length/2)

	for i := range length {
		for j := i + 1; j < length; j++ {
			d2 := junctions[i].Euclidean(junctions[j])
			pairs = append(pairs, Pair{d2, i, j})
		}
	}

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].d2 < pairs[b].d2
	})

	dsu := newDSU(length)

	var i int
	var j int

	for _, p := range pairs {
		if dsu.Union(p.i, p.j) {
			length--
			i = p.i
			j = p.j
			if length == 1 {
				break
			}
		}
	}

	return int(junctions[i].X * junctions[j].X)
}
