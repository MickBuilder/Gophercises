package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Question struct {
	title  string
	answer string
}

func extract_questions(file *os.File) []Question {
	results := make([]Question, 0)
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		question := Question{line[0], line[1]}
		results = append(results, question)
	}
	return results
}

func main() {
	file_name := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("You have %d seconds to solve the problems\n", *limit)

	questions := extract_questions(file)

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	user_score := 0

loop:
	for i, question := range questions {
		fmt.Printf("Problem #%02d: %v = ", (i + 1), question.title)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case answer := <-answerCh:
			if answer == question.answer {
				user_score++
			}
		}
	}

	fmt.Printf("You scored %d of %d. [Motherfucker]\n", user_score, len(questions))
}
