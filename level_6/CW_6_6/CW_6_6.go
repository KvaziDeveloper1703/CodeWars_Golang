/*
Write a function that takes an array of exactly 10 integers (each between 0 and 9) as input and returns a string formatted as a phone number in the format "(XXX) XXX-XXXX".

Example:
Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
Output: "(123) 456-7890"

Напишите функцию, которая принимает массив из ровно 10 целых чисел (каждое от 0 до 9) и возвращает строку, отформатированную как номер телефона в формате "(XXX) XXX-XXXX".

Пример:
Ввод: [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
Вывод: "(123) 456-7890"

https://www.codewars.com/kata/525f50e3b73515a6db000b83
*/

package main

import (
	"fmt"
)

func FormatPhoneNumber(numbers []int) string {
	if len(numbers) != 10 {
		return "Invalid input: array must contain exactly 10 integers"
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5],
		numbers[6], numbers[7], numbers[8], numbers[9])
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(FormatPhoneNumber(numbers))

	invalidNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(FormatPhoneNumber(invalidNumbers))
}
