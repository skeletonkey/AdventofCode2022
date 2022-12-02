package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type max struct {
	cals     int
	whichElf int
}

func main() {
	x := 3
	topX := make([]max, x)

	datFile, err := os.Open("elf_cals.dat")
	defer datFile.Close()

	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(datFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	elfCounter := 1
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			cals, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
			sum += cals
		} else {
			for i := 0; i < x; i++ {
				if sum > topX[i].cals {
					if i != x-1 { // it's not the last element os we need ot move things down
						for j := x - 1; j > i; j-- {
							topX[j] = topX[j-1]
						}
					}
					topX[i].cals = sum
					topX[i].whichElf = elfCounter
					break
				}
			}
			sum = 0
		}
		elfCounter++
	}

	sum = 0
	var whichElvesBuffer strings.Builder
	for _, top := range topX {
		sum += top.cals
		if whichElvesBuffer.Len() > 0 {
			whichElvesBuffer.WriteString(", ")
		}
		whichElvesBuffer.WriteString(strconv.Itoa(top.whichElf))
	}

	fmt.Printf("Sum of top %d is %d - which is elves number %s\n", x, sum, whichElvesBuffer.String())
}
