/*
Write a function that converts an integer into a string.

Напишите функцию, которая преобразует целое число в строку.

https://www.codewars.com/kata/5265326f5fda8eb1160004c8
*/

package main

import (
	"fmt"
	"strconv"
)

func NumberToString(givenNumber int) string {
	return strconv.Itoa(givenNumber)
}

func main() {
	number := 1703
	numberAsAString := NumberToString(number)
	fmt.Println(numberAsAString)
}
