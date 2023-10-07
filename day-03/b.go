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
		rucksack1 := []rune(sc.Text())
		sc.Scan()
		rucksack2 := []rune(sc.Text())
		sc.Scan()
		rucksack3 := []rune(sc.Text())
	Exit:
		for _, r1 := range rucksack1 {
			for _, r2 := range rucksack2 {
				for _, r3 := range rucksack3 {
					if r1 == r2 && r2 == r3 {
						itemValue := int(unicode.ToLower(r1) - 96)
						if unicode.IsUpper(r1) {
							totalPriorities += 26
						}
						totalPriorities += itemValue
						// fmt.Println(itemValue)
						break Exit
					}
				}
			}
		}
	}
	fmt.Println(totalPriorities)
}
