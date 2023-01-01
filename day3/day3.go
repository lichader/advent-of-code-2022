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

	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()

		lines = append(lines, line)

		if len(lines) == 3 {
			fmt.Println("lines: ", lines)

			lettersCount := countCommonChars(lines)

			for _, char := range lettersCount {
				priority := priorities[char]
				fmt.Println("Char priority: ", char, priority)
				sum = sum + priority
			}

			lines = make([]string, 0)
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

func countCommonChars(lines []string) []string {
	dict := make(map[string]int)

	var ret []string

	for index, line := range lines {
		if index == 0 {
			// initialize the dictionary
			for i := 0; i < len(line); i++ {
				char := string(line[i])
				dict[char] = 1
			}
		} else if index == 1 {
			// check 2nd line
			for i := 0; i < len(line); i++ {
				char := string(line[i])

				value, exists := dict[char]

				if value == 1 && exists {
					dict[char] = 2
				}

			}
		} else {
			// check 3rd line
			for i := 0; i < len(line); i++ {
				char := string(line[i])

				value, exists := dict[char]

				if value == 2 && exists {
					dict[char] = 3
					fmt.Println("Found common character", char)

					ret = append(ret, char)
				}

			}
		}
	}

	return ret
}

// BdbzzddChsWrRFbzBrszbhW
// MLNJHLLLLHZtSLglFNZHLJH
