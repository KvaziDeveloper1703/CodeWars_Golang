/*
Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.

https://www.codewars.com/kata/54ff3102c1bad923760001f3
*/

package main

import "fmt"

func countVowels(input string) (count int) {
	vowels := "aeiou"
	count = 0

	for _, char := range input {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	input := "This is a test string!"
	result := countVowels(input)
	fmt.Printf("Number of vowels in the string: %d\n", result)
}
