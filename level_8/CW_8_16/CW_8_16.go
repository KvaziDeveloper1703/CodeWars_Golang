/*
Given a year, return the century it belongs to.

Напиши функцию, которая принимает год и возвращает век, к которому он относится.

https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097
*/

package main

import "fmt"

func CenturyFromYear(givenYear int) int {
	if givenYear%100 == 0 {
		return givenYear / 100
	}
	return givenYear/100 + 1
}

func main() {
	year := 1703
	century := CenturyFromYear(year)
	fmt.Println(century)
}
