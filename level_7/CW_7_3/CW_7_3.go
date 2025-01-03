/*
In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

https://www.codewars.com/kata/554b4ac871d6813a03000035
*/

package main

import (
	"fmt"
	"strconv"
)

func HighAndLow(input string) string {
	min, max := 0, 0
	current := ""
	first := true

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' || i == len(input)-1 {
			if i == len(input)-1 && input[i] != ' ' {
				current += string(input[i])
			}
			num, _ := strconv.Atoi(current)
			if first {
				min, max = num, num
				first = false
			} else {
				if num > max {
					max = num
				}
				if num < min {
					min = num
				}
			}
			current = ""
		} else {
			current += string(input[i])
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}
