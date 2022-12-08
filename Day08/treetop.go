package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"strconv"
	"strings"
)

type grid [][]int

func main() {
	treeGrid := getTreeGrid()

	visibleTrees := 0
	for i, _ := range treeGrid.grid {
		for j, treeHeight := range treeGrid.grid[i] {
			if treeGrid.visible(treeHeight, i, j) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("Part 1: There are %d visible trees\n", visibleTrees)
	fmt.Printf("Part 2: Best possible scenic score is %d\n", treeGrid.scenicScoreMax())
}

type treeGrid struct {
	grid    grid
	maxUp   []int
	maxLeft []int
}

func (tg *treeGrid) visible(height int, row int, col int) (visible bool) {
	if tg.visibleDir(height, row, col, 0, 1) { // Down
		visible = true
	} else if tg.visibleDir(height, row, col, 0, -1) { // Up
		visible = true
	} else if tg.visibleDir(height, row, col, 1, 0) { // Right
		visible = true
	} else if tg.visibleDir(height, row, col, -1, 0) { // Left
		visible = true
	}

	if tg.maxLeft[col] < tg.grid[row][col] {
		tg.maxLeft[col] = tg.grid[row][col]
	}
	if tg.maxUp[row] < tg.grid[row][col] {
		tg.maxUp[row] = tg.grid[row][col]
	}

	return
}

func (tg *treeGrid) visibleDir(height int, row int, col int, hor int, vert int) bool {
	if row == 0 || col == 0 || row == len(tg.grid)-1 || col == len(tg.grid[0])-1 {
		return true
	} else if tg.grid[row+vert][col+hor] >= height ||
		(hor == -1 && tg.maxUp[row] >= height) ||
		(vert == -1 && tg.maxLeft[col] >= height) {

		return false
	} else {
		return tg.visibleDir(height, row+vert, col+hor, hor, vert)
	}
}

func (tg *treeGrid) scenicScoreMax() (max int) {
	for i, _ := range tg.grid {
		for j, height := range tg.grid[i] {
			scenicScore := tg.scenicScore(height, i, j, 0, 1, 0)  // Down
			scenicScore *= tg.scenicScore(height, i, j, 0, -1, 0) // UP
			scenicScore *= tg.scenicScore(height, i, j, 1, 0, 0)  // Right
			scenicScore *= tg.scenicScore(height, i, j, -1, 0, 0) // Left
			if max < scenicScore {
				max = scenicScore
			}
		}
	}

	return
}

func (tg *treeGrid) scenicScore(height int, row int, col int, hor int, vert int, score int) int {
	if (row == 0 && vert == -1) ||
		(row == len(tg.grid)-1 && vert == 1) ||
		(col == 0 && hor == -1) ||
		(col == len(tg.grid[0])-1 && hor == 1) {
		return score
	} else if tg.grid[row+vert][col+hor] >= height {
		return score + 1
	} else {
		return tg.scenicScore(height, row+vert, col+hor, hor, vert, score+1)
	}
}

func getTreeGrid() treeGrid {
	var treeGrid treeGrid
	treeGrid.grid = make([][]int, 0)

	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	for fs.Scan() {
		var treeRow = make([]int, 0)
		for _, v := range strings.Split(fs.Text(), "") {
			i, err := strconv.Atoi(v)
			aoc.ReportError(err)
			treeRow = append(treeRow, i)
		}

		treeGrid.grid = append(treeGrid.grid, treeRow)
	}
	treeGrid.maxUp = make([]int, len(treeGrid.grid))
	treeGrid.maxLeft = make([]int, len(treeGrid.grid[0]))

	return treeGrid
}
