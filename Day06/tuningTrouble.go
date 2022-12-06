package main

import (
	"fmt"
	"os"
	"strings"

	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
)

const packetLength = 4
const messageLength = 14

func main() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	fs.Scan()
	fmt.Printf("Characters before start of packet: %d\n", start(packetLength, fs.Text()))
	fmt.Printf("Characters before start of message: %d\n", start(messageLength, fs.Text()))
}

func start(length int, signal string) (position int) {
	chars := strings.Split(signal, "")
	for i := 0; i+length < len(chars); i++ {
		if !dupsFound(chars[i : i+length]) {
			position = i + length
			break
		}
	}

	return
}

func dupsFound(chars []string) (dups bool) {
	var seen = make(map[string]int)
	for _, char := range chars {
		seen[char]++
		if seen[char] > 1 {
			return true
		}
	}

	return
}
