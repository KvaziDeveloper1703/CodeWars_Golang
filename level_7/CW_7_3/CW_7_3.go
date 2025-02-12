/*
Write a function that takes a string of space-separated numbers as input and returns the highest and lowest numbers as a single string, separated by a space.

Напишите функцию, которая принимает строку, содержащую числа, разделённые пробелами, и возвращает самое большое и самое маленькое число в виде строки, разделённой пробелом.

https://www.codewars.com/kata/554b4ac871d6813a03000035
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(givenNumbers string) string {
	numbers := strings.Split(givenNumbers, " ")
	minimum, maximum := 0, 0
	for i, numberAsAString := range numbers {
		number, _ := strconv.Atoi(numberAsAString)
		if i == 0 || number < minimum {
			minimum = number
		}
		if i == 0 || number > maximum {
			maximum = number
		}
	}
	return fmt.Sprintf("%d %d", maximum, minimum)
}

func main() {
	numbers := "1 2 -3 4 5"
	result := HighAndLow(numbers)
	fmt.Println(result)
}
