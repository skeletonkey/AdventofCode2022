package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start int
	end   int
}

func (s section) contains(sec section) bool {
	if s.start <= sec.start && s.end >= sec.end {
		return true
	} else {
		return false
	}
}
func (s section) partialContain(sec section) bool {
	if s.end >= sec.start && s.start <= sec.end {
		return true
	} else {
		return false
	}
}

func main() {
	var data *os.File
	var err error
	if os.Getenv("TESTCASE") == "1" {
		data, err = os.Open("data_test")
	} else {
		data, err = os.Open("data")
	}
	if err != nil {
		panic(err)
	}
	defer data.Close()

	fs := bufio.NewScanner(data)
	fullContainment := 0
	partialContainment := 0
	for fs.Scan() {
		ranges := strings.Split(fs.Text(), ",")
		sections := make([]section, len(ranges))
		for i, r := range ranges {
			parts := strings.Split(r, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			sections[i] = section{start, end}
		}
		if sections[0].contains(sections[1]) || sections[1].contains(sections[0]) {
			fullContainment++
		}
		if sections[0].partialContain(sections[1]) || sections[1].partialContain(sections[0]) {
			partialContainment++
		}
	}
	fmt.Printf("Part 1 Number of ranges fully contained: %d\n", fullContainment)
	fmt.Printf("Part 2 Number of ranges partially contained: %d\n", partialContainment)
}
