package main

import (
	"context"
	"fmt"

	"github.com/kramen22/gordle/pkg/dictionary"
	"github.com/kramen22/gordle/pkg/state"
)

var exitcodes = map[string]string{
	"q":    "q",
	"Q":    "Q",
	"quit": "quite",
	"exit": "exit",
}

func main() {
	ctx := context.Background()

	dict, err := dictionary.New(ctx)
	if err != nil {
		panic(err.Error())
	}

	state := state.New(dict)
	state.StartGame()

	var input string

	fmt.Printf("Welcome to gordle! Type h for help, q to quit (when i implement that)\n\n\n")

	for {
		fmt.Print(state.GetBoardPrompt())
		fmt.Print("Enter your guess: ")

		fmt.Scanln(&input)
		switch input {
		case exitcodes[input]:
			return
		case "cheat":
			fmt.Printf("the answer is: %s\n", state.Target)
		default:
			if reason, ok := state.IsValidGuess(input); !ok {
				fmt.Println(reason)
			} else {
				won := state.GuessWord(input)
				if won {
					fmt.Println("You won!!")
				}
			}
		}
	}
}
