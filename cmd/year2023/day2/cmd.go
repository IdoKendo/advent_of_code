package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  "day2",
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

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	redRe := regexp.MustCompile(`(\d+) red`)
	greenRe := regexp.MustCompile(`(\d+) green`)
	blueRe := regexp.MustCompile(`(\d+) blue`)
	gameRe := regexp.MustCompile(`(\d+)`)
	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		gameID := splitLine[0]
		l := splitLine[1]
		possible := true
		for _, s := range strings.Split(l, ";") {
			red := redRe.FindStringSubmatch(s)
			if len(red) > 0 {
				n, _ := strconv.Atoi(red[1])
				if n > 12 {
					possible = false
				}
			}
			green := greenRe.FindStringSubmatch(s)
			if len(green) > 0 {
				n, _ := strconv.Atoi(green[1])
				if n > 13 {
					possible = false
				}
			}
			blue := blueRe.FindStringSubmatch(s)
			if len(blue) > 0 {
				n, _ := strconv.Atoi(blue[1])
				if n > 14 {
					possible = false
				}
			}
		}
		if possible {
			gameID := gameRe.FindAllString(gameID, -1)
			n, _ := strconv.Atoi(gameID[0])
			result += n
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0
	redRe := regexp.MustCompile(`(\d+) red`)
	greenRe := regexp.MustCompile(`(\d+) green`)
	blueRe := regexp.MustCompile(`(\d+) blue`)
	for _, line := range lines {
		minRed := 0
		minGreen := 0
		minBlue := 0
		splitLine := strings.Split(line, ":")
		l := splitLine[1]
		for _, s := range strings.Split(l, ";") {
			red := redRe.FindStringSubmatch(s)
			if len(red) > 0 {
				n, _ := strconv.Atoi(red[1])
				minRed = max(minRed, n)
			}
			green := greenRe.FindStringSubmatch(s)
			if len(green) > 0 {
				n, _ := strconv.Atoi(green[1])
				minGreen = max(minGreen, n)
			}
			blue := blueRe.FindStringSubmatch(s)
			if len(blue) > 0 {
				n, _ := strconv.Atoi(blue[1])
				minBlue = max(minBlue, n)
			}
		}
		result += minRed * minGreen * minBlue
	}
	return result
}
