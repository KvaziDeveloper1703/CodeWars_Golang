/*
Write a function that takes a string as input and returns a new string with all vowels removed. For this task, vowels are defined as a, e, i, o, u and y.

For example:
Input: "This website is for losers LOL!"
Output: "Ths wbst s fr lsrs LL!"

Напишите функцию, которая принимает строку и возвращает новую строку с удалёнными гласными. В рамках задачи гласными считаются a, e, i, o, u и y.

Например:
Ввод: "This website is for losers LOL!"
Вывод: "Ths wbst s fr lsrs LL!"

https://www.codewars.com/kata/52fba66badcd10859f00097e
*/

package main

import (
	"fmt"
)

func removeVowels(input string) string {
	vowels := "aeiouyAEIOUY"
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
