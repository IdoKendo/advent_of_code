package day16

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day16",
	Short: "day16",
	Long:  "day16",
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

type point struct {
	x, y int
}
type position struct {
	point     point
	direction direction
}

type priorityQueueItem struct {
	dist int
	pos  position
	idx  int
}

type priorityQueue []priorityQueueItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(priorityQueueItem)
	item.idx = len(*pq)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type maze struct {
	grid  [][]rune
	start position
	end   position
}

func newMaze(lines []string) maze {
	var grid [][]rune
	var start, end position

	for y, line := range lines {
		row := []rune(strings.TrimSpace(line))
		grid = append(grid, row)
		for x, ch := range row {
			p := point{x, y}
			switch ch {
			case 'S':
				start = position{point: p}
			case 'E':
				end = position{point: p}
			}
		}
	}

	return maze{grid, start, end}
}

type direction string

var (
	east  direction = "E"
	west  direction = "W"
	north direction = "N"
	south direction = "S"
)

func (d direction) flip() direction {
	flip := map[direction]direction{east: west, west: east, north: south, south: north}
	return flip[d]
}

func dijkstra(grid [][]rune, starts []position) map[position]int {
	delta := map[direction]point{
		east:  {0, 1},
		west:  {0, -1},
		north: {-1, 0},
		south: {1, 0},
	}

	dist := make(map[position]int)
	pq := &priorityQueue{}
	heap.Init(pq)

	for _, start := range starts {
		dist[start] = 0
		heap.Push(pq, priorityQueueItem{dist: 0, pos: start})
	}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(priorityQueueItem)
		curr := item.pos
		d := item.dist

		if currentDist, exists := dist[curr]; exists && currentDist < d {
			continue
		}

		for _, nextDir := range []direction{east, west, north, south} {
			if nextDir == curr.direction {
				continue
			}
			nextPos := position{point: curr.point, direction: nextDir}
			if currentDist, exists := dist[nextPos]; !exists || currentDist > d+1000 {
				dist[nextPos] = d + 1000
				heap.Push(pq, priorityQueueItem{dist: d + 1000, pos: nextPos})
			}
		}

		dx, dy := delta[curr.direction].x, delta[curr.direction].y
		nextX, nextY := curr.point.x+dx, curr.point.y+dy
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] != '#' {
			nextPos := position{point{nextX, nextY}, curr.direction}
			if currentDist, exists := dist[nextPos]; !exists || currentDist > d+1 {
				dist[nextPos] = d + 1
				heap.Push(pq, priorityQueueItem{dist: d + 1, pos: nextPos})
			}
		}
	}

	return dist
}

func Part1(input string) int {
	lines := strings.Split(string(input), "\n")
	m := newMaze(lines)
	starts := []position{{m.start.point, east}}
	solutions := dijkstra(m.grid, starts)
	minCost := math.MaxInt
	for s, cost := range solutions {
		if s.point.x == m.end.point.x && s.point.y == m.end.point.y && cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func Part2(input string) int {
	lines := strings.Split(string(input), "\n")
	m := newMaze(lines)
	fromStart := dijkstra(m.grid, []position{{m.start.point, east}})
	fromEnd := dijkstra(m.grid, []position{
		{m.end.point, east},
		{m.end.point, west},
		{m.end.point, north},
		{m.end.point, south},
	})
	minCost := Part1(input)
	result := make(map[point]bool)

	for pos, d1 := range fromStart {
		opposite := position{pos.point, pos.direction.flip()}
		if d2, exists := fromEnd[opposite]; exists {
			if d1+d2 == minCost {
				result[pos.point] = true
			}
		}
	}

	return len(result)
}
