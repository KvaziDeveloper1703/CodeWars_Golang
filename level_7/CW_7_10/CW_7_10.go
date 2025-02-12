/*
Write a function that takes two strings, string1 and string2, containing only lowercase letters from a to z. The function should return a new string that:
+ Is sorted in alphabetical order.
+ Contains the longest possible combination of distinct letters, each appearing only once.
+ Includes letters from both string1 and string2.

Examples:
Input:
string1 = "xyaabbbccccdefww"
string2 = "xxxxyyyyabklmopq"
Output: "abcdefklmopqwxy"

Напишите функцию, которая принимает две строки string1 и string2, содержащие только строчные буквы от a до z. Функция должна вернуть новую строку, которая:
+ Отсортирована в алфавитном порядке.
+ Содержит максимально возможное количество различных букв, каждая из которых встречается только один раз.
+ Включает буквы из обеих строк string1 и string2.

Пример:
Ввод:
string1 = "xyaabbbccccdefww"
string2 = "xxxxyyyyabklmopq"
Вывод: "abcdefklmopqwxy"

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
	string1 := "xyaabbbccccdefww"
	string2 := "xxxxyyyyabklmopq"
	result := LongestDistinctCombination(string1, string2)
	fmt.Println(result)
}
