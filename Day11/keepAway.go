package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Part 1
// const rounds = 20
//const worryReductionFactor = 3

// Part 2
const rounds = 10000
const worryReductionFactor = 1

func main() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()
	monkeys := make([]monkey, 0)
	description := make([]string, 0)
	for fs.Scan() {
		if fs.Text() == "" {
			monkeys = append(monkeys, newMonkey(description, len(monkeys)))
			description = make([]string, 0)
		} else {
			description = append(description, fs.Text())
		}
	}
	monkeys = append(monkeys, newMonkey(description, len(monkeys)))

	for round := 0; round < rounds; round++ {
		for i := range monkeys {
			for _, item := range monkeys[i].items {
				worryLevel := monkeys[i].inspect(item)
				monkeys[monkeys[i].test(worryLevel)].addItem(worryLevel)
			}
			monkeys[i].clearItems()
		}
	}

	activeMonkeys := getMostActiveMonkeys(monkeys, 2)
	monkeyBusiness := activeMonkeys[0].inspectionCount * activeMonkeys[1].inspectionCount
	fmt.Printf("Part 1 monkey business: %d\n", monkeyBusiness)
}

type monkey struct {
	items           []int
	operator        func(int) int
	test            func(int) int
	inspectionCount int
}

func newMonkey(description []string, monkeyCount int) (monkey monkey) {
	countRE := regexp.MustCompile(`Monkey (\d+):`)
	countCG := countRE.FindStringSubmatch(description[0])
	count, err := strconv.Atoi(countCG[1])
	aoc.ReportError(err)
	if count != monkeyCount {
		aoc.ReportError(fmt.Errorf("counts (%d - %d) do not match", count, monkeyCount))
	}

	itemRE := regexp.MustCompile(`Starting items: (.+)$`)
	itemCG := itemRE.FindStringSubmatch(description[1])
	for _, item := range strings.Split(itemCG[1], ", ") {
		i, err := strconv.Atoi(item)
		aoc.ReportError(err)
		monkey.items = append(monkey.items, i)
	}

	opRE := regexp.MustCompile(`Operation: new = old (\S+) (\S+)`)
	opCG := opRE.FindStringSubmatch(description[2])
	if opCG[2] == "old" {
		monkey.operator = func(x int) int { return x * x }
	} else {
		i, err := strconv.Atoi(opCG[2])
		aoc.ReportError(err)
		switch opCG[1] {
		case "+":
			monkey.operator = func(x int) int { return x + i }
		case "-":
			monkey.operator = func(x int) int { return x - i }
		case "*":
			monkey.operator = func(x int) int { return x * i }
		default:
			aoc.ReportError(fmt.Errorf("unknown opperator: %s", opCG[1]))
		}
	}

	testRE := regexp.MustCompile(`Test: divisible by (\d+)`)
	testCG := testRE.FindStringSubmatch(description[3])
	test, err := strconv.Atoi(testCG[1])
	aoc.ReportError(err)
	trueRE := regexp.MustCompile(`true: throw to monkey (\d+)`)
	trueCG := trueRE.FindStringSubmatch(description[4])
	trueMonkey, err := strconv.Atoi(trueCG[1])
	aoc.ReportError(err)
	falseRE := regexp.MustCompile(`false: throw to monkey (\d+)`)
	falseCG := falseRE.FindStringSubmatch(description[5])
	falseMonkey, err := strconv.Atoi(falseCG[1])
	aoc.ReportError(err)
	monkey.test = func(x int) int {
		if x%test == 0 {
			return trueMonkey
		} else {
			return falseMonkey
		}
	}

	return
}
func (m *monkey) inspect(worryLevel int) int {
	m.inspectionCount++
	newWL := m.operator(worryLevel)
	newWL /= worryReductionFactor
	return newWL
}
func (m *monkey) addItem(worryLevel int) {
	m.items = append(m.items, worryLevel)
}
func (m *monkey) clearItems() {
	m.items = []int{}
}
func getMostActiveMonkeys(monkeys []monkey, count int) []monkey {
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].inspectionCount > monkeys[j].inspectionCount })

	return monkeys[0:count]
}
