package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	letters "wordle/solver/internal"

	"github.com/fatih/color"
)

func checkGuess(guess string, start_word string, letterList [5]letters.Letter) [5]letters.Letter {
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
			for letterIdx := range letterList {
				if letterIdx == i {
					letterList[letterIdx].ThisLetter = append(letterList[letterIdx].ThisLetter, string(start_word[i]))
					letterList[letterIdx].LetterGuess = append(letterList[letterIdx].LetterGuess, string(start_word[i]))
				}
			}
		default:
			fmt.Printf("No Match %c \n", start_word[i])
			for letterIdx := range letterList {
				letterList[letterIdx].LetterGuess = append(letterList[letterIdx].LetterGuess, string(start_word[i]))
				// fmt.Println(letterIdx)
				// fmt.Println(letterList[letterIdx].LetterGuess)
			}

		}
	}
	return letterList
}

func selectStartWord(words []string) string {
	// TODO word checker
	var word string
	var accept string
	acceptWord := false
	for !acceptWord {
		wordIndex := rand.Intn(len(words))
		word = words[wordIndex]
		fmt.Printf("Is %s acceptable? ", word)
		fmt.Scanln(&accept)
		acceptWord = (accept == "y")
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

	correctGuess := false
	for !correctGuess {
		startWord := selectStartWord(words)
		c := color.New(color.FgHiBlack).Add(color.BgGreen).Add(color.Bold)
		c.Println("Enter your guess using the following chars")
		c.Println("= for letter and location match")
		c.Println("? for letter match")
		c.Println(". for no  match")
		fmt.Print("What is the result of your guess: ")
		var guess string
		fmt.Scanln(&guess)
		letterList = checkGuess(guess, startWord, letterList)
		// step over letterList to create regexp string
		var testString strings.Builder
		correctGuess = true
		for _, letter := range letterList {
			testString.WriteString(letter.MakeRegexString())
			correctGuess = correctGuess && letter.IsExact
		}
		var newWords []string
		for _, word := range words {
			matched, err := regexp.MatchString(testString.String(), word)
			if err != nil {
				fmt.Println("Error: ", err)
			} else if matched {
				addString := true
				for idx := range letterList {
					for letterIdx := range letterList[idx].ThisLetter {
						addString = addString && strings.Contains(word, letterList[idx].ThisLetter[letterIdx])
					}
				}
				if addString {
					newWords = append(newWords, word)
				}
			} else {
				continue
			}
		}
		fmt.Println(testString.String())
		fmt.Println("There are ", len(newWords), " words left")
		if len(newWords) == 1 {
			correctGuess = true
			fmt.Println("Your word has to be ", newWords[0])
		} else if len(newWords) == 0 {
			fmt.Println("There is something wrong, there are no words left")
			os.Exit(1)
		} else {
			words = newWords
		}
	}
	fmt.Println("You guessed it")
}
