package day23_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day23"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  string
		input string
		fn    func(string) string
	}{
		{
			want:  "7",
			input: "test1.txt",
			fn:    day23.Part1,
		},
		{
			want:  "co,de,ka,ta",
			input: "test2.txt",
			fn:    day23.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
