/*
Complete the function that takes a string as input and reverses each word individually, while keeping the original spaces intact.

Rules:
+ Each word in the string should be reversed, but their positions remain unchanged.
+ Any extra spaces between words must be retained in the output.

Реализуйте функцию, которая принимает строку и переворачивает каждое слово отдельно, сохраняя все пробелы в исходном порядке.

Правила:
+ Слова должны быть перевёрнуты, но их порядок не изменяется.
+ Дополнительные пробелы между словами должны быть сохранены.

https://www.codewars.com/kata/5259b20d6021e9e14c0010d4
*/

package main

import (
	"fmt"
	"strings"
)

func ReverseWord(givenWord string) string {
	runes := []rune(givenWord)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseWords(givenString string) string {
	words := strings.Split(givenString, " ")
	for i, word := range words {
		words[i] = ReverseWord(word)
	}
	return strings.Join(words, " ")
}

func main() {
	myString := "This is an example!"
	result := ReverseWords(myString)
	fmt.Println(result)
}
