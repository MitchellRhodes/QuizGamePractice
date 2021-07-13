package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Quiz struct {
	Question string `csv:"Question"`
	Answer   string `csv:"Answer"`
}

func main() {
	quizReader()
	// quiz := Quiz{
	// 	Question: records[0], //outputs question row
	// 	Answer:   records[1], //outputs answer row
	// }
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
		records, err := reader.Read()

		//EOF is the error returned by Read when no more input
		//is available. This breaks at end of read
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//print file
		fmt.Println(records[0])
	}

}
