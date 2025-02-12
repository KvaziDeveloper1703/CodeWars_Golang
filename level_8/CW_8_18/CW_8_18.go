/*
Given a random non-negative number, return its digits as an array in reverse order.

Дано случайное неотрицательное число. Необходимо вернуть его цифры в виде массива в обратном порядке.

https://www.codewars.com/kata/5583090cbe83f4fd8c000051
*/

package main

import (
	"fmt"
	"strconv"
)

func Digitize(givenNumber int) []int {
	givenNumberAsAString := strconv.Itoa(givenNumber)
	result := make([]int, len(givenNumberAsAString))
	for i, character := range givenNumberAsAString {
		result[len(givenNumberAsAString)-1-i] = int(character - '0')
	}
	return result
}

func main() {
	number := 35231
	digitizedNumber := Digitize(number)
	fmt.Println(digitizedNumber)
}
