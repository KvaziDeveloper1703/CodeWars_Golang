/*
We need a function that can transform a number (integer) into a string.

Examples:
123  --> "123"
999  --> "999"
-100 --> "-100"

https://www.codewars.com/kata/5265326f5fda8eb1160004c8
*/

package main

import (
	"fmt"
	"strconv"
)

func NumberToString(number int) string {
	return strconv.Itoa(number)
}

func main() {
	fmt.Println(NumberToString(123))
	fmt.Println(NumberToString(999))
	fmt.Println(NumberToString(-100))
}
