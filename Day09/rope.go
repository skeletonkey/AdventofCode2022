package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	//rope := newRope(2) // part 1
	rope := newRope(10) // part 2
	for fs.Scan() {
		parts := strings.Split(fs.Text(), " ")
		moveCount, err := strconv.Atoi(parts[1])
		aoc.ReportError(err)
		for i := 0; i < moveCount; i++ {
			rope.moveHead(parts[0])
		}
	}

	fmt.Printf("Tail visited %d positions at least once\n", len(rope.tailVisited))
}

type position struct {
	x int
	y int
}
type knot struct {
	pos position
}
type rope struct {
	knots       []knot
	tailIndex   int
	tailVisited map[position]struct{}
}

func newRope(numOfKnots int) rope {
	startPos := position{0, 0}

	var tailVisited = make(map[position]struct{})
	tailVisited[startPos] = struct{}{}

	knots := make([]knot, numOfKnots, numOfKnots)
	for i := 0; i < numOfKnots; i++ {
		knots[i] = knot{startPos}
	}

	return rope{knots, numOfKnots - 1, tailVisited}
}
func (r *rope) moveHead(dir string) {
	switch dir {
	case "R":
		r.knots[0].pos.x++
	case "L":
		r.knots[0].pos.x--
	case "U":
		r.knots[0].pos.y++
	case "D":
		r.knots[0].pos.y--
	}
	r.moveKnots()
}
func (r *rope) moveKnots() {
	for i := 1; i <= r.tailIndex; i++ {
		o := i - 1
		delta := position{int(math.Abs(float64(r.knots[o].pos.x - r.knots[i].pos.x))),
			int(math.Abs(float64(r.knots[o].pos.y - r.knots[i].pos.y)))}
		// Rules:
		// 0, 0 - do nothing
		// 1, 0 - do nothing
		// 0, 1 - do nothing
		// 1, 1 - do nothing
		// 2, 0 - move 1 in dir of x
		// 0, 2 - move 1 in dir of y
		// 1, 2 - move 1 in dir of x, 1 in dir of y
		// 2, 2 - move 1 in dir of x, 1 in dir of y
		// 3, 3 (or greater) - rope broke!!!!
		if delta.x >= 3 && delta.y >= 3 {
			aoc.ReportError(fmt.Errorf("knots[%d] and knots[%d] are too far apart: %v", o, i, r))
		} else if delta.x == 2 && delta.y == 0 { // horizontally
			r.knots[i].pos.x += (r.knots[o].pos.x - r.knots[i].pos.x) / delta.x
		} else if delta.x == 0 && delta.y == 2 { // vertically
			r.knots[i].pos.y += (r.knots[o].pos.y - r.knots[i].pos.y) / delta.y
		} else if delta.x >= 1 && delta.y >= 1 && !(delta.x == 1 && delta.y == 1) { // diagonally
			r.knots[i].pos.x += (r.knots[o].pos.x - r.knots[i].pos.x) / delta.x
			r.knots[i].pos.y += (r.knots[o].pos.y - r.knots[i].pos.y) / delta.y
		} else {
			break // once a knot doesn't move subsequent knots will not move either
		}
	}
	r.tailVisited[r.knots[r.tailIndex].pos] = struct{}{}
}
