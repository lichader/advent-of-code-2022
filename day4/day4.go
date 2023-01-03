package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const sampleFile = "./input.txt"
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	var count int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		slice := strings.Split(line, ",")

		r1Low, r1High := splitRange(slice[0])
		r2Low, r2High := splitRange(slice[1])

		if (r2Low <= r1Low && r1High <= r2High) || (r1Low <= r2Low && r2High <= r1High) {
			count++
			fmt.Println("Found range included: ", line)
		}
	}

	fmt.Println("Total fully contained pairs: ", count)
}

func splitRange(r string) (int, int) {
	splitted := strings.Split(r, "-")

	low, _ := strconv.Atoi(splitted[0])
	high, _ := strconv.Atoi(splitted[1])

	return low, high
}
