package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var fullyOverlappingPairs = 0
	for sc.Scan() {
		var s1, e1, s2, e2 int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
		if e1 >= s2 && s1 <= e2 {
			fullyOverlappingPairs++
		}

	}
	fmt.Println(fullyOverlappingPairs)
}
