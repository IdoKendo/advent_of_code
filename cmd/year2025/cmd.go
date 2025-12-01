package year2025

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2025/day1"
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
}
