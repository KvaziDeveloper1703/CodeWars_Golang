/*
Write a function that takes a number as input and returns the sum of all natural numbers below the given number that are multiples of 3 or 5. If a number is a multiple of both 3 and 5, count it only once.

Example:
For numbers below 10, the multiples of 3 or 5 are 3, 5, 6, and 9. Their sum is 23.

Напишите функцию, которая принимает число и возвращает сумму всех натуральных чисел меньше заданного числа, которые являются кратными 3 или 5. Если число кратно и 3, и 5, оно учитывается только один раз.

Пример:
Для чисел меньше 10 кратные 3 или 5: 3, 5, 6 и 9. Их сумма равна 23.

https://www.codewars.com/kata/514b92a657cdc65150000006
*/

package main

import (
	"fmt"
)

func SumMultiples(limit int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(SumMultiples(10))
	fmt.Println(SumMultiples(20))
	fmt.Println(SumMultiples(100))
}
