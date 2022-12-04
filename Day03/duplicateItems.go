package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type badgeType map[string]int

func main() {
	data, err := os.Open("data")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	fs := bufio.NewScanner(data)

	dupsTotal := 0
	badgesTotal := 0
	badges := make(badgeType)
	lineCount := 1
	for fs.Scan() {
		parts := strings.Split(fs.Text(), "")
		var chars = make(map[string]int)
		for i := 0; i < len(parts)/2; i++ {
			chars[parts[i]] = 0
			badges[parts[i]] = newBadgeValue(badges, parts[i], lineCount)
		}
		for i := len(parts) / 2; i < len(parts); i++ {
			badges[parts[i]] = newBadgeValue(badges, parts[i], lineCount)
			_, ok := chars[parts[i]]
			if ok {
				chars[parts[i]] = 1
			}
		}

		for char, val := range chars {
			if val > 0 {
				dupsTotal += getValue(char)
			}
		}

		if lineCount == 3 {
			for char, val := range badges {
				if val == 3 {
					badgesTotal += getValue(char)
				}
			}
			badges = make(badgeType)
			lineCount = 0

		}
		lineCount++
	}
	fmt.Printf("Part 1 Duplicate Score: %d\n", dupsTotal)
	fmt.Printf("Part 2 Badges Score: %d\n", badgesTotal)
}

func getValue(char string) int {
	r := []rune(char)
	if int(r[0]) > int('a') {
		return int(r[0]) - int('a') + 1
	} else {
		return int(r[0]) - int('A') + 27
	}
}

func newBadgeValue(badge badgeType, part string, line int) int {
	if line == 1 {
		return 1
	} else {
		if badge[part] >= line-1 {
			return line
		}
	}
	return 0
}
