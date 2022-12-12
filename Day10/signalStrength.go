package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"strconv"
	"strings"
)

func main() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	screen := newScreen(40, 6)

	interrupts := map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}
	register := 1
	cycle := 0
	for fs.Scan() {
		parts := strings.Split(fs.Text(), " ")

		if parts[0] == "noop" {
			screen.drawPixel(cycle, register)
			cycle++
			if _, ok := interrupts[cycle]; ok {
				interrupts[cycle] = register
			}
		} else {
			screen.drawPixel(cycle, register)
			cycle++
			if _, ok := interrupts[cycle]; ok {
				interrupts[cycle] = register
			}
			screen.drawPixel(cycle, register)
			cycle++
			if _, ok := interrupts[cycle]; ok {
				interrupts[cycle] = register
			}
			shift, err := strconv.Atoi(parts[1])
			aoc.ReportError(err)
			register += shift
		}
	}

	total := 0
	for x, y := range interrupts {
		total += x * y
	}

	fmt.Printf("Part 1 total: %d\n", total)

	fmt.Println("Screen display:")
	screen.render()
}

type screen struct {
	width     int
	height    int
	display   [][]string
	litPixel  string
	darkPixel string
}

func newScreen(width int, height int) screen {
	display := make([][]string, height, height)
	for i, _ := range display {
		display[i] = make([]string, width, width)
	}
	return screen{width: width, height: height, display: display, litPixel: "#", darkPixel: "."}
}
func (s *screen) drawPixel(cycle int, register int) {
	col := cycle
	row := 0
	for col >= s.width {
		col -= s.width
		row++
	}
	if col >= register-1 && col <= register+1 {
		s.display[row][col] = s.litPixel
	} else {
		s.display[row][col] = s.darkPixel
	}
}
func (s *screen) render() {
	for _, row := range s.display {
		fmt.Println(strings.Join(row, ""))
	}
}
