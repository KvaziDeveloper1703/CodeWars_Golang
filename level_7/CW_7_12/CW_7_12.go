/*
Complete the findNextSquare function, which finds the next integral perfect square after the given number.
A perfect square is an integer n such that sqrt(n) is also an integer. If the given number is not a perfect square, return -1.

Напиши функцию findNextSquare, которая находит следующий идеальный квадрат после переданного числа.
Идеальный квадрат — это число n, для которого sqrt(n) также является целым числом. Если переданное число не является идеальным квадратом, верни -1.

https://www.codewars.com/kata/56269eb78ad2e4ced1000013
*/

package main

import (
	"fmt"
	"math"
)

func FindNextSquare(givenSquare int64) int64 {
	sqrt := int64(math.Sqrt(float64(givenSquare)))
	if sqrt*sqrt != givenSquare {
		return -1
	}
	return (sqrt + 1) * (sqrt + 1)
}

func main() {
	square := int64(121)
	nextSquare := FindNextSquare(square)
	fmt.Println(nextSquare)
}
