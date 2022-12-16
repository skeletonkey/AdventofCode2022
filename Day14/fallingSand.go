package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	empty   = " "
	entry   = "^"
	falling = "+"
	rock    = "#"
	sand    = "0"
)

func main() {
	leakLocation := coordinates{500, 0}

	cave1 := newCave()
	cave1.setLeakLocation(leakLocation)
	fmt.Printf("Part 1: %d units of sand where required\n", cave1.fillInfinite())

	cave2 := newCave()
	cave2.setLeakLocation(leakLocation)
	// set floor
	cave2.max.y += 2
	//cave2.display = true
	fmt.Printf("Part 2: %d units of sand where required\n", cave2.fill())
}

type coordinates struct {
	x int
	y int
}

func (c coordinates) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type cave struct {
	filled                 bool
	infinity               bool
	layout                 map[int]map[int]string
	min, max, leakLocation coordinates

	display  bool
	lineNums bool
}

func newCave() (c cave) {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()
	c.layout = map[int]map[int]string{}
	c.display = false
	c.lineNums = false
	c.min = coordinates{999999, 999999}

	for fs.Scan() {
		c.buildCave(fs.Text())
	}

	return
}

func (c *cave) buildCave(input string) {
	coords := getCoordinates(input)
	for i, coord := range coords {
		if i == len(coords)-1 {
			c.setLocation(coordinates{coord.x, coord.y}, rock)
		} else {
			if hor := coords[i].x - coords[i+1].x; hor != 0 {
				var start, end int
				if hor > 0 {
					start = coords[i+1].x
					end = coords[i].x
				} else {
					start = coords[i].x
					end = coords[i+1].x
				}
				y := coords[i].y
				if y != coords[i+1].y {
					aoc.ReportError(fmt.Errorf("both x and y coordinates are changing: %v - %v (%s)", coords[i], coords[i+1], input))
				}
				for x := start; x <= end; x++ {
					c.setLocation(coordinates{x, y}, rock)
				}
			} else if vert := coords[i].y - coords[i+1].y; vert != 0 {
				var start, end int
				if vert > 0 {
					start = coords[i+1].y
					end = coords[i].y
				} else {
					start = coords[i].y
					end = coords[i+1].y
				}
				x := coords[i].x
				if x != coords[i+1].x {
					aoc.ReportError(fmt.Errorf("both x and y coordinates are changing: %v - %v (%s)", coords[i], coords[i+1], input))
				}
				for y := start; y <= end; y++ {
					c.setLocation(coordinates{x, y}, rock)
				}
			} else {
				aoc.ReportError(fmt.Errorf("coordinate doesn't move: %v - %v (%s)", coords[i], coords[i+1], input))
			}
		}
	}
}

func (c *cave) fill() int {
	sandUnits := 0
	for !c.filled {
		c.dropUnitOfSand(c.leakLocation)
		sandUnits++
	}
	return sandUnits
}
func (c *cave) dropUnitOfSand(cur coordinates) {
	c.render(cur)
	if cur.x == c.leakLocation.x && cur.y == c.leakLocation.y && !c.spaceToMove(c.leakLocation) {
		c.filled = true
		c.setLocation(cur, sand)
	} else if cur.y+1 == c.max.y { // there are no more obstacles - falls forever
		c.setLocation(cur, sand)
	} else if _, obstacle := c.layout[cur.x][cur.y+1]; !obstacle { // move down
		c.dropUnitOfSand(coordinates{cur.x, cur.y + 1})
	} else if _, obstacle = c.layout[cur.x-1][cur.y+1]; !obstacle { // move left
		c.dropUnitOfSand(coordinates{cur.x - 1, cur.y + 1})
	} else if _, obstacle = c.layout[cur.x+1][cur.y+1]; !obstacle { // move right
		c.dropUnitOfSand(coordinates{cur.x + 1, cur.y + 1})
	} else {
		c.setLocation(cur, sand)
	}
}

func (c *cave) spaceToMove(cur coordinates) bool {
	_, obstacleDownLeft := c.layout[cur.x-1][cur.y+1]
	_, obstacleDown := c.layout[cur.x][cur.y+1]
	_, obstacleDownRight := c.layout[cur.x+1][cur.y+1]
	if obstacleDown && obstacleDownRight && obstacleDownLeft {
		return false
	}
	return true
}
func (c *cave) fillInfinite() int {
	sandUnits := 0
	for true {
		c.dropUnitOfSandInfinite(c.leakLocation)
		if c.infinity {
			break
		} else {
			sandUnits++
		}
	}
	return sandUnits
}
func (c *cave) dropUnitOfSandInfinite(cur coordinates) {
	c.render(cur)
	if cur.y >= c.max.y { // there are no more obstacles - falls forever
		c.infinity = true
	} else if _, obstacle := c.layout[cur.x][cur.y+1]; !obstacle { // move down
		c.dropUnitOfSandInfinite(coordinates{cur.x, cur.y + 1})
	} else if _, obstacle = c.layout[cur.x-1][cur.y+1]; !obstacle { // move left
		c.dropUnitOfSandInfinite(coordinates{cur.x - 1, cur.y + 1})
	} else if _, obstacle = c.layout[cur.x+1][cur.y+1]; !obstacle { // move right
		c.dropUnitOfSandInfinite(coordinates{cur.x + 1, cur.y + 1})
	} else {
		c.setLocation(cur, sand)
	}

	return
}

func (c *cave) setLeakLocation(cur coordinates) {
	c.leakLocation = cur
	c.setLocation(cur, entry)
}

func (c *cave) setLocation(cur coordinates, obstacle string) {
	if _, exists := c.layout[cur.x]; !exists {
		c.layout[cur.x] = map[int]string{}
	}
	c.layout[cur.x][cur.y] = obstacle
	if c.max.x < cur.x {
		c.max.x = cur.x
	}
	if c.max.y < cur.y {
		c.max.y = cur.y
	}
	if c.min.x > cur.x {
		c.min.x = cur.x
	}
	if c.min.y > cur.y {
		c.min.y = cur.y
	}
}

func (c *cave) render(fallingLoc coordinates) {
	if !c.display {
		return
	}
	for y := 0; y <= c.max.y-c.min.y; y++ {
		if c.lineNums {
			fmt.Printf("%5d  ", y)
		} else {
			fmt.Print(" ")
		}

		for x := 0; x <= c.max.x-c.min.x; x++ {
			if obstacle, exists := c.layout[x+c.min.x][y+c.min.y]; exists {
				fmt.Print(obstacle)
			} else if x == fallingLoc.x-c.min.x && y == fallingLoc.y-c.min.y {
				fmt.Print(falling)
			} else {
				fmt.Print(empty)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	time.Sleep(100 * time.Millisecond)
}

func getCoordinates(line string) []coordinates {
	parts := strings.Split(line, " -> ")
	c := make([]coordinates, len(parts), len(parts))
	for i, part := range parts {
		coords := strings.Split(part, ",")
		x, err := strconv.Atoi(coords[0])
		aoc.ReportError(err)
		y, err := strconv.Atoi(coords[1])
		aoc.ReportError(err)
		c[i] = coordinates{x, y}
	}
	return c
}
