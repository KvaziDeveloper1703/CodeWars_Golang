/*
Write a function that takes a string of words and returns the length of the shortest word. The input string will always contain at least one word, and all inputs will be valid strings.

Напишите функцию, которая принимает строку, содержащую слова, и возвращает длину самого короткого слова. Входная строка всегда будет содержать хотя бы одно слово, и все входные данные будут корректными строками.

https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9
*/

package main

import (
	"fmt"
	"strings"
)

func FindShortestWordLength(input string) int {
	words := strings.Fields(input)

	if len(words) == 0 {
		return 0
	}

	minimumLength := len(words[0])

	for _, word := range words {
		wordLength := len(word)
		if wordLength < minimumLength {
			minimumLength = wordLength
		}
	}

	return minimumLength
}

func main() {
	fmt.Println(FindShortestWordLength("The quick brown fox"))
	fmt.Println(FindShortestWordLength("Go is awesome"))
	fmt.Println(FindShortestWordLength("Single"))
	fmt.Println(FindShortestWordLength("This is a test"))
}
