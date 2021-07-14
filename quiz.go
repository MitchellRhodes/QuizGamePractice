package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	Question string `csv:"Question"`
	Answer   string `csv:"Answer"`
}

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
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//attributes the newly made struct to a variable
	quiz := parseRows(rows)

	//loop through the struct to output questions
	//i+1 is so that we skip the 0 index, which is a header
	for i, problem := range quiz {
		fmt.Printf("Solve #%d: %s = \n", i+1, problem.Question)

		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == problem.Answer {
			fmt.Println("Good Job")
		} else {
			fmt.Println("Wrong")
		}
	}
}

//parses the read all and puts into a struct for use
func parseRows(rows [][]string) []Quiz {
	quiz := make([]Quiz, len(rows))
	for i, row := range rows {
		quiz[i] = Quiz{
			Question: row[0],
			Answer:   row[1],
		}
	}
	return quiz
}
