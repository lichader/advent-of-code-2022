package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	priorities := initPrioritiesMap()
	fmt.Println(priorities)

	const sampleFile = "./input.txt"
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
	}

	var sum int

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("line: ", line)

		lettersCount := countCommonChars(line)

		for _, char := range lettersCount {
			priority := priorities[char]
			fmt.Println("Char priority: ", char, priority)
			sum = sum + priority
		}
	}

	fmt.Println("The sum of priority is ", sum)
}

func initPrioritiesMap() map[string]int {
	prioritiesDic := make(map[string]int)
	// a-z, 1-26
	for i := 97; i <= 122; i++ {
		character := string(i)
		prioritiesDic[character] = i - 96
	}

	// A-Z, 27-52
	for i := 65; i <= 90; i++ {
		character := string(i)
		prioritiesDic[character] = i - 64 + 26
	}

	return prioritiesDic
}

func countCommonChars(line string) []string {
	dict := make(map[string]int)

	var ret []string

	var halfSize int = len(line) / 2
	for i := 0; i < len(line); i++ {
		if i < halfSize {
			// initialize
			leftChar := string(line[i])
			dict[leftChar] = 0
		} else {
			rightChar := string(line[i])

			value, inLeft := dict[rightChar]

			if value == 0 && inLeft {
				dict[rightChar] = 1
				fmt.Println("Found common character: ", rightChar)
				ret = append(ret, rightChar)
			}
		}
	}

	return ret
}

// BdbzzddChsWrRFbzBrszbhW
// MLNJHLLLLHZtSLglFNZHLJH
