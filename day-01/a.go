package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./day-01/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	strs := strings.Split(string(bytes), "\n")
	currentSum := 0
	maxSum := 0
	for i := 0; i < len(strs); i++ {
		str := strs[i]

		if len(str) == 0 {
			if currentSum > maxSum {
				maxSum = currentSum
			}
			currentSum = 0
			continue
		}

		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		currentSum += num

	}
	fmt.Printf("maxSum: %v\n", maxSum)
}
