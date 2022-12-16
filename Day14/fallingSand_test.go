package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cave_buildCave(t *testing.T) {
	assert := assert.New(t)

	input := "527,102 -> 527,106 -> 523,106 -> 523,111 -> 540,111 -> 540,106 -> 533,106 -> 533,102"

	coords := getCoordinates(input)
	assert.Equal("527,102", coords[0].String())
	assert.Equal("527,106", coords[1].String())
	assert.Equal("523,106", coords[2].String())
	assert.Equal("523,111", coords[3].String())
	assert.Equal("540,111", coords[4].String())
	assert.Equal("540,106", coords[5].String())
	assert.Equal("533,106", coords[6].String())
	assert.Equal("533,102", coords[7].String())

	cave := cave{display: true, layout: map[int]map[int]string{}, min: coordinates{999999, 0}}
	cave.buildCave(input)
	cave.render(true)
}
