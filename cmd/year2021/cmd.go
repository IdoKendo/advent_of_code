package year2021

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2021/day1"
	"github.com/idokendo/aoc/cmd/year2021/day2"
	"github.com/idokendo/aoc/cmd/year2021/day3"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2021",
	Short: "2021",
	Long:  "2021",
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
}
