/*
You get an array of numbers, return the sum of all of the positives ones.

Example: [1,-4,7,12] => 20

https://www.codewars.com/kata/5715eaedb436cf5606000381
*/

package main

import "fmt"

func SumOfPositives(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}

func main() {
	arr := []int{1, -4, 7, 12}
	fmt.Println(SumOfPositives(arr))

	arr = []int{-11, -12, -13}
	fmt.Println(SumOfPositives(arr))

	arr = []int{}
	fmt.Println(SumOfPositives(arr))
}
