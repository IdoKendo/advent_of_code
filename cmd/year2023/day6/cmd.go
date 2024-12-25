package day6

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

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

func parseRaces(input string) []Race {
	re := regexp.MustCompile(`Time:\s*(.+)\nDistance:\s*(.+)`)
	matches := re.FindStringSubmatch(input)

	timeNumbers := regexp.MustCompile(`\d+`).FindAllString(matches[1], -1)
	distNumbers := regexp.MustCompile(`\d+`).FindAllString(matches[2], -1)

	races := make([]Race, len(timeNumbers))
	for i := range timeNumbers {
		time, _ := strconv.Atoi(timeNumbers[i])
		dist, _ := strconv.Atoi(distNumbers[i])
		races[i] = Race{
			Time:   time,
			Record: dist,
		}
	}

	return races
}

type Race struct {
	Time   int
	Record int
}

func (r Race) IsBeatenByHoldingFor(milliseconds int) bool {
	return (r.Time-milliseconds)*milliseconds > r.Record
}

func Part1(input string) int {
	races := parseRaces(input)
	optionsProduct := 1
	for _, race := range races {
		options := 0
		for i := range race.Time {
			if race.IsBeatenByHoldingFor(i) {
				options += 1
			}
		}
		optionsProduct *= options
	}
	return optionsProduct
}

func Part2(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	race := parseRaces(input)[0]
	options := 0
	for i := range race.Time {
		if race.IsBeatenByHoldingFor(i) {
			options += 1
		}
	}
	return options
}
