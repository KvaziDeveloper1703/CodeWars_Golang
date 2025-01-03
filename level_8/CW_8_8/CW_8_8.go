/*
It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, the original string. You don't have to worry about strings with less than two characters.

https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0
*/

package main

import "fmt"

func RemoveFirstAndLast(word string) string {
	if len(word) <= 2 {
		return ""
	}
	return word[1 : len(word)-1]
}

func main() {
	fmt.Println(RemoveFirstAndLast("hello"))
	fmt.Println(RemoveFirstAndLast("world"))
	fmt.Println(RemoveFirstAndLast("ab"))
	fmt.Println(RemoveFirstAndLast("a"))
}
