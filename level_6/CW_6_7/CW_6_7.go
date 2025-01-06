/*
Write a function that takes a non-negative integer as input and returns the number of 1s in its binary representation.

Example:
Input: 1234
Binary representation: 10011010010
Output: 5

Напишите функцию, которая принимает неотрицательное целое число и возвращает количество единиц (1) в его двоичном представлении.

Пример:
Ввод: 1234
Двоичное представление: 10011010010
Вывод: 5

https://www.codewars.com/kata/526571aae218b8ee490006f4
*/

package main

import (
	"fmt"
	"strconv"
)

func CountOnesInBinary(number int) int {
	binary := strconv.FormatInt(int64(number), 2)
	count := 0
	for _, bit := range binary {
		if bit == '1' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountOnesInBinary(1234))
	fmt.Println(CountOnesInBinary(0))
	fmt.Println(CountOnesInBinary(15))
	fmt.Println(CountOnesInBinary(1023))
}
