package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	WordList   []string
	ChosenWord string
	Green      string = "🟩"
	Yellow     string = "🟨"
	Black      string = "⬛"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	WordList, _ = readWordList("./wordlist.txt")
}

func main() {
	ChosenWord = PickAWord()
	fmt.Println("Chosen word: ", ChosenWord)

	var guess string
	for guess != ChosenWord {
		fmt.Print("Enter Guess: ")
		fmt.Scanln(&guess)
		fmt.Println(IsItRight(strings.ToLower(guess)))
	}

	fmt.Println("You Win!")
}

func readWordList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func PickAWord() string {
	return WordList[rand.Intn(len(WordList))]
}

func IsItRight(guess string) (feedback string) {
	if len(guess) != len(ChosenWord) {
		return "Guess length does not match."
	}

	for i, c := range guess {
		switch {
		case byte(c) == ChosenWord[i]:
			feedback += Green
			break
		case strings.ContainsRune(ChosenWord, c):
			feedback += Yellow
			break
		default:
			feedback += Black
			break
		}
	}

	return feedback
}
