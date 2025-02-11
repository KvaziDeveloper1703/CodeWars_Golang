/*
Write a function that converts a full name into initials.
+ The input will always consist of exactly two words, separated by a single space.
+ The output should be two capital letters, separated by a dot (.).

Examples:
"Sam Harris" → "S.H"
"patrick feeney" → "P.F"

Напиши функцию, которая преобразует полное имя в инициалы.
+ Входные данные всегда состоят ровно из двух слов, разделённых одним пробелом.
+ Выходные данные должны содержать две заглавные буквы, разделённые точкой (.).

Примеры:
"Sam Harris" → "S.H"
"patrick feeney" → "P.F"

https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3
*/

package main

import (
	"fmt"
	"strings"
)

func Abbreviation(givenName string) string {
	parts := strings.Split(givenName, " ")
	if len(parts) != 2 {
		return ""
	}
	return strings.ToUpper(string(parts[0][0])) + "." + strings.ToUpper(string(parts[1][0]))
}

func main() {
	name := "Sam Harris"
	abbreviation := Abbreviation(name)
	fmt.Println(abbreviation)
}
