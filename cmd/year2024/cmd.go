package year2024

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2024/day1"
	"github.com/idokendo/aoc/cmd/year2024/day2"
	"github.com/idokendo/aoc/cmd/year2024/day3"
	"github.com/idokendo/aoc/cmd/year2024/day4"
	"github.com/idokendo/aoc/cmd/year2024/day5"
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
}
