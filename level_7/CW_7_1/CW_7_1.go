/*
Write a program that takes a string as input and returns the count of vowels in the string. For this task, vowels are defined as a, e, i, o, u, y. The input string will consist only of lowercase letters and spaces.

Напишите функцию, которая принимает строку и возвращает количество гласных в этой строке. В рамках задачи гласными считаются a, e, i, o, u, y. Входная строка будет содержать только строчные буквы и пробелы.

https://www.codewars.com/kata/54ff3102c1bad923760001f3
*/

package main

import "fmt"

func countVowels(givenString string) (count int) {
	vowels := "aeiouy"
	count = 0

	for _, character := range givenString {
		for _, vowel := range vowels {
			if character == vowel {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	myString := "This is a test string!"
	result := countVowels(myString)
	fmt.Printf("Number of vowels in the string: %d\n", result)
}
