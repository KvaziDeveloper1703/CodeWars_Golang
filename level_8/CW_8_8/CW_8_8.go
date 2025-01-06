/*
Write a function that takes a string as input and removes its first and last characters. The input will always be a string with at least two characters.

Напишите функцию, которая принимает строку и удаляет её первый и последний символы. Входная строка всегда содержит как минимум два символа.

https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0
*/

package main

import "fmt"

func RemoveFirstAndLast(word string) string {
	if len(word) <= 2 {
		return ""
	}
	return word[1 : len(word)-1]
}

func main() {
	fmt.Println(RemoveFirstAndLast("hello"))
	fmt.Println(RemoveFirstAndLast("world"))
	fmt.Println(RemoveFirstAndLast("ab"))
	fmt.Println(RemoveFirstAndLast("a"))
}
