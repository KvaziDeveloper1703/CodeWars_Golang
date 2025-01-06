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

func RepeatString(n int, s string) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}

func main() {
	fmt.Println(RepeatString(6, "I"))
	fmt.Println(RepeatString(5, "Hello"))
	fmt.Println(RepeatString(0, "Test"))
	fmt.Println(RepeatString(-3, "Oops"))
}
