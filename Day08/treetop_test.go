package main

import (
	"fmt"
	"testing"
)

func Test_treeGrid_visibleDir(t *testing.T) {
	type fields struct {
		grid    grid
		maxUp   []int
		maxLeft []int
	}
	type args struct {
		height int
		row    int
		col    int
		hor    int
		vert   int
	}

	testGrid := grid{[]int{1, 4, 3, 4}, []int{4, 2, 3, 1}, []int{1, 2, 3, 4}, []int{1, 4, 3, 1}}
	testFields := fields{testGrid, make([]int, 4), make([]int, 4)}

	tests := []struct {
		args args
		want bool
	}{
		{
			args{testGrid[1][1], 1, 1, 1, 0},
			false,
		},
		{
			args{testGrid[1][1], 1, 1, -1, 0},
			false,
		},
		{
			args{testGrid[1][1], 1, 1, 0, 1},
			false,
		},
		{
			args{testGrid[1][1], 1, 1, 0, -1},
			false,
		},
		{
			args{testGrid[1][2], 1, 2, 1, 0},
			true,
		},
		{
			args{testGrid[1][2], 1, 2, -1, 0},
			false,
		},
		{
			args{testGrid[1][2], 1, 2, 0, 1},
			false,
		},
		{
			args{testGrid[1][2], 1, 2, 0, -1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Cell: %d, %d - Moving: %d, %d", tt.args.row, tt.args.col, tt.args.hor, tt.args.vert), func(t *testing.T) {
			tg := &treeGrid{
				grid:    testFields.grid,
				maxUp:   testFields.maxUp,
				maxLeft: testFields.maxLeft,
			}
			if got := tg.visibleDir(tt.args.height, tt.args.row, tt.args.col, tt.args.hor, tt.args.vert); got != tt.want {
				t.Errorf("visibleDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
