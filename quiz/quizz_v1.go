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

func main() {
	file_name := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	questions := map[string]string{}

	// READ CSV FILE
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

		questions[line[0]] = line[1]
	}

	i := 1
	score := 0
	var user_ans int
	for question, answer := range questions {
		fmt.Printf("Problem #%02d: %s = ", i, question)
		fmt.Scanf("%d ", &user_ans)

		if ans, _ := strconv.Atoi(answer); ans == user_ans {
			score++
		}
		i++
	}

	fmt.Printf("You scored %d of %d.\n", score, len(questions))
	fmt.Printf("The time limit of the quiz is %d.\n", *limit)
}
