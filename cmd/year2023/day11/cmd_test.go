package day11_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2023/day11"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  374,
			input: "test1.txt",
			fn:    day11.Part1,
		},
		{
			want:  1030,
			input: "test2.txt",
			fn:    day11.Example10,
		},
		{
			want:  8410,
			input: "test2.txt",
			fn:    day11.Example100,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
