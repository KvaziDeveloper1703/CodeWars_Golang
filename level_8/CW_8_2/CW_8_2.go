/*
In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

https://www.codewars.com/kata/55685cd7ad70877c23000102
*/

package main

import "fmt"

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(MakeNegative(1))  // -1
	fmt.Println(MakeNegative(-1)) // -1
	fmt.Println(MakeNegative(0))  // 0
}
