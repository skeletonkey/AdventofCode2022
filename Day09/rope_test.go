package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rope(t *testing.T) {
	assert := assert.New(t)

	rope := newRope(2)

	assert.Equal(rope.knots[0].pos, rope.knots[rope.tailIndex].pos, "Head and Tail started at same place")
	assert.Equal(1, len(rope.tailVisited), "Tail's only been in 1 place")

	rope.moveHead("U")
	assert.Equal(position{0, 1}, rope.knots[0].pos, "Head move 1 up")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail in same position")
	assert.Equal(1, len(rope.tailVisited), "Tail's only been in 1 place")

	rope.moveHead("D")
	assert.Equal(position{0, 0}, rope.knots[0].pos, "Head move 1 down")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail in same position")

	rope.moveHead("D")
	assert.Equal(position{0, -1}, rope.knots[0].pos, "Head move 1 down")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail in same position")

	rope.moveHead("D")
	assert.Equal(position{0, -2}, rope.knots[0].pos, "Head move 1 down")
	assert.Equal(position{0, -1}, rope.knots[rope.tailIndex].pos, "Tail followed 1 down")
	assert.Equal(2, len(rope.tailVisited), "Tail's been in 2 places")

	rope.moveHead("U")
	assert.Equal(position{0, -1}, rope.knots[0].pos, "Head move 1 Up")
	assert.Equal(position{0, -1}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(2, len(rope.tailVisited), "Tail's been in 2 places")

	rope.moveHead("R")
	assert.Equal(position{1, -1}, rope.knots[0].pos, "Head move Right")
	assert.Equal(position{0, -1}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(2, len(rope.tailVisited), "Tail's been in 2 places")

	rope.moveHead("U")
	assert.Equal(position{1, 0}, rope.knots[0].pos, "Head move Up")
	assert.Equal(position{0, -1}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(2, len(rope.tailVisited), "Tail's been in 2 places")

	rope.moveHead("U")
	assert.Equal(position{1, 1}, rope.knots[0].pos, "Head move Up")
	assert.Equal(position{1, 0}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(3, len(rope.tailVisited), "Tail's been in 3 places")

	rope.moveHead("L")
	assert.Equal(position{0, 1}, rope.knots[0].pos, "Head move Left")
	assert.Equal(position{1, 0}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(3, len(rope.tailVisited), "Tail's been in 3 places")

	rope.moveHead("L")
	assert.Equal(position{-1, 1}, rope.knots[0].pos, "Head move Left")
	assert.Equal(position{0, 1}, rope.knots[rope.tailIndex].pos, "Tail moved")
	assert.Equal(4, len(rope.tailVisited), "Tail's been in 4 places")

	rope.moveHead("U")
	assert.Equal(position{-1, 2}, rope.knots[0].pos, "Head move Up")
	assert.Equal(position{0, 1}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(4, len(rope.tailVisited), "Tail's been in 4 places")

	rope.moveHead("U")
	assert.Equal(position{-1, 3}, rope.knots[0].pos, "Head move Up")
	assert.Equal(position{-1, 2}, rope.knots[rope.tailIndex].pos, "Tail moved")
	assert.Equal(5, len(rope.tailVisited), "5 places")

	rope.moveHead("U")
	assert.Equal(position{-1, 4}, rope.knots[0].pos, "Head move Up")
	assert.Equal(position{-1, 3}, rope.knots[rope.tailIndex].pos, "Tail moved Up")
	assert.Equal(6, len(rope.tailVisited), "6 places")

	rope.moveHead("L")
	assert.Equal(position{-2, 4}, rope.knots[0].pos, "Head move Left")
	assert.Equal(position{-1, 3}, rope.knots[rope.tailIndex].pos, "Tail stayed the same")
	assert.Equal(6, len(rope.tailVisited), "6 places")

	rope.moveHead("L")
	assert.Equal(position{-3, 4}, rope.knots[0].pos, "Head move Left")
	assert.Equal(position{-2, 4}, rope.knots[rope.tailIndex].pos, "Tail moved left same")
	assert.Equal(7, len(rope.tailVisited), "7 places")

	rope.moveHead("L")
	assert.Equal(position{-4, 4}, rope.knots[0].pos, "Head move Left")
	assert.Equal(position{-3, 4}, rope.knots[rope.tailIndex].pos, "Tail moved left same")
	assert.Equal(8, len(rope.tailVisited), "8 places")
}

func Test_diag(t *testing.T) {
	assert := assert.New(t)

	rope := newRope(3)

	assert.Equal(rope.knots[0].pos, rope.knots[rope.tailIndex].pos, "Head and Tail started at same place")
	assert.Equal(1, len(rope.tailVisited), "Tail's only been in 1 place")

	rope.moveHead("U")
	assert.Equal(position{0, 1}, rope.knots[0].pos, "Head move 1 up")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail in same position")

	rope.moveHead("R")
	assert.Equal(position{1, 1}, rope.knots[0].pos, "Head move 1 right")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail in same position")

	rope.moveHead("U")
	assert.Equal(position{1, 2}, rope.knots[0].pos, "Head move 1 up")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail in same position")

	rope.moveHead("R")
	assert.Equal(position{2, 2}, rope.knots[0].pos, "Head move 1 right")
	assert.Equal(position{0, 0}, rope.knots[rope.tailIndex].pos, "Tail moves diagonal")

	rope.moveHead("U")
	assert.Equal(position{2, 3}, rope.knots[0].pos, "Head move 1 up")
	assert.Equal(position{1, 1}, rope.knots[rope.tailIndex].pos, "Tail in same position")

	rope.moveHead("R")
	assert.Equal(position{3, 3}, rope.knots[0].pos, "Head move 1 right")
	assert.Equal(position{1, 1}, rope.knots[rope.tailIndex].pos, "Tail moves diagonal")
}
