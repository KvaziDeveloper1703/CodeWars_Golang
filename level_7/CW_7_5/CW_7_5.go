/*
Write a function that transforms a string by repeating each character based on its position in the string. The first repetition should be capitalized, and subsequent repetitions should be lowercase. The transformed characters should be joined with hyphens. The input string will only contain letters from a to z and A to Z.

Examples:
accum("abcd") → "A-Bb-Ccc-Dddd"
accum("RqaEzty") → "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"

Напишите функцию, которая преобразует строку, повторяя каждый символ столько раз, сколько соответствует его позиции в строке. Первая буква в повторении должна быть заглавной, а остальные — строчными. Преобразованные символы должны быть соединены дефисами. Входная строка будет содержать только буквы от a до z и от A до Z.

Примеры:
accum("abcd") → "A-Bb-Ccc-Dddd"
accum("RqaEzty") → "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"

https://www.codewars.com/kata/5667e8f4e3f572a8f2000039
*/

package main

import (
	"fmt"
	"strings"
)

func Accum(givenString string) string {
	var result []string

	for i, character := range givenString {
		characterAsAString := string(character)
		lowerCharacter := strings.ToLower(characterAsAString)
		repeatedString := strings.Repeat(lowerCharacter, i+1)
		firstUpperCharacter := strings.ToUpper(string(repeatedString[0]))
		transformedString := firstUpperCharacter + repeatedString[1:]
		result = append(result, transformedString)
	}

	return strings.Join(result, "-")
}

func main() {
	myString := "RqaEzty"
	result := Accum(myString)
	fmt.Println(result)
}
