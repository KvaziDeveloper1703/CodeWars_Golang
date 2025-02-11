/*
Write a function that takes a boolean value as input and returns the string "Yes" if the value is true, or "No" if the value is false.

Напишите функцию, которая принимает логическое значение и возвращает строку "Yes", если значение равно true, или "No", если значение равно false.

https://www.codewars.com/kata/53369039d7ab3ac506000467
*/

package main

import "fmt"

func BoolToString(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

func main() {
	boolValue := true
	stringValue := BoolToString(boolValue)
	fmt.Println(stringValue)
}
