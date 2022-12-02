package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	datFile = "tournament.dat"

	// Points
	rock     = 1
	paper    = 2
	scissors = 3
	lose     = 0
	draw     = 3
	win      = 6
)

func main() {
	data, err := os.Open(datFile)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	fs := bufio.NewScanner(data)

	elfThrow := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
	}
	myThrow := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	outcome := map[string]int{
		"X": lose,
		"Y": draw,
		"Z": win,
	}

	partOneTotalScore := 0
	partTwoTotalScore := 0
	for fs.Scan() {
		parts := strings.Split(fs.Text(), " ")
		partOneTotalScore += myThrow[parts[1]] + matchResult(elfThrow[parts[0]], myThrow[parts[1]])
		partTwoTotalScore += whatToThrow(elfThrow[parts[0]], outcome[parts[1]]) + outcome[parts[1]]
	}

	fmt.Printf("Part 1 Tournament Score: %d\n", partOneTotalScore)
	fmt.Printf("Part 2 Tournament Score: %d\n", partTwoTotalScore)
}

func matchResult(elfThrow int, myThrow int) int {
	if elfThrow == myThrow {
		return draw
	} else if elfThrow == rock {
		if myThrow == paper {
			return win
		} else {
			return lose
		}
	} else if elfThrow == paper {
		if myThrow == scissors {
			return win
		} else {
			return lose
		}
	} else { // scissors
		if myThrow == rock {
			return win
		} else {
			return lose
		}
	}
}

func whatToThrow(elfThrow int, outcome int) int {
	if outcome == draw {
		return elfThrow
	} else if outcome == win {
		switch elfThrow {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	} else { // lose
		switch elfThrow {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}
	log.Panicf("Couldn't figure out the proper throw given elf throw of %d and outcome of %d", elfThrow, outcome)
	return 0
}
