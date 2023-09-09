package main

import (
	"bufio"
	"fmt"
	"os"
)

const MARKER_WIDTH = 14

func main(){
	input,err := os.Open("./input.txt")
	if (err != nil) {
		fmt.Println((err))
	}
	defer input.Close()
	sc := bufio.NewScanner(input)
	
	total := 0

	for sc.Scan(){
		runeArr := []rune(sc.Text())
		Outer:
		for i := range runeArr {
			if i < MARKER_WIDTH-1 {
				continue;
			}

			characters := make(map[rune]bool, 4);
			for j := 0; j < MARKER_WIDTH; j++ {
				currentCharacter := runeArr[i - j]
				if (characters[currentCharacter]) {
					continue Outer
				} 
				characters[currentCharacter] = true;
			}

			total += i + 1
			break;
		}
	}

	fmt.Printf("total: %v\n", total)
}