package main

import "fmt"

var GenreTable = map[string]Searcheble {
	"Super-Hero": superheroMovie{},
	"Cult": cultMovie{},
	"Default": defaultMovie{},
	"Terror": TerrorMovie{},
}

func Run() {
	userInput := RequestUserInputAsString("Gander: ")

	if(!IsGenreValid(userInput)) {
		fmt.Println("Invalid Genre")
		return
	}

	rating, err := GenreTable[userInput].CollectRatings()

	if(err != nil) {
		fmt.Println("Something was wrong")
		return
	}

	fmt.Println("Rating: %f Error: %v", rating, err)
}

func RequestUserInputAsString(message string) (string) {
	var userInput string
	fmt.Println(message)
	fmt.Scan(&userInput)
	return userInput
}

func IsGenreValid(Genre string) (bool) {
	return true
}