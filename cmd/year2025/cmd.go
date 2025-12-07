package year2025

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2025/day1"
	"github.com/idokendo/aoc/cmd/year2025/day2"
	"github.com/idokendo/aoc/cmd/year2025/day3"
	"github.com/idokendo/aoc/cmd/year2025/day4"
	"github.com/idokendo/aoc/cmd/year2025/day5"
	"github.com/idokendo/aoc/cmd/year2025/day6"
	"github.com/idokendo/aoc/cmd/year2025/day7"
	"github.com/idokendo/aoc/cmd/year2025/day8"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2025",
	Short: "2025",
	Long:  "2025",
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
}
