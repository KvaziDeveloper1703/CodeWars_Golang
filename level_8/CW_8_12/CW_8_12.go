/*
Write a function that removes all spaces from a given string and returns the resulting string.

Напиши функцию, которая удаляет все пробелы из переданной строки и возвращает получившуюся строку.

https://www.codewars.com/kata/57eae20f5500ad98e50002c5
*/

package main

import (
	"fmt"
	"strings"
)

func NoSpace(givenString string) string {
	return strings.ReplaceAll(givenString, " ", "")
}

func main() {
	myString := "8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd"
	result := NoSpace(myString)
	fmt.Println(result)
}
