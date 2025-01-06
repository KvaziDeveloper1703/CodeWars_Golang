/*
Write a function that takes an array of integers as input and returns the integer that appears an odd number of times. There will always be exactly one integer in the array that satisfies this condition.

Examples:
[7] → 7 (appears 1 time, which is odd)
[0] → 0 (appears 1 time, which is odd)
[1, 1, 2] → 2 (appears 1 time, which is odd)
[0, 1, 0, 1, 0] → 0 (appears 3 times, which is odd)
[1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1] → 4 (appears 1 time, which is odd)

Напишите функцию, которая принимает массив целых чисел и возвращает число, которое встречается нечётное количество раз. В массиве всегда будет ровно одно число, удовлетворяющее этому условию.

Примеры:
[7] → 7 (встречается 1 раз, что является нечётным числом)
[0] → 0 (встречается 1 раз, что является нечётным числом)
[1, 1, 2] → 2 (встречается 1 раз, что является нечётным числом)
[0, 1, 0, 1, 0] → 0 (встречается 3 раза, что является нечётным числом)
[1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1] → 4 (встречается 1 раз, что является нечётным числом)

https://www.codewars.com/kata/54da5a58ea159efa38000836
*/

package main

import (
	"fmt"
)

func FindOdd(seq []int) int {
	result := 0

	for _, num := range seq {
		result ^= num
	}

	return result
}

func main() {
	fmt.Println(FindOdd([]int{7}))
	fmt.Println(FindOdd([]int{0}))
	fmt.Println(FindOdd([]int{1, 1, 2}))
	fmt.Println(FindOdd([]int{0, 1, 0, 1, 0}))
	fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}
