package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readInput reads a string from the user.
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		line, err := reader.ReadString('\n')
		sentence := strings.TrimSpace(line)
		if err != nil || sentence == "" {
			fmt.Println("Please input a valid string.")
		} else {
			return sentence
		}
	}
}

// checkPalindrome checks if a given string is a palindrome.
func checkPalindrome(sentence string) bool {
	left, right := 0, len(sentence)-1
	for left < right {
		if sentence[left] != sentence[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// countWordFrequency counts the frequency of each word in the given sentence.
func countWordFrequency(sentence string) map[string]int {
	wordFrequency := make(map[string]int)
	for _, word := range strings.Fields(sentence) {
		wordFrequency[word]++
	}
	return wordFrequency
}

// printWordFrequency prints the word frequency map.
func printWordFrequency(wordFrequency map[string]int) {
	fmt.Println("Word Frequency: ")
	for word, frequency := range wordFrequency {
		fmt.Printf("%s: %d\n", word, frequency)
	}
}

// checkPalindromeOption handles the palindrome checking option.
func checkPalindromeOption() {
	sentence := readInput("Input a string to check for palindrome: ")
	if checkPalindrome(sentence) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}

// countWordFrequencyOption handles the word frequency counting option.
func countWordFrequencyOption() {
	sentence := readInput("Input a sentence to count word frequency: ")
	wordFrequency := countWordFrequency(sentence)
	printWordFrequency(wordFrequency)
	fmt.Println()
}

// exitOption handles the exit option.
func exitOption() {
	fmt.Println("Exiting program.")
}

// main function handles user choices and calls the appropriate functions.
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Check if a string is a palindrome")
		fmt.Println("2. Count word frequency in a sentence")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1/2/3): ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			checkPalindromeOption()
		case "2":
			countWordFrequencyOption()
		case "3":
			exitOption()
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
