package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main(){
	bytes, err := os.ReadFile("./day-01/input.txt")
	if (err != nil) {
		fmt.Println(err)
	}
	strs := strings.Split(string(bytes), "\n")
	var caloriesPerElf []int;
	currentSum := 0;
	for i := 0; i < len(strs); i++ {
		str := strs[i];

		if len(str) == 0 {
			caloriesPerElf = append(caloriesPerElf, currentSum);
			currentSum = 0;
			continue;
		}

		num, err := strconv.Atoi(str);
		if err != nil {
			panic(err);
		}
		
		currentSum += num;
	}
	slices.Sort(caloriesPerElf);
	fmt.Printf("caloriesPerElf: %v\n", caloriesPerElf);

	fmt.Println("sum: ", caloriesPerElf[len(caloriesPerElf)-3] + caloriesPerElf[len(caloriesPerElf)-1] + caloriesPerElf[len(caloriesPerElf)-2]);
}