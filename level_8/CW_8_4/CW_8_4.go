/*
Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'
'word'   =>  'drow'

https://www.codewars.com/kata/5168bb5dfe9a00b126000018
*/

package main

import "fmt"

func Solution(word string) string {
	runes := []rune(word)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(Solution("world"))
	fmt.Println(Solution("word"))
	fmt.Println(Solution("a"))
	fmt.Println(Solution(""))
}
