package year2023

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2023/day1"
	"github.com/idokendo/aoc/cmd/year2023/day10"
	"github.com/idokendo/aoc/cmd/year2023/day11"
	"github.com/idokendo/aoc/cmd/year2023/day12"
	"github.com/idokendo/aoc/cmd/year2023/day2"
	"github.com/idokendo/aoc/cmd/year2023/day3"
	"github.com/idokendo/aoc/cmd/year2023/day4"
	"github.com/idokendo/aoc/cmd/year2023/day5"
	"github.com/idokendo/aoc/cmd/year2023/day6"
	"github.com/idokendo/aoc/cmd/year2023/day7"
	"github.com/idokendo/aoc/cmd/year2023/day8"
	"github.com/idokendo/aoc/cmd/year2023/day9"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "2023",
	Long:  "2023",
}

func execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	Cmd.AddCommand(day1.Cmd)
	Cmd.AddCommand(day2.Cmd)
	Cmd.AddCommand(day3.Cmd)
	Cmd.AddCommand(day4.Cmd)
	Cmd.AddCommand(day5.Cmd)
	Cmd.AddCommand(day6.Cmd)
	Cmd.AddCommand(day7.Cmd)
	Cmd.AddCommand(day8.Cmd)
	Cmd.AddCommand(day9.Cmd)
	Cmd.AddCommand(day10.Cmd)
	Cmd.AddCommand(day11.Cmd)
	Cmd.AddCommand(day12.Cmd)
}
