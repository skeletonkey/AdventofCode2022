package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_monkey(t *testing.T) {
	assert := assert.New(t)

	monkey1 := newMonkey([]string{
		"Monkey 0:",
		"  Starting items: 79, 98",
		"  Operation: new = old * 19",
		"  Test: divisible by 23",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 3",
	}, 0)
	assert.Equal(79, monkey1.items[0], "Correct items")
	assert.Equal(98, monkey1.items[1], "Correct items")
	assert.Equal(38, monkey1.operator(2), "Correct operator result")
	assert.Equal(2, monkey1.test(23), "Correct test result")

	worryLevel := monkey1.inspect(79)
	assert.Equal(500, worryLevel, "Correct Worry Level")
	assert.Equal(3, monkey1.test(worryLevel), "Throw to monkey 3")
	assert.Equal(1, monkey1.inspectionCount, "Inspected 1 item")

	monkey2 := newMonkey([]string{
		"Monkey 1:",
		"  Starting items: 74",
		"  Operation: new = old + 3",
		"  Test: divisible by 17",
		"    If true: throw to monkey 0",
		"    If false: throw to monkey 1",
	}, 1)
	assert.Equal(74, monkey2.items[0], "Correct items")
	assert.Equal(5, monkey2.operator(2), "Correct operator result")
	assert.Equal(1, monkey2.test(23), "Correct test result")

	monkey2.addItem(100)
	assert.Equal(100, monkey2.items[1], "New item at end of list")

	monkeys := []monkey{monkey1, monkey2}
	activeMonkeys := getMostActiveMonkeys(monkeys, 1)
	assert.Equal(1, len(activeMonkeys), "expected 1 monkey")
	assert.Equal(1, activeMonkeys[0].inspectionCount, "Inspected 1 item")

	monkey2.clearItems()
	assert.Equal(0, len(monkey2.items), "Items have been cleared")

	monkey2.addItem(563)
	assert.Equal(188, monkey2.inspect(563))
	assert.Equal(334, monkey2.inspect(1000))
	assert.Equal(333, monkey2.inspect(996))

	monkey3 := newMonkey([]string{
		"Monkey 3:",
		"Starting items: 79, 60, 97",
		"Operation: new = old * old",
		"Test: divisible by 13",
		"If true: throw to monkey 1",
		"If false: throw to monkey 3	",
	}, 3)

	assert.Equal(2080, monkey3.inspect(79), "Monkey 3 inspect 79")
}
