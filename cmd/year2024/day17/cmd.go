package day17

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day17",
	Short: "day17",
	Long:  "day17",
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

type instruction struct {
	opcode  int
	operand int
}

type program struct {
	a, b, c      int
	instructions []instruction
	repr         []int
	p            string
}

func newProgram(input string) program {
	re := regexp.MustCompile(`Register A: (\d+)\s+Register B: (\d+)\s+Register C: (\d+)\s+Program: ([\d,]+)`)
	matches := re.FindStringSubmatch(input)
	a, _ := strconv.Atoi(matches[1])
	b, _ := strconv.Atoi(matches[2])
	c, _ := strconv.Atoi(matches[3])

	instructionStrings := strings.Split(matches[4], ",")
	instructions := make([]instruction, len(instructionStrings)/2)
	repr := make([]int, len(instructionStrings))
	for i := 0; i < len(instructionStrings); i += 2 {
		opcode, _ := strconv.Atoi(instructionStrings[i])
		operand, _ := strconv.Atoi(instructionStrings[i+1])
		instructions[i/2] = instruction{opcode, operand}
		repr[i] = opcode
		repr[i+1] = operand
	}

	return program{
		a:            a,
		b:            b,
		c:            c,
		instructions: instructions,
		repr:         repr,
		p:            matches[4],
	}
}

func (p program) execute() []int {
	outputs := []int{}

	for ip := 0; ip < len(p.instructions); ip++ {
		opcode := p.instructions[ip].opcode
		literal := p.instructions[ip].operand
		operand := literal
		switch operand {
		case 4:
			operand = p.a
		case 5:
			operand = p.b
		case 6:
			operand = p.b
		case 7:
			panic("invalid program")
		}

		switch opcode {
		case 0: // adv - divide A by 2^value
			p.a >>= operand
		case 1: // bxl - XOR B with literal
			p.b ^= literal
		case 2: // bst - set B to value mod 8
			p.b = operand % 8
		case 3: // jnz - jump if A is not zero
			if p.a != 0 {
				ip = literal - 1
			}
		case 4: // bxc - XOR B with C
			p.b ^= p.c
		case 5: // out - output value mod 8
			outputs = append(outputs, operand%8)
		case 6: // bdv - divide A by 2^value, store in B
			p.b = p.a >> operand
		case 7: // cdv - divide A by 2^value, store in C
			p.c = p.a >> operand
		}
	}

	return outputs
}

func (p program) valueForCopy() int {
	var current int
	var outputs []int
	for digit := len(p.repr) - 1; digit >= 0; digit-- {
		for i := 0; i < math.MaxInt32; i++ {
			p.a = current + (1<<(digit*3))*i
			p.b = 0
			p.c = 0
			outputs = p.execute()
			if len(outputs) < digit || !slices.Equal(outputs[digit:], p.repr[digit:]) {
				continue
			}
			current = p.a
			break
		}
	}
	return current
}

func Part1(input string) string {
	p := newProgram(input)
	outputs := p.execute()
	res := make([]string, len(outputs))
	for i, v := range outputs {
		res[i] = strconv.Itoa(v)
	}

	return strings.Join(res, ",")
}

func Part2(input string) string {
	p := newProgram(input)
	a := p.valueForCopy()
	return strconv.Itoa(a)
}
