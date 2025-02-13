/*
Implement a function that takes three integers (a, b, c) as input. The function should return true if a valid triangle can be formed with these side lengths, and false otherwise.

Реализуйте функцию, которая принимает три целых числа (a, b, c) в качестве входных данных. Функция должна возвращать true, если из этих значений можно составить валидный треугольник, и false в противном случае.

https://www.codewars.com/kata/56606694ec01347ce800001b
*/

package main

import (
	"fmt"
)

func IsTriangle(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return a+b > c && a+c > b && b+c > a
}

func main() {
	a := 4
	b := 2
	c := 3
	result := IsTriangle(a, b, c)
	fmt.Println(result)
}
