package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const colWidth = 4

type board struct {
	data [][]string
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

	instructionRE := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	var part1 board
	var part2 board
	buildingBoard := true
	fs := bufio.NewScanner(data)
	for fs.Scan() {
		if fs.Text() == "" {
			buildingBoard = false
			continue
		}
		if buildingBoard {
			part1.add(fs.Text())
			part2.add(fs.Text())
		} else {
			parts := instructionRE.FindStringSubmatch(fs.Text())
			count, err := strconv.Atoi(parts[1])
			reportError(err)
			from, err := strconv.Atoi(parts[2])
			reportError(err)
			to, err := strconv.Atoi(parts[3])
			reportError(err)
			part1.move9000(count, from, to)
			part2.move9001(count, from, to)
		}
	}
	fmt.Printf("Part 1 crates: %s\n", part1.topRow())
	fmt.Printf("Part 2 crates: %s\n", part2.topRow())
}

func reportError(err error) {
	if err != nil {
		panic(err)
	}
}

func newBoard() *board {
	b := board{make([][]string, 0)}
	return &b
}

func (b *board) add(line string) {
	chars := strings.Split(line, "")

	dataIndex := 0
	for i := 0; i < len(chars); i += colWidth {
		if len(b.data) <= dataIndex {
			b.data = append(b.data, []string{})
		}
		if chars[i] == "[" {
			b.data[dataIndex] = append(b.data[dataIndex], chars[i+1])
		}
		dataIndex++
	}
}

func (b *board) move9000(count int, from int, to int) {
	from--
	to--
	for i := 0; i < count; i++ {
		b.data[to] = append(b.data[to], "")
		copy(b.data[to][1:], b.data[to][0:])
		b.data[to][0] = b.data[from][0]
		b.data[from] = b.data[from][1:]
	}
}

func (b *board) move9001(count int, from int, to int) {
	from--
	to--
	for i := 0; i < count; i++ {
		b.data[to] = append(b.data[to], "")
	}
	copy(b.data[to][count:], b.data[to][0:])
	copy(b.data[to][0:count], b.data[from][0:count])
	b.data[from] = b.data[from][count:]
}

func (b *board) topRow() string {
	var buf bytes.Buffer
	for i := 0; i < len(b.data); i++ {
		if len(b.data[i]) > 0 {
			buf.WriteString(b.data[i][0])
		} else {
			buf.WriteString(" ")
		}
	}

	return buf.String()
}

func (b *board) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(b.data); i++ {
		for j := 0; j < len(b.data[i]); j++ {
			if b.data[i][j] == "" {
				buf.WriteString("-")
			} else {
				buf.WriteString(b.data[i][j])
			}
		}
	}
	return buf.String()
}
