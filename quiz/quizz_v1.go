package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	questions := extract_questions(file)

	user_score := 0
	var user_ans int
	for i, question := range questions {
		fmt.Printf("Problem #%02d: %v = ", (i + 1), question.title)
		fmt.Scanf("%d ", &user_ans)

		if ans, _ := strconv.Atoi(question.answer); ans == user_ans {
			user_score++
		}
	}

	fmt.Printf("You scored %d of %d.\n", user_score, len(questions))
	fmt.Printf("The time limit of the quiz is %d.\n", *limit)
}
