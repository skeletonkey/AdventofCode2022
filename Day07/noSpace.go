package main

import (
	"fmt"
	aoc "github/skeletonkey/AdventofCode2022/adventOfCode"
	"os"
	"regexp"
	"strconv"
)

type mapOfDirs map[string]*dir
type mapOfFiles map[string]int

func main() {
	fs := aoc.GetData(os.Getwd())
	defer aoc.Cleanup()

	cdRE := regexp.MustCompile(`^\$ cd (\S+)$`)
	lsString := "$ ls"
	lsOutputRE := regexp.MustCompile(`^(dir|\d+)\s+(\S+)$`)

	var curDir *dir
	rootDir := newDir("/", nil)
	for fs.Scan() {
		if matches := cdRE.FindStringSubmatch(fs.Text()); matches != nil {
			if matches[1] == "/" {
				curDir = rootDir
			} else if matches[1] == ".." {
				curDir = curDir.parentDir
			} else {
				curDir = curDir.subDir(matches[1])
			}
		} else if matches := lsOutputRE.FindStringSubmatch(fs.Text()); matches != nil {
			if matches[1] == "dir" {
				_ = curDir.subDir(matches[2])
			} else {
				size, err := strconv.Atoi(matches[1])
				aoc.ReportError(err)
				curDir.addFile(matches[1], size)
			}
		} else {
			if fs.Text() != lsString {
				panic("Unparsed input: " + fs.Text())
			}
		}
	}

	calculateFullSize(rootDir)

	part1 := calculatePart1(rootDir)
	part2 := calculatePart2(rootDir)

	fmt.Printf("Part 1 total is %d\n", part1)
	fmt.Printf("Part 2 total is %d\n", part2)
}

func calculateFullSize(d *dir) {
	d.sizeFull = d.size // size of files
	for _, sd := range d.subDirs {
		calculateFullSize(sd)
		d.sizeFull += sd.sizeFull
	}
}

const part1SizeMax = 100000

func calculatePart1(d *dir) (total int) {
	if d.sizeFull <= part1SizeMax {
		total += d.sizeFull
	}
	for _, sd := range d.subDirs {
		total += calculatePart1(sd)
	}

	return
}

const totalSpace = 70000000
const updateSpaceRequired = 30000000

func calculatePart2(d *dir) (size int) {
	spaceNeeded := updateSpaceRequired - (totalSpace - d.sizeFull)

	return findMinDirSize(d, spaceNeeded, d.sizeFull)
}
func findMinDirSize(d *dir, minNeeded int, currentMinDir int) int {
	if d.sizeFull < currentMinDir && d.sizeFull >= minNeeded {
		currentMinDir = d.sizeFull
	}
	for _, sd := range d.subDirs {
		currentMinDir = findMinDirSize(sd, minNeeded, currentMinDir)
	}
	return currentMinDir
}

type dir struct {
	files     mapOfFiles
	name      string
	parentDir *dir
	size      int // size of only this dir
	sizeFull  int // size including subdirs
	subDirs   mapOfDirs
}

func newDir(name string, parentDir *dir) *dir {
	return &dir{
		files:     make(mapOfFiles, 0),
		name:      name,
		parentDir: parentDir,
		size:      0,
		sizeFull:  0,
		subDirs:   make(mapOfDirs, 0)}
}

func (d *dir) subDir(name string) *dir {
	if _, ok := d.subDirs[name]; !ok {
		d.subDirs[name] = newDir(name, d)
	}
	return d.subDirs[name]
}

func (d *dir) addFile(name string, size int) {
	if _, ok := d.files[name]; ok {
		panic(fmt.Sprintf("File (%s) already exists for dir %s\n", name, d.name))
	}
	d.files[name] = size
	d.size += size
}
