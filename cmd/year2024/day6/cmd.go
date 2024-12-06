package day6

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day6",
	Short: "day6",
	Long:  "day6",
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

type direction [2]int

var (
	north direction = [2]int{0, -1}
	east  direction = [2]int{1, 0}
	south direction = [2]int{0, 1}
	west  direction = [2]int{-1, 0}
)

type guard struct {
	x   int
	y   int
	dir direction
}

func (g *guard) turn() {
	switch g.dir {
	case north:
		g.dir = east
	case east:
		g.dir = south
	case south:
		g.dir = west
	case west:
		g.dir = north
	}
}

func (g *guard) print(w io.Writer) {
	switch g.dir {
	case north:
		fmt.Fprint(w, "^")
	case east:
		fmt.Fprint(w, ">")
	case south:
		fmt.Fprint(w, "v")
	case west:
		fmt.Fprint(w, "<")
	}
}

type state struct {
	dir direction
	x   int
	y   int
}

type Map struct {
	g         guard
	size      [2]int
	obstacles map[[2]int]bool
	visited   map[[2]int]bool
	states    map[state]bool
}

func newMap(size [2]int) *Map {
	return &Map{
		size:      size,
		obstacles: map[[2]int]bool{},
		visited:   map[[2]int]bool{},
		states:    map[state]bool{},
	}
}

func newMapFrom(input string) *Map {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	width := len(lines)
	height := len(lines[0])

	m := newMap([2]int{width, height})
	for i := 0; i < width; i++ {
		pos := strings.Split(lines[i], "")
		for j := 0; j < height; j++ {
			switch pos[j] {
			case "#":
				m.obstacles[[2]int{j, i}] = true
			case ".":
				continue
			default:
				m.initGuard(i, j, pos[j])
			}
		}
	}

	return m
}

func (m *Map) initGuard(y, x int, sign string) {
	var direction direction
	switch sign {
	case "^":
		direction = north
	case ">":
		direction = east
	case "v":
		direction = south
	case "<":
		direction = west
	}
	m.g = guard{x, y, direction}
}

func (m *Map) printMap(w io.Writer) {
	for i := 0; i < m.size[0]; i++ {
		for j := 0; j < m.size[1]; j++ {
			if m.obstacles[[2]int{j, i}] {
				fmt.Fprint(w, "#")
			} else if m.g.x == j && m.g.y == i {
				m.g.print(w)
			} else if m.visited[[2]int{j, i}] {
				fmt.Fprint(w, "X")

			} else {
				fmt.Fprint(w, ".")
			}
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w)
}

func (m *Map) tick() {
	m.states[state{dir: m.g.dir, x: m.g.x, y: m.g.y}] = true

	nextX := m.g.x + m.g.dir[0]
	nextY := m.g.y + m.g.dir[1]
	if m.obstacles[[2]int{nextX, nextY}] {
		m.g.turn()
		return
	}
	if !m.visited[[2]int{m.g.x, m.g.y}] {
		if m.g.x < m.size[0] && m.g.y < m.size[1] && m.g.x >= 0 && m.g.y >= 0 {
			m.visited[[2]int{m.g.x, m.g.y}] = true
		}
	}

	m.g.x = nextX
	m.g.y = nextY
}

func (m *Map) finish() {
	for m.g.x <= m.size[0] && m.g.y <= m.size[1] && m.g.x >= 0 && m.g.y >= 0 {
		m.tick()
	}
}

func (m *Map) isInfinite() bool {
	for {
		if m.g.x < 0 || m.g.x >= m.size[1] || m.g.y < 0 || m.g.y >= m.size[0] {
			return false
		}

		s := state{dir: m.g.dir, x: m.g.x, y: m.g.y}
		if m.states[s] {
			return true
		}
		m.states[s] = true

		m.tick()
	}
}

func Part1(input string) int {
	m := newMapFrom(input)
	m.finish()
	return len(m.visited)
}

type result struct {
	x, y     int
	infinite bool
}

func Part2(input string) int {
	m := newMapFrom(input)
	m.finish()
	optionals := m.visited
	blocks := 0

	results := make(chan result)
	wg := sync.WaitGroup{}
	wg.Add(len(optionals))

	for option := range optionals {
		i := option[0]
		j := option[1]

		go func(x, y int) {
			defer wg.Done()
			m := newMapFrom(input)
			p := [2]int{x, y}
			if m.obstacles[p] {
				results <- result{x, y, false}
				return
			}
			m.obstacles[p] = true
			results <- result{x, y, m.isInfinite()}
		}(i, j)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.infinite {
			blocks++
		}
	}

	return blocks
}
