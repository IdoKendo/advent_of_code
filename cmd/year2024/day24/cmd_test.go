package day24_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day24"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  string
		input string
		fn    func(string) string
	}{
		{
			want:  "4",
			input: "test1.txt",
			fn:    day24.Part1,
		},
		{
			want:  "2024",
			input: "test1b.txt",
			fn:    day24.Part1,
		},
		{
			want:  "z00,z01,z02,z05",
			input: "test2.txt",
			fn:    day24.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
