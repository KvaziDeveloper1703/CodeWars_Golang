/*
Complete the square sum function so that it squares each number passed into it and then sums the results together.

Example: [1,2,2] => 9
*/

package main

import "fmt"

func SquareSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2}))
	fmt.Println(SquareSum([]int{0, 3, 4, 5}))
	fmt.Println(SquareSum([]int{}))
}
