package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	//"slices"

	"github.com/fatih/color"
)

//func check_guess(guess string, start_word string) bool {
//	var guessRegexp []string

//	for i, r := range guess{
//
//	}
//return slices.Contains(start_word, guess)
//}

func select_start_word(words []string) string {
	// TODO word checker
	var word string
	var accept string
	accept_word := false
	for !accept_word {
		wordIndex := rand.Intn(len(words))
		word = words[wordIndex]
		fmt.Printf("Is %s acceptable? ", word)
		fmt.Scanln(&accept)
		accept_word = (accept == "y")
	}
	return word
}

func main() {
	var words []string
	file, err := os.Open("./word-list.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		word := fscanner.Text()
		words = append(words, word)
	}

	//start_word := select_start_word(words)
	correctGuess := false
	c := color.New(color.FgHiBlack).Add(color.BgGreen).Add(color.Bold)
	c.Println("Enter your guess using the following chars")
	c.Println("= for letter and location match")
	c.Println("? for letter match")
	c.Println(". for no  match")
	for !correctGuess {
		fmt.Print("What is the result of your guess: ")
		var guess string
		fmt.Scanln(&guess)
		//correctGuess = check_guess(guess, start_word)
	}
	fmt.Println("You guessed it")
}
