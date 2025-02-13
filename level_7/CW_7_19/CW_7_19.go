/*
Write a function that returns the n-th term of the following series, which is the sum of the first n terms of the sequence.

Series: 1+(1/4)+(1/7)+(1/10)+(1/13)+(1/16)

Rules:
+ The result must be rounded to 2 decimal places and returned as a string;
+ If n == 0, the function should return "0.00";
+ The input will always be a natural number.

Напишите функцию, которая возвращает n-й член следующего ряда, то есть сумму первых n членов последовательности.

Ряд: 1+(1/4)+(1/7)+(1/10)+(1/13)+(1/16)

Правила:
+ Результат должен быть округлён до 2 знаков после запятой и возвращён в виде строки;
+ Если n == 0, функция должна вернуть "0.00";
+ Входные данные — всегда натуральное число.

https://www.codewars.com/kata/555eded1ad94b00403000071
*/

package main

import (
	"fmt"
)

func SeriesSum(N int) string {
	if N == 0 {
		return "0.00"
	}

	sum := 0.0
	denominator := 1.0

	for i := 0; i < N; i++ {
		sum += 1 / denominator
		denominator += 3
	}

	return fmt.Sprintf("%.2f", sum)
}

func main() {
	N := 5
	result := SeriesSum(N)
	fmt.Println(result)
}
