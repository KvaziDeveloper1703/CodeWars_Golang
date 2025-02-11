/*
We need a function that converts a string into a number.

Нам нужна функция, которая преобразует строку в число.

https://www.codewars.com/kata/544675c6f971f7399a000e79
*/

package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(givenString string) int {
	number, err := strconv.Atoi(givenString)
	if err != nil {
		fmt.Println("Error converting string to number:", err)
		return 0
	}
	return number
}

func main() {
	myString := "1234"
	number := StringToNumber(myString)
	fmt.Println(number)
}
