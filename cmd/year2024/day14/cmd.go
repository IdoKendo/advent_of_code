package day14

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
	Use:   "day14",
	Short: "day14",
	Long:  "day14",
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

	result := Part1(string(content), 103, 101)
	fmt.Println("Part 1 result: ", result)

	content, err = os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input2.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result = Part2(string(content), 103, 101)
	fmt.Println("Part 2 result: ", result)
}

type point struct {
	x, y int
}

func wrap(value, limit int) int {
	return (value%limit + limit) % limit
}

type robot struct {
	position point
	velocity point
}

func (r *robot) move(width, height int) {
	r.position.x = wrap(r.position.x+r.velocity.x, width)
	r.position.y = wrap(r.position.y+r.velocity.y, height)
}

func parseRobotSlice(input string) []robot {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	robots := make([]robot, len(lines))
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])
		robots[i] = robot{
			position: point{
				x: px,
				y: py,
			},
			velocity: point{
				x: vx,
				y: vy,
			},
		}
	}
	return robots
}

func safetyFactor(robots []robot, height, width int) int {
	quadrants := make([]int, 4)
	midRow := height / 2
	midCol := width / 2

	for _, robot := range robots {
		if robot.position.x < midCol {
			if robot.position.y < midRow {
				quadrants[0]++
			} else if robot.position.y > midRow {
				quadrants[1]++
			}
		} else if robot.position.x > midCol {
			if robot.position.y < midRow {
				quadrants[2]++
			} else if robot.position.y > midRow {
				quadrants[3]++
			}
		}
	}

	product := 1
	for _, q := range quadrants {
		product *= q
	}

	return product
}

func Part1(input string, height, width int) int {
	robots := parseRobotSlice(input)
	for s := 0; s < 100; s++ {
		for i := range robots {
			r := &robots[i]
			r.move(width, height)
		}
	}

	return safetyFactor(robots, height, width)
}

func printMap(robots []robot, height, width int) {
	graph := make([][]int, height)
	for i := range height {
		graph[i] = make([]int, width)
	}

	positionMap := make(map[point]int)
	for _, r := range robots {
		positionMap[r.position]++
	}

	for y, row := range graph {
		for x := range row {
			graph[y][x] += positionMap[point{x, y}]
		}
	}

	fmt.Println()
	for _, row := range graph {
		for _, v := range row {
			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Part2(input string, height, width int) int {
	robots := parseRobotSlice(input)
	sf := math.MaxInt
	elapsedTime := 0
	for i := 0; i < 10000; i++ {
		newSf := safetyFactor(robots, height, width)
		if newSf < sf {
			printMap(robots, height, width)
			sf = newSf
			elapsedTime = i
		}
		for j := range robots {
			r := &robots[j]
			r.move(width, height)
		}
	}
	return elapsedTime
}
