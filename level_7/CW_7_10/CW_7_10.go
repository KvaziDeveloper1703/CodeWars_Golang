/*
Write a function that takes two strings, s1 and s2, containing only lowercase letters from a to z. The function should return a new string that:
+ Is sorted in alphabetical order.
+ Contains the longest possible combination of distinct letters, each appearing only once.
+ Includes letters from both s1 and s2.

Examples:
Input:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
Output: "abcdefklmopqwxy"

Input:
a = "abcdefghijklmnopqrstuvwxyz"
b = "abcdefghijklmnopqrstuvwxyz"
Output: "abcdefghijklmnopqrstuvwxyz"

Напишите функцию, которая принимает две строки s1 и s2, содержащие только строчные буквы от a до z. Функция должна вернуть новую строку, которая:
+ Отсортирована в алфавитном порядке.
+ Содержит максимально возможное количество различных букв, каждая из которых встречается только один раз.
+ Включает буквы из обеих строк s1 и s2.

Примеры:
Ввод:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
Вывод: "abcdefklmopqwxy"

Ввод:
a = "abcdefghijklmnopqrstuvwxyz"
b = "abcdefghijklmnopqrstuvwxyz"
Вывод: "abcdefghijklmnopqrstuvwxyz"

https://www.codewars.com/kata/5656b6906de340bd1b0000ac
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func LongestDistinctCombination(givenString1, givenString2 string) string {
	combinedString := givenString1 + givenString2
	uniqueLetters := make(map[rune]bool)

	for _, char := range combinedString {
		uniqueLetters[char] = true
	}

	letters := make([]string, 0, len(uniqueLetters))
	for char := range uniqueLetters {
		letters = append(letters, string(char))
	}

	sort.Strings(letters)

	return strings.Join(letters, "")
}

func main() {
	s1 := "xyaabbbccccdefww"
	s2 := "xxxxyyyyabklmopq"
	fmt.Println(LongestDistinctCombination(s1, s2))

	s1 = "hello"
	s2 = "world"
	fmt.Println(LongestDistinctCombination(s1, s2))
}
