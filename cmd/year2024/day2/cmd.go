package day2

import (
	"fmt"
	"os"
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

func isReportSafe(levels []string) bool {
	level1, _ := strconv.Atoi(levels[0])
	level2, _ := strconv.Atoi(levels[1])
	direction := 0
	safe := true
	if level1 > level2 {
		direction = -1
	} else {
		direction = 1
	}
	for i := 1; i < len(levels); i++ {
		prev, _ := strconv.Atoi(levels[i-1])
		this, _ := strconv.Atoi(levels[i])
		diff := (this - prev) * direction
		if diff > 3 || diff < 1 {
			safe = false
			break
		}
	}
	return safe
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	reports := lines[:len(lines)-1]
	safeReports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")
		safe := isReportSafe(levels)
		if safe {
			safeReports++
		}
	}

	return safeReports
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	reports := lines[:len(lines)-1]
	safeReports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")
		safe := isReportSafe(levels)
		if safe {
			safeReports++
			continue
		}
		for i := 0; !safe && i < len(levels); i++ {
			l := make([]string, len(levels))
			copy(l, levels)
			poppedLevels := append(l[:i], l[i+1:]...)
			safe = isReportSafe(poppedLevels)
		}
		if safe {
			safeReports++
		}
	}

	return safeReports
}
