/*
Write a function that takes two parameters: an integer repetitions and a string value. The function should return a new string that consists of the string value repeated exactly repetitions times.

Напиш ите функцию, которая принимает два параметра: целое число n и строку s. Функция должна возвращать новую строку, состоящую из строки s, повторённой ровно n раз.

https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
*/

package main

import (
	"fmt"
	"strings"
)

func RepeatString(repetitions int, value string) string {
	if repetitions <= 0 {
		return ""
	}
	return strings.Repeat(value, repetitions)
}

func main() {
	fmt.Println(RepeatString(6, "I"))
	fmt.Println(RepeatString(5, "Hello"))
	fmt.Println(RepeatString(0, "Test"))
	fmt.Println(RepeatString(-3, "Oops"))
}
