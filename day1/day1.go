package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	const sampleFile = "./day1_input.txt"
	readFile, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var allCalories []int
	allCalories = append(allCalories, 0)
	index := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		rawInput := fileScanner.Text()

		if len(rawInput) == 0 {
			index++
			allCalories = append(allCalories, 0)
			continue
		}

		calories, err := strconv.Atoi(rawInput)
		if err != nil {
			errorMessage := fmt.Sprintf("Received an error on parsing %s", rawInput)
			fmt.Println(errorMessage)
		}

		allCalories[index] = allCalories[index] + calories
	}

	readFile.Close()

	sort.Ints(allCalories)
	fmt.Println(allCalories)

	fmt.Println(fmt.Sprintf("The maximum calrories are: %d", allCalories[len(allCalories)-1]))

	var topThree []int = allCalories[(len(allCalories) - 3):]

	fmt.Println("Top 3 are: ", topThree)

	sumOfTop3 := 0
	for _, v := range topThree {
		sumOfTop3 += v
	}

	fmt.Println("The total calories of top 3 are: ", sumOfTop3)
}
