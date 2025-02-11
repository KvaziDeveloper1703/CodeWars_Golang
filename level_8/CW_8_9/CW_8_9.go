/*
Write a function that takes two parameters: an integer n and a string s. The function should return a new string where the input string s is repeated exactly n times.

Напишите функцию, которая принимает два параметра: целое число n и строку s. Функция должна вернуть новую строку, в которой строка s повторяется ровно n раз.

https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
*/

package main

import (
	"fmt"
	"strings"
)

func RepeatString(N int, givenString string) string {
	if N <= 0 {
		return ""
	}
	return strings.Repeat(givenString, N)
}

func main() {
	N := 5
	word := "Hello"
	result := RepeatString(N, word)
	fmt.Println(result)
}
