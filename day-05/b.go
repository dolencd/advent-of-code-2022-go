package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	crates []rune
}

func fromString(str string) stack {
	return stack{
		crates: []rune(str),
	}
}

func (s *stack) pop(count int) []rune {
	last := s.crates[len(s.crates)-count:]
	s.crates = s.crates[:len(s.crates)-count]
	return last
}

func (s *stack) push(r []rune) {
	s.crates = append(s.crates, r...)
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)
	stacks := []stack{
		// fromString("ZN"),
		// fromString("MCD"),
		// fromString("P"),
		fromString("STHFWR"),
		fromString("SGDQW"),
		fromString("BTW"),
		fromString("DRWTNQZJ"),
		fromString("FBHGLVTZ"),
		fromString("LPTCVBSG"),
		fromString("ZBRTWGP"),
		fromString("NGMTCJR"),
		fromString("LGBW"),
	}
	for sc.Scan() {
		var move, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &move, &from, &to)
		stacks[to-1].push(stacks[from-1].pop(move))
	}

	for i := range stacks {
		stack := stacks[i]
		fmt.Print(string(stack.pop(1)))
	}
}
