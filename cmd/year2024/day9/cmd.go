package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "day9",
	Long:  "day9",
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

func newDiskMap(input string) []string {
	input = strings.TrimSuffix(input, "\n")
	numbers := strings.Split(input, "")
	digits := make([]int, len(numbers))
	for i, number := range numbers {
		digits[i], _ = strconv.Atoi(number)
	}
	diskMap := []string{}

	for i, j := 0, 0; i < len(digits); i, j = i+2, j+1 {
		occupied := digits[i]
		for x := 0; x < occupied; x++ {
			diskMap = append(diskMap, fmt.Sprintf("%d", j))
		}
		if i < len(digits)-1 {
			free := digits[i+1]
			for x := 0; x < free; x++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	return diskMap
}

func defrag(diskMap []string) {
	i := 0
	j := len(diskMap) - 1
	for j > i {
		if diskMap[i] != "." {
			i++
			continue
		}
		for diskMap[j] == "." {
			j--
		}
		diskMap[i] = diskMap[j]
		diskMap[j] = "."
		i++
		j--
	}
}

func checksum(diskMap []string) int {
	result := 0
	for i, block := range diskMap {
		if block == "." {
			continue
		}
		d, _ := strconv.Atoi(block)
		result += i * d
	}
	return result
}

func Part1(input string) int {
	diskMap := newDiskMap(input)
	defrag(diskMap)
	return checksum(diskMap)
}

type File struct {
	ID   string
	Size int
}

func newFileSlice(input string) []File {
	input = strings.TrimSuffix(input, "\n")
	numbers := strings.Split(input, "")
	digits := make([]int, len(numbers))
	for i, number := range numbers {
		digits[i], _ = strconv.Atoi(number)
	}
	diskMap := []File{}

	for i, j := 0, 0; i < len(digits); i, j = i+2, j+1 {
		diskMap = append(diskMap, File{
			ID:   fmt.Sprintf("%d", j),
			Size: digits[i],
		})
		if i < len(digits)-1 {
			diskMap = append(diskMap, File{
				ID:   ".",
				Size: digits[i+1],
			})
		}
	}

	return diskMap
}

func smartDefrag(diskMap []File) []File {
	i := 0
	j := len(diskMap) - 1
	for j > 0 {
		for diskMap[i].ID != "." {
			i++
		}
		for diskMap[j].ID == "." {
			j--
		}
		left := diskMap[:i]
		if i >= len(diskMap) || i+1 > j {
			j--
			i = 0
			continue
		}
		middle := diskMap[i+1 : j]
		right := diskMap[j+1:]
		if diskMap[i].Size == diskMap[j].Size {
			diskMap = append(append(append(append(append([]File{}, left...), diskMap[j]), middle...), diskMap[i]), right...)
			j--
			i = 0
		} else if diskMap[i].Size > diskMap[j].Size {
			diff := File{
				ID:   ".",
				Size: diskMap[i].Size - diskMap[j].Size,
			}
			diskMap[i].Size -= diff.Size
			diskMap = append(append(append(append(append(append([]File{}, left...), diskMap[j]), diff), middle...), diskMap[i]), right...)
			j--
			i = 0
		} else {
			i++
			continue
		}
	}

	return diskMap
}

func diskMapFromFileSlice(fileSlice []File) []string {
	diskMap := []string{}

	for _, file := range fileSlice {
		for range file.Size {
			diskMap = append(diskMap, file.ID)
		}
	}

	return diskMap
}

func Part2(input string) int {
	fileSlice := newFileSlice(input)
	fileSlice = smartDefrag(fileSlice)
	diskMap := diskMapFromFileSlice(fileSlice)
	return checksum(diskMap)
}
