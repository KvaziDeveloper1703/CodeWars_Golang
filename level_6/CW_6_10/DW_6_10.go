/*
Write a function that takes a string as input and returns a new string where each character is replaced with "(" if it appears only once in the original string, or ")" if it appears more than once. The function should ignore capitalization when determining if a character is a duplicate.

Examples:
"din" → "((("
"recede" → "()()()"
"Success" → ")())())"
"(( @" → "))(("

Напишите функцию, которая принимает строку и возвращает новую строку, где каждый символ заменяется на "(", если он встречается в исходной строке только один раз, или на ")", если он встречается более одного раза. Функция должна игнорировать регистр при определении дубликатов.

Примеры:
"din" → "((("
"recede" → "()()()"
"Success" → ")())())"
"(( @" → "))(("

https://www.codewars.com/kata/54b42f9314d9229fd6000d9c
*/

package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(input string) string {
	input = strings.ToLower(input)
	runeCount := make(map[rune]int)

	for _, char := range input {
		runeCount[char]++
	}

	result := make([]rune, len(input))
	for i, char := range input {
		if runeCount[char] > 1 {
			result[i] = ')'
		} else {
			result[i] = '('
		}
	}

	return string(result)
}

func main() {
	fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("recede"))
	fmt.Println(DuplicateEncode("Success"))
	fmt.Println(DuplicateEncode("(( @"))
}
