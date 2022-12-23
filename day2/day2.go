package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shape struct {
	name   string
	score  int
	winOn  string
	loseOn string
}

const (
	shapeRock      = "rock"
	shapeScissor   = "scissor"
	shapePaper     = "paper"
	scoreWin       = 6
	scoreDraw      = 3
	myStrategyWin  = "Z"
	myStrategyDarw = "Y"
	myStrategyLose = "X"
)

func main() {
	const sampleFile = "./day2_input.txt"
	fileStream, err := os.Open(sampleFile)
	if err != nil {
		fmt.Println("Unable to read the file", err)
		return
	}

	rock := Shape{shapeRock, 1, shapeScissor, shapePaper}
	paper := Shape{shapePaper, 2, shapeRock, shapeScissor}
	scissor := Shape{shapeScissor, 3, shapePaper, shapeRock}

	optionsMap := map[string]Shape{
		"A":          rock,
		"B":          paper,
		"C":          scissor,
		shapeRock:    rock,
		shapePaper:   paper,
		shapeScissor: scissor,
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	var score int
	for fileScanner.Scan() {

		rawInput := fileScanner.Text()

		yourOption, myOption := getShapes(rawInput)

		yourShape := optionsMap[yourOption]

		var myShape Shape
		if myOption == myStrategyWin {
			myShape = optionsMap[yourShape.loseOn]
		} else if myOption == myStrategyDarw {
			myShape = yourShape
		} else {
			myShape = optionsMap[yourShape.winOn]
		}

		fmt.Println("Your option and my option:", yourShape, myShape)

		score = score + myShape.score

		if myShape.winOn == yourShape.name {
			score = score + scoreWin
		} else if myShape == yourShape {
			score = score + scoreDraw
		}

		fmt.Println("New score: ", score)

	}

	fmt.Println("Final score: ", score)
}

func getShapes(input string) (string, string) {
	slice := strings.Split(input, " ")

	return slice[0], slice[1]
}
