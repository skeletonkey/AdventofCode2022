package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"sort"
)

var stepsTaken int64

func main() {
	hillMap := newChart()
	fmt.Printf("\nPart 1: Shortest Path requires %d steps\n", hillMap.findFewestSteps())
}

type position struct {
	col int
	row int
}

type chart struct {
	end      position
	grid     [][]rune
	start    position
	startRow int
	steps    []int
}

func newChart() (c chart) {
	c.end = position{-1, -1}
	c.start = position{-1, -1}
	c.steps = make([]int, 0)
	c.load()
	return
}

func (c *chart) load() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	c.grid = make([][]rune, 0)
	rowCount := 0
	for fs.Scan() {
		c.grid = append(c.grid, []rune(fs.Text()))
		if c.start.col == -1 || c.end.col == -1 {
			for i, v := range c.grid[rowCount] {
				if v == 'S' {
					c.start = position{i, rowCount}
					c.grid[c.start.row][c.start.col] = 'a'
				} else if v == 'E' {
					c.end = position{i, rowCount}
					c.grid[c.end.row][c.end.col] = 'z'
				}
			}
		}
		rowCount++
	}
}

func (c *chart) findFewestSteps() int {
	c.findPaths(c.start, newPath())
	sort.Ints(c.steps)
	if len(c.steps) > 1 {
		return c.steps[0]
	} else {
		return 0
	}

}

func (c *chart) findPaths(cur position, path path) {
	if c.end == cur {
		c.steps = append(c.steps, path.length)
		path.remove(cur)
		return
	}
	path.add(cur)
	stepsTaken++
	fmt.Printf("Steps Taken: %d\r", stepsTaken)

	// Up
	if cur.row != 0 && c.moveAllowed(cur, position{cur.col, cur.row - 1}, path) {
		c.findPaths(position{cur.col, cur.row - 1}, path)
	}
	// Down
	if cur.row != len(c.grid)-1 && c.moveAllowed(cur, position{cur.col, cur.row + 1}, path) {
		c.findPaths(position{cur.col, cur.row + 1}, path)
	}
	// Left
	if cur.col != 0 && c.moveAllowed(cur, position{cur.col - 1, cur.row}, path) {
		c.findPaths(position{cur.col - 1, cur.row}, path)
	}
	// Right
	if cur.col != len(c.grid[cur.row])-1 && c.moveAllowed(cur, position{cur.col + 1, cur.row}, path) {
		c.findPaths(position{cur.col + 1, cur.row}, path)
	}

	path.remove(cur)
}
func (c *chart) moveAllowed(current position, next position, path path) bool {
	if path.exists(next) {
		return false
	}
	return c.grid[next.row][next.col]-c.grid[current.row][current.col] == 0 ||
		c.grid[next.row][next.col]-c.grid[current.row][current.col] == 1
}

type path struct {
	locs   map[int]map[int]struct{}
	length int
}

func newPath() (p path) {
	p.locs = make(map[int]map[int]struct{}, 0)
	return
}
func (p *path) add(pos position) {
	if p.exists(pos) {
		aoc.ReportError(fmt.Errorf("can not add an existing position: %v", pos))
	} else {
		if _, found := p.locs[pos.row]; !found {
			p.locs[pos.row] = make(map[int]struct{})
		}
		p.locs[pos.row][pos.col] = struct{}{}
		p.length++
	}
}
func (p *path) remove(pos position) {
	if p.exists(pos) {
		delete(p.locs[pos.row], pos.col)
		p.length--
	}
}
func (p *path) exists(pos position) bool {
	if _, found := p.locs[pos.row][pos.col]; found {
		return true
	} else {
		return false
	}
}
