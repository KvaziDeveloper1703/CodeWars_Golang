/*
Write a function that takes a string and converts it to "Jaden-Cased" format, where each word starts with a capital letter. Contractions should also follow proper capitalization, as shown in the example below.

Example:
Input: "How can mirrors be real if our eyes aren't real"
Output: "How Can Mirrors Be Real If Our Eyes Aren't Real"

Напишите функцию, которая принимает строку и преобразует её в формат "Jaden-Cased", где каждое слово начинается с заглавной буквы. Сокращения также должны быть правильно капитализированы, как показано в примере ниже.

Пример:
Ввод: "How can mirrors be real if our eyes aren't real"
Вывод: "How Can Mirrors Be Real If Our Eyes Aren't Real"

https://www.codewars.com/kata/5390bac347d09b7da40006f6
*/

package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}

func main() {
	input := "How can mirrors be real if our eyes aren't real"
	output := ToJadenCase(input)
	fmt.Println(output)
}
