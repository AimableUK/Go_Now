package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	str, err := input.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		val := wordFreq(str)
		fmt.Println(val)
	}
}

func wordFreq(inputedWord string) {

	x := ""
	for char := range inputedWord {
		if char == " " {

		}
		x += char
	}
}
