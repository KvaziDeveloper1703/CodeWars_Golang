/*
Write a function to sort a string of numbers by their "weight," where the weight is the sum of the digits.

Rules:
+ Sort numbers by their "weight" (sum of digits).
+ If weights are equal, sort them as strings alphabetically.

Example:
Input: "56 65 74 100 99 68 86 180 90"
Output: "100 180 90 56 65 74 68 86 99"

Напишите функцию для сортировки строки с числами по их "весу", где вес — это сумма цифр числа.

Правила:
+ Сортировать числа по "весу" (сумме цифр).
+ При одинаковом весе сортировать как строки в алфавитном порядке.

Пример:
Ввод: "56 65 74 100 99 68 86 180 90"
Вывод: "100 180 90 56 65 74 68 86 99"

https://www.codewars.com/kata/55c6126177c9441a570000cc
*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func weight(number string) int {
	sum := 0
	for _, digit := range number {
		if unicode.IsDigit(digit) {
			sum += int(digit - '0')
		}
	}
	return sum
}

func orderWeight(input string) string {
	numbers := strings.Fields(input)

	sort.SliceStable(numbers, func(i, j int) bool {
		weightI := weight(numbers[i])
		weightJ := weight(numbers[j])
		if weightI == weightJ {
			return numbers[i] < numbers[j]
		}
		return weightI < weightJ
	})

	return strings.Join(numbers, " ")
}

func main() {
	input := "56 65 74 100 99 68 86 180 90"
	output := orderWeight(input)
	fmt.Println(output)
}
