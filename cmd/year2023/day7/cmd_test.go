package day7_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2023/day7"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  6440,
			input: "test1.txt",
			fn:    day7.Part1,
		},
		{
			want:  5905,
			input: "test2.txt",
			fn:    day7.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
