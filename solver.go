package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	letters "wordle/solver/internal"

	//"slices"

	"github.com/fatih/color"
)

func check_guess(guess string, start_word string, letterList [5]letters.Letter) [5]letters.Letter {
	fmt.Println(start_word)
	var sb strings.Builder
	sb.WriteString("^")
	for i, r := range guess {
		switch r {
		case '=':
			fmt.Printf("Exact Match %c \n", start_word[i])
			letterList[i].IsExact = true
			letterList[i].ExactLetter = string(start_word[i])
		case '?':
			fmt.Printf("Letter Match %c \n", start_word[i])
			letterList[i].ThisLetter = append(letterList[i].ThisLetter, string(start_word[i]))
			for letterIdx, _ := range letterList {
				if letterIdx == i {
					continue
				}
				letterList[letterIdx].LetterGuess = append(letterList[letterIdx].LetterGuess, string(start_word[i]))
				fmt.Println(letterList[letterIdx].LetterGuess)
			}
		default:
			fmt.Printf("No Match %c \n", start_word[i])
			for letterIdx, _ := range letterList {
				letterList[letterIdx].LetterGuess = append(letterList[letterIdx].LetterGuess, string(start_word[i]))
				fmt.Println(letterIdx)
				fmt.Println(letterList[letterIdx].LetterGuess)
			}

		}
	}
	return letterList
}

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

	letterList := [5]letters.Letter{}

	start_word := select_start_word(words)
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
		letterList = check_guess(guess, start_word, letterList)
		// step over letterList to create regexp string
		var testString strings.Builder
		correctGuess = true
		for _, letter := range letterList {
			testString.WriteString(string(letter.MakeRegexString()))
			correctGuess = correctGuess && letter.IsExact
		}
		fmt.Println(testString.String())
	}
	fmt.Println("You guessed it")
}
