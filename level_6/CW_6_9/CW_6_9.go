/*
Write a function that takes a string as input and returns the count of distinct case-insensitive alphabetic characters and numeric digits that appear more than once in the string. The input string will contain only letters and numeric digits.

Examples:
"abcde" → 0 (no characters repeat more than once)
"aabbcde" → 2 ('a' and 'b')
"aabBcde" → 2 ('a' occurs twice and 'b' twice, case-insensitive)

Напишите функцию, которая принимает строку и возвращает количество различных символов (букв и цифр), которые встречаются более одного раза, не учитывая регистр. Входная строка будет содержать только буквы и цифры.

Примеры:
"abcde" → 0 (нет символов, которые повторяются более одного раза)
"aabbcde" → 2 ('a' и 'b')
"aabBcde" → 2 ('a' встречается дважды, 'b' — дважды, регистр не учитывается)

https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1
*/

package main

import (
	"fmt"
	"strings"
)

func CountDuplicates(input string) int {
	countMap := make(map[rune]int)

	lowerInput := strings.ToLower(input)

	for _, char := range lowerInput {
		countMap[char]++
	}

	duplicateCount := 0
	for _, count := range countMap {
		if count > 1 {
			duplicateCount++
		}
	}

	return duplicateCount
}

func main() {
	fmt.Println(CountDuplicates("abcde"))
	fmt.Println(CountDuplicates("aabbcde"))
	fmt.Println(CountDuplicates("aabBcde"))
	fmt.Println(CountDuplicates("indivisibilities"))
}
