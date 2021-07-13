package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Quiz struct {
	Question string `csv:"Question"`
	Answer   int    `csv:"Answer"`
}

const (
	question int = iota
	answer
)

func main() {
	quizReader()
}

func quizReader() {
	//open the file
	quizFile, err := os.Open("./questionsAndAnswers.csv")
	if err != nil {
		panic(err)
	}
	//initialize reader
	reader := csv.NewReader(quizFile)

	//discarding the header
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	//read all of the file
	for {
		row, err := reader.Read()

		//EOF is the error returned by Read when no more input
		//is available. This breaks at end of read
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//convert string answer to int
		answer, err := strconv.Atoi(row[answer])
		if err != nil {
			log.Fatal(err)
		}

		//print file
		fmt.Println(Quiz{
			Question: row[question],
			Answer:   answer,
		})
	}

}
