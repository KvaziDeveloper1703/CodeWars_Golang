/*
Given a triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29

Implement a function that calculates the sum of the numbers in the n-th row of this triangle.
The row index starts at 1.

Дан треугольник последовательных нечётных чисел:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29

Реализуйте функцию, которая вычисляет сумму чисел в n-й строке этого треугольника.
Индексация строк начинается с 1.

https://www.codewars.com/kata/55fd2d567d94ac3bc9000064
*/

package main

import "fmt"

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func main() {
	n := 4
	result := RowSumOddNumbers(n)
	fmt.Println(result)
}
