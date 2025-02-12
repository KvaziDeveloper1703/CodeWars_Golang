/*
Write a function that takes two integers, a and b, which can be positive or negative. The function should calculate and return the sum of all integers between and including a and b. If a and b are equal, return either of the numbers.

Напишите функцию, которая принимает два целых числа a и b, которые могут быть положительными или отрицательными. Функция должна вычислить и вернуть сумму всех целых чисел между a и b, включая сами числа. Если a и b равны, вернуть любое из этих чисел.

https://www.codewars.com/kata/55f2b110f61eb01779000053
*/

package main

import (
	"fmt"
)

func SumBetween(number1, number2 int) int {
	if number1 > number2 {
		number1, number2 = number2, number1
	}

	sum := 0
	for i := number1; i <= number2; i++ {
		sum += i
	}

	return sum
}

func main() {
	number1 := 63
	number2 := 27
	result := SumBetween(number1, number2)
	fmt.Println(result)
}
