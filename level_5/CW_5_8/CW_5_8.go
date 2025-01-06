/*
Write a function that takes an integer prod and finds two consecutive Fibonacci numbers, F(n) and F(n+1):
+ Return (F(n),F(n+1),true) if F(n)×F(n+1)=prod.
+ Otherwise, return (F(n),F(n+1),false), where F(n)×F(n+1)>prod.

Examples:
Input: 714 → Output: (21, 34, true)
Input: 800 → Output: (34, 55, false)

Напишите функцию, которая принимает число prod и находит два последовательных числа Фибоначчи F(n) и F(n+1):
+ Вернуть (F(n),F(n+1),true), если F(n)×F(n+1)=prod.
+ Иначе вернуть (F(n),F(n+1),false), где F(n)×F(n+1)>prod.

Примеры:
Ввод: 714 → Вывод: (21, 34, true)
Ввод: 800 → Вывод: (34, 55, false)

https://www.codewars.com/kata/5541f58a944b85ce6d00006a
*/

package main

import (
	"fmt"
)

func ProductFib(prod uint64) (uint64, uint64, bool) {
	f1, f2 := uint64(0), uint64(1)

	for f1*f2 < prod {
		f1, f2 = f2, f1+f2
	}

	return f1, f2, f1*f2 == prod
}

func main() {
	fmt.Println(ProductFib(714))
	fmt.Println(ProductFib(800))
	fmt.Println(ProductFib(4895))
	fmt.Println(ProductFib(4))
}
