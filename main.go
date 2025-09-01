package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have some chances based on the difficult to guess the correct number.")

	play := "y"

PLAY:
	for strings.ToLower(play) == "y" {
		fmt.Println("\nPlease select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")

		fmt.Print("\nEnter your choice: ")
		var choice int16
		maxAttempt := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Great! You have selected the Easy difficulty level.")
			maxAttempt = 10
		case 2:
			fmt.Println("Great! You have selected the Medium difficulty level.")
			maxAttempt = 5
		case 3:
			fmt.Println("Great! You have selected the Hard difficulty level.")
			maxAttempt = 3
		default:
			log.Fatal("Unknown difficulty. The difficult range are 1 to 3 (easy to hard)")
		}
		fmt.Println("Let's start the game!")

		target := rand.Intn(101)
		if target == 0 {
			target = 100
		}

		var input int
		attempt := 1
		start := time.Now()
		for attempt < maxAttempt {
			fmt.Print("\nEnter your guess: ")
			fmt.Scan(&input)

			if input > target {
				fmt.Printf("Incorrect! The number is less than %d.\n", input)
			} else if input < target {
				fmt.Printf("Incorrect! The number is greater than %d.\n", input)
			} else {
				duration := time.Since(start)
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts and with time: %v\n",
					attempt,
					duration,
				)

				fmt.Print("Do you want to play again? [y/n] ")
				fmt.Scan(&play)
				continue PLAY
			}

			attempt++
		}

		fmt.Printf("Game Over! You exceeeded max attempts. The correct number is %d\n", target)
		fmt.Print("Do you want to play again? [y/n] ")
		fmt.Scan(&play)
	}
}
