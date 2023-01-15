package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"regexp"
	"strconv"
)

const (
	humanName = "humn"
	rootName  = "root"
)

type collectionType map[string]*node

type node struct {
	name       string
	value      int
	valueSet   bool
	args       []string
	dependants []string
	operator   string
}

func newNode(name string) *node {
	return &node{
		name:       name,
		value:      0,
		valueSet:   false,
		args:       make([]string, 2),
		dependants: make([]string, 0),
		operator:   "",
	}
}

func (n *node) addDependant(name string) {
	n.dependants = append(n.dependants, name)
}
func (n *node) setValue(nodes map[string]*node) {
	node1, _ := nodes[n.args[0]]
	node2, _ := nodes[n.args[1]]
	if node1.valueSet && node2.valueSet {
		switch n.operator {
		case "+":
			n.value = node1.value + node2.value
		case "-":
			n.value = node1.value - node2.value
		case "*":
			n.value = node1.value * node2.value
		case "/":
			n.value = node1.value / node2.value
		case "=":
			n.value = node1.value - node2.value
		default:
			aoc.ReportError(fmt.Errorf("unknown operator (%s) with args: %d, %d", n.operator, node1.value, node2.value))
		}

		n.valueSet = true
	}
}

//	func (n *node) setDependency(name string) {
//		n.dependencies = append(n.dependencies, name)
//	}
func (n *node) resolveDependants(nodes collectionType) {
	if n.valueSet {
		for _, nodeName := range n.dependants {
			nodes[nodeName].setValue(nodes)
			if nodes[nodeName].valueSet {
				nodes[nodeName].resolveDependants(nodes)
			}
		}
	}
}

func main() {
	//part1()
	//part2()
	part2Math()
}

func part1() {
	nodeCollection := make(collectionType, 0)
	nodeCollection[rootName] = newNode(rootName)

	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	mathRE := regexp.MustCompile(`(\w+): (\w+) ([+-/*]) (\w+)`)
	answerRE := regexp.MustCompile(`(\w+): (\d+)`)
	for fs.Scan() {
		if mathRE.Match(fs.Bytes()) {
			parts := mathRE.FindStringSubmatch(fs.Text())

			if _, exists := nodeCollection[parts[1]]; !exists {
				nodeCollection[parts[1]] = newNode(parts[1])
			}
			if _, exists := nodeCollection[parts[2]]; !exists {
				nodeCollection[parts[2]] = newNode(parts[2])
			}
			if _, exists := nodeCollection[parts[4]]; !exists {
				nodeCollection[parts[4]] = newNode(parts[4])
			}
			nodeCollection[parts[1]].operator = parts[3]
			nodeCollection[parts[1]].args = []string{parts[2], parts[4]}
			nodeCollection[parts[1]].setValue(nodeCollection)
			if nodeCollection[parts[1]].valueSet {
				if len(nodeCollection[parts[1]].dependants) > 0 {
					nodeCollection[parts[1]].resolveDependants(nodeCollection)
				}
			} else {
				if !nodeCollection[parts[2]].valueSet {
					nodeCollection[parts[2]].addDependant(parts[1])
				}
				if !nodeCollection[parts[4]].valueSet {
					nodeCollection[parts[4]].addDependant(parts[1])
				}
			}
		} else {
			parts := answerRE.FindStringSubmatch(fs.Text())
			if _, exists := nodeCollection[parts[1]]; !exists {
				nodeCollection[parts[1]] = newNode(parts[1])
			}
			nodeValue, err := strconv.Atoi(parts[2])
			aoc.ReportError(err)
			nodeCollection[parts[1]].value = nodeValue
			nodeCollection[parts[1]].valueSet = true
			nodeCollection[parts[1]].resolveDependants(nodeCollection)
		}
		if nodeCollection[rootName].valueSet {
			break
		}
	}

	fmt.Printf("Part 1: Root value is %d\n", nodeCollection[rootName].value)
}

func part2() {
	number := 0
	for !part2Exec(number) {
		number++
		fmt.Printf("Trying: %d\r", number)
	}

	fmt.Printf("Part 2: I need to yell %d\n", number)
}
func part2Exec(start int) bool {
	nodeCollection := make(collectionType, 0)
	nodeCollection[rootName] = newNode(rootName)
	nodeCollection[humanName] = newNode(humanName)
	nodeCollection[humanName].value = start
	nodeCollection[humanName].valueSet = true

	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	mathRE := regexp.MustCompile(`(\w+): (\w+) ([=+-/*]) (\w+)`)
	answerRE := regexp.MustCompile(`(\w+): (\d+)`)
	for fs.Scan() {
		if mathRE.Match(fs.Bytes()) {
			parts := mathRE.FindStringSubmatch(fs.Text())

			if _, exists := nodeCollection[parts[1]]; !exists {
				nodeCollection[parts[1]] = newNode(parts[1])
			}
			if _, exists := nodeCollection[parts[2]]; !exists {
				nodeCollection[parts[2]] = newNode(parts[2])
			}
			if _, exists := nodeCollection[parts[4]]; !exists {
				nodeCollection[parts[4]] = newNode(parts[4])
			}
			nodeCollection[parts[1]].operator = parts[3]
			nodeCollection[parts[1]].args = []string{parts[2], parts[4]}
			nodeCollection[parts[1]].setValue(nodeCollection)
			if nodeCollection[parts[1]].valueSet {
				if len(nodeCollection[parts[1]].dependants) > 0 {
					nodeCollection[parts[1]].resolveDependants(nodeCollection)
				}
			} else {
				if !nodeCollection[parts[2]].valueSet {
					nodeCollection[parts[2]].addDependant(parts[1])
				}
				if !nodeCollection[parts[4]].valueSet {
					nodeCollection[parts[4]].addDependant(parts[1])
				}
			}
		} else {
			parts := answerRE.FindStringSubmatch(fs.Text())
			if _, exists := nodeCollection[parts[1]]; !exists {
				nodeCollection[parts[1]] = newNode(parts[1])
			}
			nodeValue, err := strconv.Atoi(parts[2])
			aoc.ReportError(err)
			nodeCollection[parts[1]].value = nodeValue
			nodeCollection[parts[1]].valueSet = true
			nodeCollection[parts[1]].resolveDependants(nodeCollection)
		}
		if nodeCollection[rootName].valueSet {
			break
		}
	}

	return nodeCollection[rootName].value == 0
}

func part2Math() {
	nodeCollection := make(collectionType, 0)

	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	humnLocation := ""
	mathRE := regexp.MustCompile(`(\w+): (\w+) ([=+-/*]) (\w+)`)
	answerRE := regexp.MustCompile(`(\w+): (\d+)`)
	for fs.Scan() {
		if mathRE.Match(fs.Bytes()) {
			parts := mathRE.FindStringSubmatch(fs.Text())
			nodeCollection[parts[1]] = newNode(parts[1])
			nodeCollection[parts[1]].operator = parts[3]
			nodeCollection[parts[1]].args = []string{parts[2], parts[4]}
			if humnLocation == "" && (parts[2] == humanName || parts[4] == humanName) {
				humnLocation = parts[1]
			}
		} else {
			parts := answerRE.FindStringSubmatch(fs.Text())
			nodeCollection[parts[1]] = newNode(parts[1])
			nodeValue, err := strconv.Atoi(parts[2])
			aoc.ReportError(err)
			nodeCollection[parts[1]].value = nodeValue
			nodeCollection[parts[1]].valueSet = true
		}
	}

	val1 := calulateNum(nodeCollection, nodeCollection[rootName].args[0])
	val2 := calulateNum(nodeCollection, nodeCollection[rootName].args[1])
	if part2Exec(val1 - val2) {
		fmt.Printf("Part 2: I need to yell %d\n", val1-val2)
	} else if part2Exec(val2 - val1) {
		fmt.Printf("Part 2: I need to yell %d\n", val2-val1)
	} else {
		aoc.ReportError(fmt.Errorf("%d and %d are the wrong numbers\n", val1, val2))
	}
}

func calulateNum(nodes collectionType, name string) (value int) {
	if nodes[name].valueSet {
		return nodes[name].value
	} else {

	}
	return
}
