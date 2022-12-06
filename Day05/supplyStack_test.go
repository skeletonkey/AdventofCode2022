package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_board_add_move(t *testing.T) {
	b := newBoard()
	part2 := newBoard()

	tests := []struct {
		name   string
		line   string
		result board
	}{
		{
			"No data",
			"",
			board{[][]string{}},
		},
		{
			"First Row",
			"    [D]",
			board{[][]string{[]string{}, []string{"D"}}},
		},
		{
			"Second row",
			"[A]",
			board{[][]string{[]string{"A"}, []string{"D"}}},
		}, {
			"Third Row",
			"[Z] [M] [P]",
			board{[][]string{[]string{"A", "Z"}, []string{"D", "M"}, []string{"P"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b.add(tt.line)
			part2.add(tt.line)
			assert.Equal(t, tt.result.String(), b.String(), tt.name)
		})
	}

	moveTests := []struct {
		name   string
		count  int
		from   int
		to     int
		result board
		topRow string
	}{
		{
			"First move",
			1, 2, 1,
			board{[][]string{[]string{"D", "A", "Z"}, []string{"M"}, []string{"P"}}},
			"DMP",
		},
		{
			"Second move",
			1, 1, 3,
			board{[][]string{[]string{"A", "Z"}, []string{"M"}, []string{"D", "P"}}},
			"AMD",
		},
		{
			"Third move",
			1, 2, 1,
			board{[][]string{[]string{"M", "A", "Z"}, []string{}, []string{"D", "P"}}},
			"M D",
		},
		{
			"Fourth move",
			2, 1, 2,
			board{[][]string{[]string{"Z"}, []string{"A", "M"}, []string{"D", "P"}}},
			"ZAD",
		},
	}
	for _, tt := range moveTests {
		t.Run(tt.name, func(t *testing.T) {
			b.move9000(tt.count, tt.from, tt.to)
			assert.Equal(t, tt.result.String(), b.String(), tt.name)
			assert.Equal(t, tt.topRow, b.topRow())
		})
	}

	movePart2Tests := []struct {
		name   string
		count  int
		from   int
		to     int
		result board
		topRow string
	}{
		{
			"First move",
			1, 2, 1,
			board{[][]string{[]string{"D", "A", "Z"}, []string{"M"}, []string{"P"}}},
			"DMP",
		},
		{
			"Second move",
			1, 1, 3,
			board{[][]string{[]string{"A", "Z"}, []string{"M"}, []string{"D", "P"}}},
			"AMD",
		},
		{
			"Third move",
			1, 2, 1,
			board{[][]string{[]string{"M", "A", "Z"}, []string{}, []string{"D", "P"}}},
			"M D",
		},
		{
			"Fourth move",
			2, 1, 2,
			board{[][]string{[]string{"Z"}, []string{"M", "A"}, []string{"D", "P"}}},
			"ZMD",
		},
		{
			"Fifth move",
			2, 2, 1,
			board{[][]string{[]string{"M", "A", "Z"}, []string{}, []string{"D", "P"}}},
			"M D",
		},
		{
			"Sixth move",
			2, 3, 1,
			board{[][]string{[]string{"D", "P", "M", "A", "Z"}, []string{}, []string{}}},
			"D  ",
		},
		{
			"Seventh move",
			4, 1, 2,
			board{[][]string{[]string{"Z"}, []string{"D", "P", "M", "A"}, []string{}}},
			"ZD ",
		},
	}
	for _, tt := range movePart2Tests {
		t.Run(tt.name, func(t *testing.T) {
			part2.move9001(tt.count, tt.from, tt.to)
			assert.Equal(t, tt.result.String(), part2.String(), tt.name)
			assert.Equal(t, tt.topRow, part2.topRow())
		})
	}
}
