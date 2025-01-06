/*
Write a function that takes a string containing one or more words as input and returns the same string, but with all words that are five or more letters long reversed. The input string will consist only of letters and spaces. Spaces will be present only when the input contains multiple words.

Напишите функцию, которая принимает строку, содержащую одно или несколько слов, и возвращает эту же строку, но со словами длиной пять и более букв, записанными в обратном порядке. Входная строка будет содержать только буквы и пробелы. Пробелы будут присутствовать только в случае, если строка содержит несколько слов.

https://www.codewars.com/kata/5264d2b162488dc400000001
*/

package main

import (
	"fmt"
	"strings"
)

func ReverseLongWords(input string) string {
	words := strings.Fields(input)

	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverseString(word)
		}
	}

	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(ReverseLongWords("Hey fellow warriors"))
	fmt.Println(ReverseLongWords("This is a test"))
	fmt.Println(ReverseLongWords("Hello world"))
}
