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
	isFile bool;
	name string;
	children []element;
	parent *element;
	size int;
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


func (e *element) getTotalSizeOfDirectoriesSmallerThan(maxSize int) int {
	totalSize := 0
	directories := e.getDirectoriesOfSizeLessThan(maxSize);
	for _, directory := range(directories) {
		fmt.Println("adding", directory.name, directory.size)
		totalSize += directory.size
	}
	return totalSize
}

func (e *element) getDirectoriesOfSizeLessThan(maxSize int) []*element {
	if (e.isFile) {
		return make([]*element, 0, 10)	
	}

	outputDirectories := make([]*element, 0, 10)
	totalSize := e.size;
	if  totalSize < maxSize {
		fmt.Println("got one", e.name)
		outputDirectories = append(outputDirectories, e);
	}
	for _, child := range(e.children) {
		if child.isFile {
			continue
		}
		newChildren := child.getDirectoriesOfSizeLessThan(maxSize)
		outputDirectories = append(outputDirectories, newChildren...)
	}
	
	return outputDirectories
}

// func (e *element) getTotalSize() int {
// 	if e.size != 0 {
// 		return e.size
// 	}
// 	size := 0
// 	for _, child := range(e.children) {
// 		size += child.getTotalSize()
// 	}
// 	return size
// }

func (e *element) addSizeToSelfAndAllParents(sizeToAdd int) {
	e.size += sizeToAdd
	if e.parent != nil {
		e.parent.addSizeToSelfAndAllParents(sizeToAdd)
	}
}

func main(){
	input,err := os.Open("./input.txt")
	if (err != nil) {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)
	
	root := element {
		isFile: false,
		name: "root",
	};

	directories := make([]*element, 1)
	directories[0] = &root

	current := &root

	for sc.Scan(){
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
			newNode := element {
				name: line[1],
				parent: current,
				isFile: false,
			}
			current.children = append(current.children, newNode)
			directories = append(directories, &newNode)
			continue
		}

		size, _ := strconv.Atoi(line[0])
		current.children = append(current.children, element {
			name: line[1],
			size: size,
			parent: current,
			isFile: true,
		})
		current.addSizeToSelfAndAllParents(size)
	}

	totalSize := 0
	for _, dir := range(root.findDirectories()) {
		size := dir.size
		fmt.Println("size", dir.name, size)
		if size < 100000 {
			totalSize += size
		}
	}

	fmt.Printf("totalSize: %v\n", totalSize)
	// fmt.Printf("root.getTotalSizeOfDirectoriesSmallerThan(100000): %v\n", root.getTotalSizeOfDirectoriesSmallerThan(100000))
}
