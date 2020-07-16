package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min) + min

	// uncomment to see the random  during the game
	fmt.Println("Here is the random number:", randomNumber)

	fmt.Println("Random Number Guessing Game in Go")
	fmt.Println("Enter a number between 1 and 100")

	guesses := 0
	for {
		guesses++
		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Oops! something went wrong, please try again", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)

		guess, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid entry, please enter a number. Error->", err)
			continue
		}

		if guess > randomNumber {
			fmt.Println(guess, "is too high, try again")
		} else if guess < randomNumber {
			fmt.Println(guess, "is too low,try again")
		} else {
			fmt.Println("Congrats!!! You got the random number after", guesses, "guesses")
			fmt.Println("Game Over!")
			break
		}
	}
}
