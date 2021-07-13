package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	quizReader()
}

func quizReader() {
	//open the file
	quizFile, err := os.Open("./questionsAndAnswers.csv")
	if err != nil {
		fmt.Println("Error occured ::", err)
	}
	//initialize reader
	reader := csv.NewReader(quizFile)

	//read all file
	quiz, _ := reader.ReadAll()

	//print file
	fmt.Println(quiz)
}
