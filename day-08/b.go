package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

type forest struct {
	trees [][]rune
}

func (f *forest) viewLeft(cx, cy int) int {
	ownHeight := f.trees[cx][cy]
	total := 0
	for x := cx - 1; x >= 0; x-- {
		foundHeight := f.trees[x][cy]
		if foundHeight > ownHeight {
			break
		}
		total++
		if foundHeight == ownHeight {
			break
		}
	}
	return total
}

func (f *forest) viewRight(cx, cy int) int {
	ownHeight := f.trees[cx][cy]
	total := 0
	for x := cx + 1; x < len(f.trees); x++ {
		foundHeight := f.trees[x][cy]
		if foundHeight > ownHeight {
			break
		}
		total++
		if foundHeight == ownHeight {
			break
		}
	}
	return total
}

func (f *forest) viewUp(cx, cy int) int {
	ownHeight := f.trees[cx][cy]
	total := 0
	for y := cy - 1; y >= 0; y-- {
		foundHeight := f.trees[cx][y]
		if foundHeight > ownHeight {
			break
		}
		total++
		if foundHeight == ownHeight {
			break
		}
	}
	return total
}
func (f *forest) viewDown(cx, cy int) int {
	ownHeight := f.trees[cx][cy]
	total := 0
	for y := cy + 1; y < len(f.trees[0]); y++ {
		foundHeight := f.trees[cx][y]
		if foundHeight > ownHeight {
			break
		}
		total++
		if foundHeight == ownHeight {
			break
		}
	}
	return total
}

func (f *forest) width() int {
	return len(f.trees[0])
}

func (f *forest) height() int {
	return len(f.trees)
}

func createForest(sc *bufio.Scanner) forest {
	var trees [][]rune
	for sc.Scan() {
		row := []rune{}
		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		trees = append(trees, row)
	}
	return forest{trees}
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	forest := createForest(sc)

	maxScore := 0
	// Rows
	for x := 0; x < forest.width(); x++ {
		for y := 0; y < forest.height(); y++ {
			currentScore := forest.viewDown(x, y) * forest.viewLeft(x, y) * forest.viewRight(x, y) * forest.viewUp(x, y)
			if currentScore > maxScore {
				maxScore = currentScore
			}
		}
	}

	fmt.Printf("maxScore: %v\n", maxScore)
}
