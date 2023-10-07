package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type element struct {
	isFile   bool
	name     string
	children []element
	parent   *element
	size     int
}

func (e *element) findDirectories() []*element {
	var directories []*element

	var dfs func(current *element)
	dfs = func(current *element) {
		if !current.isFile {
			directories = append(directories, current)
		}
		for i := range current.children {
			dfs(&current.children[i])
		}
	}

	dfs(e)
	return directories
}

func (e *element) addSizeToSelfAndAllParents(sizeToAdd int) {
	e.size += sizeToAdd
	if e.parent != nil {
		e.parent.addSizeToSelfAndAllParents(sizeToAdd)
	}
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	root := element{
		isFile: false,
		name:   "root",
	}

	current := &root

	for sc.Scan() {
		lineString := sc.Text()
		line := strings.Fields(lineString)

		if lineString == "$ cd /" {
			continue
		}

		if line[0] == "$" {
			if line[1] == "ls" {
				continue
			}

			if line[1] == "cd" {
				if line[2] == ".." {
					current = current.parent
					continue
				}
				searchIndex := slices.IndexFunc(current.children, func(e element) bool {
					return e.name == line[2]
				})
				if searchIndex == -1 {
					fmt.Println("Failed to find child", searchIndex, current.name, lineString)
					panic("")
				}
				current = &current.children[searchIndex]
				continue
			}
		}

		if line[0] == "dir" {
			newNode := element{
				name:   line[1],
				parent: current,
				isFile: false,
			}
			current.children = append(current.children, newNode)
			continue
		}

		size, _ := strconv.Atoi(line[0])
		current.children = append(current.children, element{
			name:   line[1],
			size:   size,
			parent: current,
			isFile: true,
		})
		current.addSizeToSelfAndAllParents(size)
	}

	directories := root.findDirectories()

	fmt.Printf("usedSize: %v\n", root.size)

	TOTAL_SIZE_TO_FREE := root.size - (70000000 - 30000000)
	fmt.Printf("TOTAL_SIZE_TO_FREE: %v\n", TOTAL_SIZE_TO_FREE)
	minDirSize := 99999999
	for _, dir := range directories {
		if dir.size >= TOTAL_SIZE_TO_FREE && dir.size < minDirSize {
			minDirSize = dir.size
			fmt.Printf("minDirSize: %v %v\n", minDirSize, dir.name)
		}
	}
}
