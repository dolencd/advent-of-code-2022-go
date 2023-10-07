package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type register struct {
	cycle       int
	x           int
	totalScore  int
	currentLine []string
}

func (r *register) processCycle() {

	column := r.cycle % 40

	if AbsDiff(column, r.x) < 2 {
		r.currentLine = append(r.currentLine, "#")
	} else {
		r.currentLine = append(r.currentLine, ".")
	}
	r.cycle += 1
	if column == 39 {
		fmt.Printf("Line: %s\n", strings.Join(r.currentLine, ""))
		r.currentLine = make([]string, 40)
	}
}

func (r *register) noop() {
	r.processCycle()
}

func (r *register) addx(x int) {
	r.processCycle()
	r.processCycle()
	r.x += x
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	register := register{
		x: 1,
	}

	for sc.Scan() {
		ops := strings.Fields(sc.Text())
		switch ops[0] {
		case "noop":
			register.noop()
		case "addx":
			x, err := strconv.Atoi(ops[1])
			if err != nil {
				log.Fatal(err)
			}
			register.addx(x)
		}
	}

	fmt.Printf("register: %v\n", register)
}

func AbsDiff(a int, b int) int {
	diff := a - b

	if diff < 0 {
		diff = -diff
	}

	return diff
}
