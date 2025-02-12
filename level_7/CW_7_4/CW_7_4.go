/*
Write a function that takes a non-empty string as input and returns its middle character(s):
+ If the string has an odd length, return the single middle character.
+ If the string has an even length, return the two middle characters.

Напишите функцию, которая принимает непустую строку и возвращает её средний символ или символы:
+ Если длина строки нечётная, вернуть один средний символ.
+ Если длина строки чётная, вернуть два средних символа.

https://www.codewars.com/kata/56747fd5cb988479af000028
*/

package main

import "fmt"

func GetMiddleCharacter(givenString string) string {
	length := len(givenString)
	middle := length / 2

	if length%2 == 0 {
		return givenString[middle-1 : middle+1]
	}

	return string(givenString[middle])
}

func main() {
	myString := "test"
	result := GetMiddleCharacter(myString)
	fmt.Println(result)
}
