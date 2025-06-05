package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/questions"
	"strings"
)

func Run(questions []questions.Question) (correctAnswers uint) {
	fmt.Println("Welcome to Quiz Game!")

	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}

		continue
	}

	return correctAnswers
}

func askQuestion(question questions.Question) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\nEnter the capital of %s\n", question.Country)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading the entered text! Please try again")
			continue
		}

		return strings.TrimRight(strings.ToLower(input), "\r\n") == strings.ToLower(question.Capital)
	}
}
