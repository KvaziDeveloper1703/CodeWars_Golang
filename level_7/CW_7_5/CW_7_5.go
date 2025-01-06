/*
Write a function that transforms a string by repeating each character based on its position in the string. The first repetition should be capitalized, and subsequent repetitions should be lowercase. The transformed characters should be joined with hyphens. The input string will only contain letters from a to z and A to Z.

Examples:
accum("abcd") → "A-Bb-Ccc-Dddd"
accum("RqaEzty") → "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") → "C-Ww-Aaa-Tttt"

Напишите функцию, которая преобразует строку, повторяя каждый символ столько раз, сколько соответствует его позиции в строке. Первая буква в повторении должна быть заглавной, а остальные — строчными. Преобразованные символы должны быть соединены дефисами. Входная строка будет содержать только буквы от a до z и от A до Z.

Примеры:
accum("abcd") → "A-Bb-Ccc-Dddd"
accum("RqaEzty") → "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") → "C-Ww-Aaa-Tttt"

https://www.codewars.com/kata/5667e8f4e3f572a8f2000039
*/

package main

import (
	"fmt"
	"strings"
)

func Accum(given_string string) string {
	var result []string

	for i, character := range given_string {
		repeated_string := strings.Repeat(strings.ToLower(string(character)), i+1)
		transformed_string := strings.ToUpper(string(repeated_string[0])) + repeated_string[1:]
		result = append(result, transformed_string)
	}

	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("abcd"))
	fmt.Println(Accum("RqaEzty"))
	fmt.Println(Accum("cwAt"))
}
