/*
Write a function that calculates the digital root of a non-negative integer n. The digital root is the result of recursively summing all the digits of n until a single-digit number is obtained.

For example:
Input: 16 → Steps: 1 + 6 = 7 → Output: 7
Input: 942 → Steps: 9 + 4 + 2 = 15, then 1 + 5 = 6 → Output: 6

Напишите функцию, которая вычисляет цифровой корень неотрицательного целого числа n. Цифровой корень — это результат последовательного суммирования всех цифр числа n до тех пор, пока не останется одна цифра.

Пример:
Ввод: 16 → Шаги: 1 + 6 = 7 → Вывод: 7
Ввод: 942 → Шаги: 9 + 4 + 2 = 15, затем 1 + 5 = 6 → Вывод: 6

https://www.codewars.com/kata/541c8630095125aba6000c00
*/

package main

import (
	"fmt"
)

func DigitalRoot(number int) int {
	for number >= 10 {
		sum := 0
		for number > 0 {
			sum += number % 10
			number /= 10
		}
		number = sum
	}
	return number
}

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(942))
	fmt.Println(DigitalRoot(132189))
	fmt.Println(DigitalRoot(493193))
}
