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
		if (s1 <= s2 && e1 >= e2) || (s2 <= s1 && e2 >= e1) {
			fullyOverlappingPairs++
		}

	}
	fmt.Println(fullyOverlappingPairs)
}
