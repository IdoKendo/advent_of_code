package day9

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "day9",
	Long:  "day9",
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

type Perimeter map[image.Point]bool

func newPerimeter(tiles []image.Point) Perimeter {
	perimeter := make(Perimeter)
	current := tiles[0]

	for _, tile := range tiles[1:] {
		for current != tile {
			perimeter[current] = true

			if current.X != tile.X {
				if current.X < tile.X {
					current.X++
				} else {
					current.X--
				}
				continue
			}

			if current.Y < tile.Y {
				current.Y++
			} else {
				current.Y--
			}
		}
	}

	first := tiles[0]

	for current != first {
		perimeter[current] = true

		if current.X != first.X {
			if current.X < first.X {
				current.X++
			} else {
				current.X--
			}
			continue
		}

		if current.Y < first.Y {
			current.Y++
		} else {
			current.Y--
		}
	}

	return perimeter
}

func (p Perimeter) containsRectangle(a, b image.Point) bool {
	minX := min(a.X, b.X)
	minY := min(a.Y, b.Y)
	maxX := max(a.X, b.X)
	maxY := max(a.Y, b.Y)

	for point := range p {
		// left edge
		if point.Y == minY && point.X > minX && point.X < maxX {
			if _, ok := p[image.Point{point.X, point.Y + 1}]; ok {
				return false
			}
		}
		// right edge
		if point.Y == maxY && point.X > minX && point.X < maxX {
			if _, ok := p[image.Point{point.X, point.Y - 1}]; ok {
				return false
			}
		}
		// top edge
		if point.X == minX && point.Y > minY && point.Y < maxY {
			if _, ok := p[image.Point{point.X + 1, point.Y}]; ok {
				return false
			}
		}
		// bottom edge
		if point.X == maxX && point.Y > minY && point.Y < maxY {
			if _, ok := p[image.Point{point.X - 1, point.Y}]; ok {
				return false
			}
		}
	}
	return true
}

func tiles(input string) []image.Point {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	tiles := make([]image.Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles[i] = image.Point{x, y}
	}

	return tiles
}

func Part1(input string) int {
	redTiles := tiles(input)
	largestRectangleArea := 0
	for _, a := range redTiles {
		for _, b := range redTiles {
			if a == b {
				continue
			}
			y := math.Abs(float64(b.Y-a.Y)) + 1
			x := math.Abs(float64(b.X-a.X)) + 1
			area := int(x * y)
			largestRectangleArea = max(area, largestRectangleArea)
		}
	}
	return largestRectangleArea
}

func Part2(input string) int {
	redTiles := tiles(input)
	perimeter := newPerimeter(redTiles)
	memo := make(map[string]bool)
	largestRectangleArea := 0
	for _, a := range redTiles {
		for _, b := range redTiles {
			if a == b {
				continue
			}
			minX := min(a.X, b.X)
			minY := min(a.Y, b.Y)
			maxX := max(a.X, b.X)
			maxY := max(a.Y, b.Y)
			y := maxY - minY + 1
			x := maxX - minX + 1
			area := x * y
			if area <= largestRectangleArea {
				continue
			}
			key := fmt.Sprintf("%d,%d,%d,%d", minX, minY, maxX, maxY)
			contained, seen := memo[key]
			if !seen {
				contained = perimeter.containsRectangle(a, b)
				memo[key] = contained
			}
			if !contained {
				continue
			}
			largestRectangleArea = area
		}
	}
	return largestRectangleArea
}
