/*
Write a function that takes two integers, a and b, which can be positive or negative. The function should calculate and return the sum of all integers between and including a and b. If a and b are equal, return either of the numbers.

Examples:
(1, 0) → 1 (1 + 0 = 1)
(1, 2) → 3 (1 + 2 = 3)
(-1, 2) → 2 (-1 + 0 + 1 + 2 = 2)

Напишите функцию, которая принимает два целых числа a и b, которые могут быть положительными или отрицательными. Функция должна вычислить и вернуть сумму всех целых чисел между a и b, включая сами числа. Если a и b равны, вернуть любое из этих чисел.

Примеры:
(1, 0) → 1 (1 + 0 = 1)
(1, 2) → 3 (1 + 2 = 3)
(-1, 2) → 2 (-1 + 0 + 1 + 2 = 2)

https://www.codewars.com/kata/55f2b110f61eb01779000053
*/

package main

import (
	"fmt"
)

func SumBetween(a, b int) int {
	if a > b {
		a, b = b, a
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}

func main() {
	fmt.Println(SumBetween(1, 0))
	fmt.Println(SumBetween(1, 2))
	fmt.Println(SumBetween(-1, 2))
	fmt.Println(SumBetween(5, 5))
}
