/*
Write a function that takes a string as input and removes its first and last characters. The input will always be a string with at least two characters.

Напишите функцию, которая принимает строку и удаляет её первый и последний символы. Входная строка всегда содержит как минимум два символа.

https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0
*/

package main

import "fmt"

func RemoveFirstAndLast(givenWord string) string {
	if len(givenWord) <= 2 {
		return ""
	}
	return givenWord[1 : len(givenWord)-1]
}

func main() {
	word := "word"
	result := RemoveFirstAndLast(word)
	fmt.Println(result)
}
