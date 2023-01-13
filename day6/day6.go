package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 14

func main() {
	sampleFile := os.Args[1]
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fmt.Println("Read line", line)

		for i := 0; i <= len(line)-size; i++ {
			subString := line[i:(i + size)]

			fmt.Println(subString)

			if uniqueChars(subString) {
				fmt.Println("Index found: ", i+size)
				break
			}
		}
	}

	fileStream.Close()
}

func uniqueChars(input string) bool {
	m := make(map[rune]int)

	for i, c := range input {
		m[c] = i
	}

	if len(m) == size {
		return true
	}

	return false
}
