/*
Write a function that takes a non-empty string as input and returns its middle character(s):
+ If the string has an odd length, return the single middle character.
+ If the string has an even length, return the two middle characters.

Examples:
Input: "test" → Output: "es"
Input: "testing" → Output: "t"
Input: "middle" → Output: "dd"
Input: "A" → Output: "A"

Напишите функцию, которая принимает непустую строку и возвращает её средний символ или символы:
+ Если длина строки нечётная, вернуть один средний символ.
+ Если длина строки чётная, вернуть два средних символа.

Примеры:
Ввод: "test" → Вывод: "es"
Ввод: "testing" → Вывод: "t"
Ввод: "middle" → Вывод: "dd"
Ввод: "A" → Вывод: "A"

https://www.codewars.com/kata/56747fd5cb988479af000028
*/

package main

import "fmt"

func GetMiddleCharacter(input string) string {
	length := len(input)
	middle := length / 2

	if length%2 == 0 {
		return input[middle-1 : middle+1]
	}

	return string(input[middle])
}

func main() {
	fmt.Println(GetMiddleCharacter("test"))
	fmt.Println(GetMiddleCharacter("testing"))
	fmt.Println(GetMiddleCharacter("middle"))
	fmt.Println(GetMiddleCharacter("A"))
}
