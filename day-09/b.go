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

	snake := make([]point, 10)

	tailPositions := make(map[point]bool)
	tailPositions[snake[len(snake)-1]] = true
	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moveDistance, err := strconv.Atoi(sc.Text()[2:])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < moveDistance; i++ {
			switch direction {
			case 'L':
				snake[0].x--
			case 'R':
				snake[0].x++
			case 'U':
				snake[0].y++
			case 'D':
				snake[0].y--
			}

			snake = followHead(snake)
			tailPositions[snake[len(snake)-1]] = true
		}
	}

	fmt.Printf("len(tailPositions): %v\n", len(tailPositions))
}

func followHead(snake []point) (newsnake []point) {
	newSnake := snake
	for i := 1; i < len(snake); i++ {
		front := newSnake[i-1]
		rear := newSnake[i]
		switch (point{x: front.x - rear.x, y: front.y - rear.y}) {
		case point{x: -2, y: 2}, point{x: -1, y: 2}, point{x: -2, y: 1}:
			rear.x--
			rear.y++
		case point{x: 0, y: 2}:
			rear.y++
		case point{x: 2, y: 2}, point{x: 1, y: 2}, point{x: 2, y: 1}:
			rear.x++
			rear.y++
		case point{x: 2, y: 0}:
			rear.x++
		case point{x: 2, y: -1}, point{x: 2, y: -2}, point{x: 1, y: -2}:
			rear.x++
			rear.y--
		case point{x: 0, y: -2}:
			rear.y--
		case point{x: -1, y: -2}, point{x: -2, y: -2}, point{x: -2, y: -1}:
			rear.x--
			rear.y--
		case point{x: -2, y: 0}:
			rear.x--
		}
		newSnake[i] = rear
	}

	return newSnake
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
