package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	inspectionCount int
	items           []int
	operator        rune
	operationValue  int
	divisibleCheck  int
	throwToIfTrue   int
	throwToIfFalse  int
}

func (m *monkey) applyOperationToItem(item int) int {
	var operationValue int

	if m.operationValue == 0 {
		operationValue = item
	} else {
		operationValue = m.operationValue
	}

	var result int
	switch m.operator {
	case '+':
		result = item + operationValue
	case '*':
		result = item * operationValue
	default:
		panic("Unsupported operator")
	}

	m.inspectionCount++

	return result / 3
}

func (m *monkey) determineMonkeyToThrowTo(item int) int {
	if item%m.divisibleCheck == 0 {
		return m.throwToIfTrue
	} else {
		return m.throwToIfFalse
	}
}

func (m *monkey) processRound(allMonkeys []*monkey) {
	for _, item := range m.items {
		newItem := m.applyOperationToItem(item)
		nextMonkey := m.determineMonkeyToThrowTo(newItem)
		// fmt.Printf("item with score %v will be thrown to %v with new score %v\n", item, nextMonkey, newItem)
		allMonkeys[nextMonkey].items = append(allMonkeys[nextMonkey].items, newItem)
	}
	m.items = make([]int, 0)
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	var monkeyRefs []*monkey

	for sc.Scan() {
		sc.Scan()
		var newMonkey monkey
		for _, item := range strings.Split(sc.Text()[len("  Starting items: "):], ", ") {
			worryLevel, _ := strconv.Atoi(item)
			newMonkey.items = append(newMonkey.items, worryLevel)
		}

		sc.Scan()
		var possibleOtherValue string
		fmt.Sscanf(sc.Text(), "  Operation: new = old %c %s", &newMonkey.operator, &possibleOtherValue)

		if possibleOtherValue == "old" {
			newMonkey.operationValue = 0
		} else {
			newMonkey.operationValue, _ = strconv.Atoi(possibleOtherValue)
		}

		sc.Scan()
		fmt.Sscanf(sc.Text(), "  Test: divisible by %d", &newMonkey.divisibleCheck)

		sc.Scan()
		fmt.Sscanf(sc.Text(), "    If true: throw to monkey %d", &newMonkey.throwToIfTrue)
		sc.Scan()
		fmt.Sscanf(sc.Text(), "    If false: throw to monkey %d", &newMonkey.throwToIfFalse)

		sc.Scan()
		monkeyRefs = append(monkeyRefs, &newMonkey)
	}

	for round := 0; round < 20; round++ {
		// fmt.Printf("\nRound: %v\n", round+1)
		for _, monkey := range monkeyRefs {
			monkey.processRound(monkeyRefs)
		}

		// for i, monkey := range monkeyRefs {
		// 	fmt.Printf("monkey %v items: %v\n", i, monkey.items)
		// }
	}

	// for i, monkey := range monkeyRefs {
	// 	fmt.Printf("monkey %v inspected times: %v\n", i, monkey.inspectionCount)
	// }

	scores := make([]int, len(monkeyRefs))
	for i, monkey := range monkeyRefs {
		scores[i] = monkey.inspectionCount
	}

	sort.Ints(scores)
	fmt.Printf("scores: %v\n", scores)

	fmt.Printf("(scores[len(scores)-1] * scores[len(scores)-2]): %v\n", (scores[len(scores)-1] * scores[len(scores)-2]))
}
