package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	head := point{0, 0}
	tail := head
	tailPositions := make(map[point]bool)
	tailPositions[tail] = true
	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moveDistance, err := strconv.Atoi(sc.Text()[2:])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < moveDistance; i++ {
			oldHead := head

			switch direction {
			case 'L':
				head.x--
			case 'R':
				head.x++
			case 'U':
				head.y++
			case 'D':
				head.y--
			}

			if !isAdjacent(tail, head) {
				tail = oldHead
				tailPositions[tail] = true
			}
		}
	}

	fmt.Printf("len(tailPositions): %v\n", len(tailPositions))
}

func isAdjacent(a point, b point) bool {
	return Abs(b.x-a.x) <= 1 && Abs(b.y-a.y) <= 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
