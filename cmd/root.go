package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/idokendo/aoc/cmd/year2021"
	"github.com/idokendo/aoc/cmd/year2023"
	"github.com/idokendo/aoc/cmd/year2024"
	"github.com/idokendo/aoc/cmd/year2025"
	"github.com/idokendo/aoc/templates"
	"github.com/spf13/cobra"
)

var templateFS = templates.FS

type TemplateData struct {
	Year string
	Day  string
}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "aoc",
	Long:  "aoc",
}

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap [year] [day]",
	Short: "Create scaffolding for a new day",
	Long:  "Create scaffolding for a new day",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		year := args[0]
		day := args[1]

		dir := fmt.Sprintf("cmd/year%s/day%s", year, day)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cmdFile := filepath.Join(dir, "cmd.go")
		tmpl, err := template.ParseFS(templateFS, "day_cmd.go.tmpl")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, TemplateData{Year: year, Day: day})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		cmdContent := buf.String()
		if err := os.WriteFile(cmdFile, []byte(cmdContent), 0644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		testFile := filepath.Join(dir, "cmd_test.go")
		tmpl, err = template.ParseFS(templateFS, "day_cmd_test.go.tmpl")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		buf.Reset()
		err = tmpl.Execute(&buf, TemplateData{Year: year, Day: day})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		testContent := buf.String()
		if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		emptyFiles := []string{"test1.txt", "test2.txt", "input1.txt", "input2.txt"}
		for _, file := range emptyFiles {
			test1File := filepath.Join(dir, file)
			if err := os.WriteFile(test1File, []byte(""), 0644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		yearDir := fmt.Sprintf("cmd/year%s", year)
		yearCmdFile := filepath.Join(yearDir, "cmd.go")
		if _, err := os.Stat(yearCmdFile); os.IsNotExist(err) {
			tmpl, err = template.ParseFS(templateFS, "year_cmd.go.tmpl")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			buf.Reset()
			err = tmpl.Execute(&buf, TemplateData{Year: year, Day: day})
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			yearContent := buf.String()
			if err := os.WriteFile(yearCmdFile, []byte(yearContent), 0644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			rootContent, err := os.ReadFile("cmd/root.go")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			importLine := fmt.Sprintf("\t\"github.com/idokendo/aoc/cmd/year%s\"", year)
			lines := strings.Split(string(rootContent), "\n")
			importStart := -1
			importEnd := -1
			for i, line := range lines {
				if strings.Contains(line, "import (") {
					importStart = i
				}
				if importStart != -1 && strings.Contains(line, ")") && i > importStart {
					importEnd = i
					break
				}
			}
			if importEnd != -1 {
				lines = append(lines[:importEnd], append([]string{importLine}, lines[importEnd:]...)...)
			}

			initStart := -1
			for i, line := range lines {
				if strings.Contains(line, "func init() {") {
					initStart = i
				}
				if initStart != -1 && strings.Contains(line, "}") && i > initStart && !strings.Contains(line, "func") {
					lines = append(lines[:i], append([]string{fmt.Sprintf("\trootCmd.AddCommand(year%s.Cmd)", year)}, lines[i:]...)...)
					break
				}
			}
			newRootContent := strings.Join(lines, "\n")
			if err := os.WriteFile("cmd/root.go", []byte(newRootContent), 0644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {

			yearContent, err := os.ReadFile(yearCmdFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			lines := strings.Split(string(yearContent), "\n")

			importLine := fmt.Sprintf("\t\"github.com/idokendo/aoc/cmd/year%s/day%s\"", year, day)
			importStart := -1
			importEnd := -1
			for i, line := range lines {
				if strings.Contains(line, "import (") {
					importStart = i
				}
				if importStart != -1 && strings.Contains(line, ")") && i > importStart {
					importEnd = i
					break
				}
			}
			if importEnd != -1 {
				lines = append(lines[:importEnd], append([]string{importLine}, lines[importEnd:]...)...)
			}

			initStart := -1
			for i, line := range lines {
				if strings.Contains(line, "func init() {") {
					initStart = i
				}
				if initStart != -1 && strings.Contains(line, "}") && i > initStart && !strings.Contains(line, "func") {
					lines = append(lines[:i], append([]string{fmt.Sprintf("\tCmd.AddCommand(day%s.Cmd)", day)}, lines[i:]...)...)
					break
				}
			}
			newYearContent := strings.Join(lines, "\n")
			if err := os.WriteFile(yearCmdFile, []byte(newYearContent), 0644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
		fmt.Println("Scaffolding created for year", year, "day", day)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
	rootCmd.AddCommand(year2021.Cmd)
	rootCmd.AddCommand(year2023.Cmd)
	rootCmd.AddCommand(year2024.Cmd)
	rootCmd.AddCommand(year2025.Cmd)
}
