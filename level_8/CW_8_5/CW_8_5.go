/*
Write a function that converts an integer into a string.

For example:
Input: 123 → Output: "123"
Input: 999 → Output: "999"
Input: -100 → Output: "-100"

Напишите функцию, которая преобразует целое число в строку.

Например:
Ввод: 123 → Вывод: "123"
Ввод: 999 → Вывод: "999"
Ввод: -100 → Вывод: "-100"

https://www.codewars.com/kata/5265326f5fda8eb1160004c8
*/

package main

import (
	"fmt"
	"strconv"
)

func NumberToString(number int) string {
	return strconv.Itoa(number) // Integer to ASCII
}

func main() {
	fmt.Println(NumberToString(123))
	fmt.Println(NumberToString(999))
	fmt.Println(NumberToString(-100))
}
