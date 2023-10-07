package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var totalPriorities = 0
	for sc.Scan() {
		var rucksack = sc.Text()
		rucksackLeft := []rune(rucksack[:len(rucksack)/2])
		rucksackRight := []rune(rucksack[len(rucksack)/2:])
	Exit:
		for _, rucksackLeftItem := range rucksackLeft {
			for _, rucksackRightItem := range rucksackRight {
				if rucksackLeftItem != rucksackRightItem {
					continue
				}

				itemValue := int(unicode.ToLower(rucksackLeftItem) - 96)

				if unicode.IsUpper(rucksackLeftItem) {
					totalPriorities += 26
				}
				totalPriorities += itemValue
				break Exit
			}
		}
	}
	fmt.Println(totalPriorities)
}
