package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxGuesses = 10
const MaxNumber = 100

func main() {
	guess_count := 0
	rand.Seed(time.Now().UnixNano())
	correct_answer := rand.Intn(MaxNumber)

	for guess_count < MaxGuesses {
		fmt.Printf("Guess a number (0-%v): ", MaxNumber)
		current_guess := 0
		_, err := fmt.Scanln(&current_guess)

		if err != nil {
			fmt.Println(err)
			return
		}

		current_guess = int(current_guess)

		if current_guess == correct_answer {
			fmt.Printf("Correct! You got it in %v guess(es).\n", guess_count)
			fmt.Println()
			return
		} else if current_guess < correct_answer {
			fmt.Println("Too low.")
		} else {
			fmt.Println("Too high.")
		}

		guess_count++
		fmt.Println()
	}

	fmt.Printf("You lose. The correct answer was: %v", correct_answer)
	fmt.Println()
}
