package day5_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2023/day5"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  35,
			input: "test1.txt",
			fn:    day5.Part1,
		},
		{
			want:  46,
			input: "test2.txt",
			fn:    day5.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
