package year2024

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2024/day1"
	"github.com/idokendo/aoc/cmd/year2024/day10"
	"github.com/idokendo/aoc/cmd/year2024/day11"
	"github.com/idokendo/aoc/cmd/year2024/day12"
	"github.com/idokendo/aoc/cmd/year2024/day13"
	"github.com/idokendo/aoc/cmd/year2024/day14"
	"github.com/idokendo/aoc/cmd/year2024/day15"
	"github.com/idokendo/aoc/cmd/year2024/day16"
	"github.com/idokendo/aoc/cmd/year2024/day17"
	"github.com/idokendo/aoc/cmd/year2024/day18"
	"github.com/idokendo/aoc/cmd/year2024/day19"
	"github.com/idokendo/aoc/cmd/year2024/day2"
	"github.com/idokendo/aoc/cmd/year2024/day20"
	"github.com/idokendo/aoc/cmd/year2024/day21"
	"github.com/idokendo/aoc/cmd/year2024/day22"
	"github.com/idokendo/aoc/cmd/year2024/day23"
	"github.com/idokendo/aoc/cmd/year2024/day24"
	"github.com/idokendo/aoc/cmd/year2024/day25"
	"github.com/idokendo/aoc/cmd/year2024/day3"
	"github.com/idokendo/aoc/cmd/year2024/day4"
	"github.com/idokendo/aoc/cmd/year2024/day5"
	"github.com/idokendo/aoc/cmd/year2024/day6"
	"github.com/idokendo/aoc/cmd/year2024/day7"
	"github.com/idokendo/aoc/cmd/year2024/day8"
	"github.com/idokendo/aoc/cmd/year2024/day9"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2024",
	Short: "2024",
	Long:  "2024",
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
	Cmd.AddCommand(day13.Cmd)
	Cmd.AddCommand(day14.Cmd)
	Cmd.AddCommand(day15.Cmd)
	Cmd.AddCommand(day16.Cmd)
	Cmd.AddCommand(day17.Cmd)
	Cmd.AddCommand(day18.Cmd)
	Cmd.AddCommand(day19.Cmd)
	Cmd.AddCommand(day20.Cmd)
	Cmd.AddCommand(day21.Cmd)
	Cmd.AddCommand(day22.Cmd)
	Cmd.AddCommand(day23.Cmd)
	Cmd.AddCommand(day24.Cmd)
	Cmd.AddCommand(day25.Cmd)
}
