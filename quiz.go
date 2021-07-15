package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Question string `csv:"Question"`
	Answer   string `csv:"Answer"`
}

func main() {
	csvFilename := flag.String("csv", "./quizzes/addition.csv",
		"A csv in the format of 'Question,Answer'. ./quizzes/subtraction.csv and ./quizzes/multiplication.csv are also available")

	timeLimit := flag.Int("Timelimit", 30, "Timelimit formatted in seconds.")
	flag.Parse()

	quizReader(*csvFilename, *timeLimit)

}

func quizReader(fileName string, timeLimit int) {
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

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	//loop through the struct to output questions
	//i+1 is so that we skip the 0 index, which is a header
	for i, problem := range quiz {

		fmt.Printf("Solve #%d: %s = \n", i+1, problem.Question)

		answerChannel := make(chan string)

		//anonymous function immediately invoked. Sets userinput to
		//its own channel so it doesn't block
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerChannel <- userAnswer
		}()

		select {

		//if you get an answer from the timer channel
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d correct. \n", numberCorrect, len(quiz))
			return

		//if you get an answer from the answer channel.
		case userAnswer := <-answerChannel:
			if userAnswer == problem.Answer {
				numberCorrect++
			}
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
