package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Quiz struct {
	Question string `csv:"Question"`
	Answer   string `csv:"Answer"`
}

func main() {
	csvFilename := flag.String("csv", "./addition.csv",
		"A csv in the format of 'Question,Answer'. subtraction.csv and multiplication.csv are also available")
	flag.Parse()
	quizReader(*csvFilename)

}

func quizReader(fileName string) {
	//open the file
	quizFile, err := os.Open(fileName)
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
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//attributes the newly made struct to a variable
	quiz := parseRows(rows)
	var numberCorrect int = 0

	//loop through the struct to output questions
	//i+1 is so that we skip the 0 index, which is a header
	for i, problem := range quiz {
		fmt.Printf("Solve #%d: %s = \n", i+1, problem.Question)

		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == problem.Answer {
			numberCorrect++
		}
	}

	fmt.Printf("You got %d out of %d correct. \n", numberCorrect, len(quiz))
}

//parses the read all and puts into a struct for use
func parseRows(rows [][]string) []Quiz {
	quiz := make([]Quiz, len(rows))
	for i, row := range rows {
		quiz[i] = Quiz{
			Question: row[0],
			Answer:   strings.TrimSpace(row[1]),
		}
	}
	return quiz
}
