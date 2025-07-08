package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

var (
	digits    = "0123456789"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	var passwordStorage []string
	for {
		charSet, maxLength := getCharSet()
		length := getLength(maxLength)

		password := generateUniquePassword(charSet, length, passwordStorage)
		passwordStorage = append(passwordStorage, password)

		fmt.Printf("Generated password: %s\n", password)
		fmt.Println()
		fmt.Println("Password storage:")
		for i, pwd := range passwordStorage {
			fmt.Printf("Password %d: %s\n", i+1, pwd)
		}
		fmt.Println()
		if !askYesNo("Generate another password with new parameters? (y/n): ") {
			break
		}
	}
}

func getCharSet() ([]string, int) {
	var maxLength int
	for {
		var charSet []string
		if askYesNo("Use digits (0-9)? (y/n): ") {
			charSet = append(charSet, digits)
			maxLength += len(digits)
		}
		if askYesNo("Use lowercase letters (a-z)? (y/n): ") {
			charSet = append(charSet, lowercase)
			maxLength += len(lowercase)
		}
		if askYesNo("Use uppercase letters (A-Z)? (y/n): ") {
			charSet = append(charSet, uppercase)
			maxLength += len(uppercase)
		}
		if len(charSet) > 0 {
			return charSet, maxLength
		}
		fmt.Println("You must select at least one character set. Please try again.")
	}
}

func askYesNo(prompt string) bool {
	for {
		fmt.Print(prompt)
		var response string
		fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}

func getLength(maxLength int) int {
	for {
		var length int
		fmt.Printf("Enter the desired password length (min 3, max %d): ", maxLength)
		fmt.Scanln(&length)
		if length >= 3 && length <= maxLength {
			return length
		}
		fmt.Printf("Invalid length. Please enter a number between 3 and %d.\n", maxLength)
	}
}

func generateUniquePassword(charSet []string, length int, passwordStorage []string) string {
	var allChars string
	for _, set := range charSet {
		allChars += set
	}
	for {
		passwordRunes := []rune{}
		usedChars := map[rune]bool{}

		for _, set := range charSet {
			ch := randomUniqueChar(set, usedChars)
			passwordRunes = append(passwordRunes, ch)
			usedChars[ch] = true
		}
		for len(passwordRunes) < length {
			ch := randomUniqueChar(allChars, usedChars)
			passwordRunes = append(passwordRunes, ch)
			usedChars[ch] = true
		}

		r.Shuffle(len(passwordRunes), func(i, j int) {
			passwordRunes[i], passwordRunes[j] = passwordRunes[j], passwordRunes[i]
		})

		password := string(passwordRunes)
		if !slices.Contains(passwordStorage, password) {
			return password
		}
	}
}

func randomUniqueChar(set string, usedChars map[rune]bool) rune {
	for {
		ch := rune(set[r.Intn(len(set))])
		if !usedChars[ch] {
			return ch
		}
	}
}
