/*
Create a function that accepts a parameter representing a name and returns the message: "Hello, <name> how are you doing today?"

Создай функцию, которая принимает параметр — имя и возвращает сообщение: "Hello, <имя> how are you doing today?"

https://www.codewars.com/kata/55a70521798b14d4750000a4
*/

package main

import (
	"fmt"
)

func Greet(givenName string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", givenName)
}

func main() {
	name := "Viktor"
	greeting := Greet(name)
	fmt.Println(greeting)
}
