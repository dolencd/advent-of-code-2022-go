package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	var forest [][]rune
	for sc.Scan() {
		row := []rune{}
		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	maxX := len(forest)
	maxY := len(forest[0])
	isVisible := make(map[point]bool)

	// Rows
	for x := 0; x < maxX; x++ {
		maxSoFar := -1
		for y := 0; y < maxY; y++ {
			// Left to right
			currentTreeHeight := int(forest[x][y]) - 48
			if currentTreeHeight > maxSoFar {
				isVisible[point{x, y}] = true
				maxSoFar = currentTreeHeight
			}
		}

		maxSoFar = -1
		for y := maxY - 1; y >= 0; y-- {
			// Right to left
			currentTreeHeight := int(forest[x][y]) - 48
			if currentTreeHeight > maxSoFar {
				isVisible[point{x, y}] = true
				maxSoFar = currentTreeHeight
			}
		}
	}

	for y := 0; y < maxY; y++ {
		maxSoFar := -1
		for x := 0; x < maxX; x++ {
			// Top to bottom
			currentTreeHeight := int(forest[x][y]) - 48
			if currentTreeHeight > maxSoFar {
				isVisible[point{x, y}] = true
				maxSoFar = currentTreeHeight
			}
		}

		maxSoFar = -1
		for x := maxX - 1; x >= 0; x-- {
			// Bottom to top
			currentTreeHeight := int(forest[x][y]) - 48
			if currentTreeHeight > maxSoFar {
				isVisible[point{x, y}] = true
				maxSoFar = currentTreeHeight
			}
		}
	}

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			tmp := isVisible[point{x, y}]
			if tmp {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	fmt.Printf("len(isVisible): %v\n", len(isVisible))
}
