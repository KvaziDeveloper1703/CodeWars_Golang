/*
Write a program that takes a string as input and returns the string reversed.

Напишите функцию, которая принимает строку и возвращает её в перевёрнутом виде.

https://www.codewars.com/kata/5168bb5dfe9a00b126000018
*/

package main

import "fmt"

func Solution(givenWord string) string {
	runes := []rune(givenWord)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	word := "word"
	result := Solution(word)
	fmt.Println(result)
}
