package day5

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  "day5",
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

func parseNumbers(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		numbers[i], _ = strconv.Atoi(part)
	}
	return numbers
}

type Mapping struct {
	start int
	end   int
	level int
}

func pop(slice *[]Mapping) Mapping {
	last := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return last
}

func Part1(input string) int {
	var translated []int
	var origin []int
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "seeds:") {
			origin = parseNumbers(strings.TrimPrefix(line, "seeds:"))
			translated = make([]int, len(origin))
		} else if strings.HasSuffix(line, "map:") {
			for i, item := range origin {
				if item != -1 {
					translated[i] = item
				}
			}
			copy(origin, translated)
			translated = make([]int, len(origin))
		} else if len(line) > 0 {
			numbers := parseNumbers(line)
			for i, item := range origin {
				if item == -1 {
					continue
				}
				if item >= numbers[1] && item < numbers[1]+numbers[2] {
					res := numbers[0] - numbers[1] + item
					origin[i] = -1
					translated[i] = res
				}
			}
		}
	}
	for i, item := range origin {
		if item != -1 {
			translated[i] = item
		}
	}
	minLocation := math.MaxInt
	for _, location := range translated {
		minLocation = min(minLocation, location)
	}
	return minLocation
}

func Part2(input string) int {
	categories := strings.Split(input, "\n\n")
	var mappings []Mapping

	seeds := parseNumbers(strings.TrimPrefix(categories[0], "seeds:"))
	for i := 0; i < len(seeds); i += 2 {
		mappings = append(mappings, Mapping{seeds[i], seeds[i] + seeds[i+1], 1})
	}

	minLocation := math.MaxInt

	for len(mappings) > 0 {
		mapping := pop(&mappings)

		if mapping.level == 8 {
			minLocation = min(mapping.start, minLocation)
			continue
		}

		matches := regexp.MustCompile(`(\d+) (\d+) (\d+)`).
			FindAllStringSubmatch(categories[mapping.level], -1)
		newStart := mapping.start
		newEnd := mapping.end
		for _, match := range matches {
			dstStart, _ := strconv.Atoi(match[1])
			srcStart, _ := strconv.Atoi(match[2])
			length, _ := strconv.Atoi(match[3])
			srcEnd := srcStart + length
			diff := dstStart - srcStart

			if mapping.end <= srcStart || srcEnd <= mapping.start {
				continue
			}

			if mapping.start < srcStart {
				mappings = append(mappings, Mapping{mapping.start, srcStart, mapping.level})
				newStart = srcStart
			}

			if srcEnd < mapping.end {
				mappings = append(mappings, Mapping{srcEnd, mapping.end, mapping.level})
				newEnd = srcEnd
			}

			newStart += diff
			newEnd += diff
			break
		}

		mappings = append(mappings, Mapping{newStart, newEnd, mapping.level + 1})
	}

	return minLocation
}
