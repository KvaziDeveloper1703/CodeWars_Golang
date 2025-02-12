/*
Write a function that takes a string as input and returns a new string with all vowels removed. For this task, vowels are defined as a, e, i, o, u and y.

Напишите функцию, которая принимает строку и возвращает новую строку с удалёнными гласными. В рамках задачи гласными считаются a, e, i, o, u и y.

https://www.codewars.com/kata/52fba66badcd10859f00097e
*/

package main

import (
	"fmt"
)

func removeVowels(givenString string) string {
	vowels := "aeiouyAEIOUY"
	result := ""

	for _, character := range givenString {
		isVowel := false
		for _, vowel := range vowels {
			if character == vowel {
				isVowel = true
				break
			}
		}
		if !isVowel {
			result += string(character)
		}
	}

	return result
}

func main() {
	myString := "This website is for losers LOL!"
	result := removeVowels(myString)
	fmt.Println(result)
}
