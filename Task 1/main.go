package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, err := input.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
	} else {
		val := wordFreq(str)
		fmt.Println("Word frequencies:", val)
	}
}

func wordFreq(inputedWord string) map[string]int {
	dict := make(map[string]int)

	// FieldsFunc splits the string based on any character that is NOT a letter or number
	words := strings.FieldsFunc(inputedWord, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		// Normalize to lowercase for case-insensitive counting
		lowerWord := strings.ToLower(word)
		dict[lowerWord]++
	}

	return dict
}
