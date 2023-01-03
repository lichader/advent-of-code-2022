package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const sampleFile = "./input.txt"
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
	}
}
