package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	scoresPerHand := make(map[string]int);
	scoresPerHand["X"] = 1; // A Rock
	scoresPerHand["Y"] = 2; // B Paper
	scoresPerHand["Z"] = 3; // C Scissors

	winOrLose := make(map[string]int);
	winOrLose["AX"] = 3 + 0
	winOrLose["AY"] = 1 + 3
	winOrLose["AZ"] = 2 + 6
	winOrLose["BX"] = 1 + 0
	winOrLose["BY"] = 2 + 3
	winOrLose["BZ"] = 3 + 6
	winOrLose["CX"] = 2 + 0
	winOrLose["CY"] = 3 + 3
	winOrLose["CZ"] = 1 + 6

	bytes, err := os.ReadFile("./day-02/input.txt")
	if (err != nil) {
		fmt.Println(err)
	}

	strs := strings.Split(string(bytes), "\n")
	totalScore := 0;
	for _, str := range(strs) {
		runes := []rune(str);
		totalScore += winOrLose[string(runes[0]) + string(runes[2])]
	}
	fmt.Printf("totalScore: %v\n", totalScore)
}