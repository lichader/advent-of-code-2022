package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file is given in the command line")
		return
	}

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var allCalories []int
	allCalories = append(allCalories, 0)
	index := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		rawInput := fileScanner.Text()

		if len(rawInput) == 0 {
			fmt.Println("Encountered space. Skip")
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

	fmt.Println(allCalories)

	var max int
	for i, e := range allCalories {
		if i == 0 || e > max {
			max = e
		}
	}

	fmt.Println(max)
}
