package main

import (
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"strings"
)

func main() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	var left, right signal
	for fs.Scan() {
		if (left == signal{}) {

		} else if (right == signal{}) {

		} else {

		}
	}

}

type signal struct {
	data []string
}

func (s *signal) load(text string) {
	text = strings.Trim(text, "[]")
	s.data = strings.Split(text, ",")
}
