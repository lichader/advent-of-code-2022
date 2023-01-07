package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	moveRegex := regexp.MustCompile("^move (\\d+) from (\\d+) to (\\d+)$")
	crateRegex := regexp.MustCompile("^\\[[A-Z]{1}\\]$")
	const sampleFile = "./input.txt"
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	var stacks [9][]string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if moveRegex.MatchString(line) {
			res := moveRegex.FindAllStringSubmatch(line, -1)

			moves, _ := strconv.Atoi(res[0][1])
			sourceCrate, _ := strconv.Atoi(res[0][2])
			destCrate, _ := strconv.Atoi(res[0][3])

			fmt.Println("Parsing line to moves: ", moves, sourceCrate, destCrate)
		} else if strings.HasPrefix(line, "[") {
			fmt.Println("line: ", line, " and length: ", len(line))

			stackIndex := 0
			for i := 0; i < len(line); i += 4 {
				crate := line[i : i+3]

				if crateRegex.MatchString(crate) {
					crateName := crateRegex.FindStringSubmatch(crate)[0]

					stacks[stackIndex] = append([]string{crateName}, stacks[stackIndex]...)
				}

				fmt.Println(stackIndex)
				stackIndex++
			}
		} else if len(line) == 0 {
			for i, stack := range stacks {
				fmt.Println("index: ", i, ", stack: ", stack)
			}
		} else {
			fmt.Println("Skip line: ", line)
		}
	}
}
