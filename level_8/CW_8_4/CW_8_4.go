/*
Write a program that takes a string as input and returns the string reversed.

For example:
Input: 'world' → Output: 'dlrow'
Input: 'word' → Output: 'drow'

Напишите функцию, которая принимает строку и возвращает её в перевёрнутом виде.

Например:
Ввод: 'world' → Вывод: 'dlrow'
Ввод: 'word' → Вывод: 'drow'

https://www.codewars.com/kata/5168bb5dfe9a00b126000018
*/

package main

import "fmt"

func Solution(word string) string {
	runes := []rune(word)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(Solution("world"))
	fmt.Println(Solution("word"))
	fmt.Println(Solution("a"))
	fmt.Println(Solution(""))
}
