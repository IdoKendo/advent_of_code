package day17_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day17"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  string
		input string
		fn    func(string) string
	}{
		{
			want:  "4,6,3,5,6,3,5,2,1,0",
			input: "test1.txt",
			fn:    day17.Part1,
		},
		{
			want:  "117440",
			input: "test2.txt",
			fn:    day17.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
