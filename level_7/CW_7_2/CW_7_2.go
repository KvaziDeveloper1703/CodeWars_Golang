/*
Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.

https://www.codewars.com/kata/52fba66badcd10859f00097e
*/

package main

import (
	"fmt"
)

func removeVowels(input string) string {
	vowels := "aeiouAEIOU"
	result := ""

	for _, char := range input {
		isVowel := false
		for _, vowel := range vowels {
			if char == vowel {
				isVowel = true
				break
			}
		}
		if !isVowel {
			result += string(char)
		}
	}

	return result
}

func main() {
	comment := "This website is for losers LOL!"
	result := removeVowels(comment)
	fmt.Println(result)
}
