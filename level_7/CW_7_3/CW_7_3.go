/*
Write a function that takes a string of space-separated numbers as input and returns the highest and lowest numbers as a single string, separated by a space.

For example:
Input: "1 2 3 4 5"
Output: "5 1"

Напишите функцию, которая принимает строку, содержащую числа, разделённые пробелами, и возвращает самое большое и самое маленькое число в виде строки, разделённой пробелом.

Например:
Ввод: "1 2 3 4 5"
Вывод: "5 1"

https://www.codewars.com/kata/554b4ac871d6813a03000035
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(input string) string {
	numbers := strings.Split(input, " ")
	min, max := 0, 0
	for i, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		if i == 0 || num < min {
			min = num
		}
		if i == 0 || num > max {
			max = num
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}
