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
	cycle      int
	x          int
	totalScore int
}

func (r *register) processCycle() {
	r.cycle += 1
	if (r.cycle-20)%40 == 0 {
		score := r.cycle * r.x
		fmt.Printf("In cycle %v got a score %v with x %v\n", r.cycle, score, r.x)
		r.totalScore += score
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
