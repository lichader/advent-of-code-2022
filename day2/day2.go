package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("test")

	const sampleFile = "./day2_input.txt"
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
		return
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		rawInput := fileScanner.Text()

		fmt.Println(rawInput)
	}
}
