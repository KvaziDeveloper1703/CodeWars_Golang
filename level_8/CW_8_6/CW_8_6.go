/*
Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.

https://www.codewars.com/kata/53369039d7ab3ac506000467
*/

package main

import "fmt"

func BoolToWord(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

func main() {
	fmt.Println(BoolToWord(true))
	fmt.Println(BoolToWord(false))
}
