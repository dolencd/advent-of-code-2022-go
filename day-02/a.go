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
	winOrLose["AX"] = 3
	winOrLose["AY"] = 6
	winOrLose["AZ"] = 0
	winOrLose["BX"] = 0
	winOrLose["BY"] = 3
	winOrLose["BZ"] = 6
	winOrLose["CX"] = 6
	winOrLose["CY"] = 0
	winOrLose["CZ"] = 3

	bytes, err := os.ReadFile("./day-02/input.txt")
	if (err != nil) {
		fmt.Println(err)
	}

	strs := strings.Split(string(bytes), "\n")
	totalScore := 0;
	for _, str := range(strs) {
		runes := []rune(str);
		totalScore += scoresPerHand[string(runes[2])] + winOrLose[string(runes[0]) + string(runes[2])]
	}
	fmt.Printf("totalScore: %v\n", totalScore)
}