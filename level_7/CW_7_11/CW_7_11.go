/*
Complete the function so that it returns true if the first argument ends with the second argument. Otherwise, return false.

Examples:
solution("abc", "bc") → true

Доработай функцию так, чтобы она возвращала true, если первая строка заканчивается на вторую строку. В противном случае функция должна вернуть false.

Пример:
solution("abc", "bc") → true

https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d
*/

package main

import (
	"fmt"
)

func Solution(givenString, givenEnding string) bool {
	stringLength, endingLength := len(givenString), len(givenEnding)
	if endingLength > stringLength {
		return false
	}
	return givenString[stringLength-endingLength:] == givenEnding
}

func main() {
	myString := "abc"
	ending := "bc"
	result := Solution(myString, ending)
	fmt.Println(result)
}
