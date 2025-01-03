/*
Write a program that calculates the summation of all numbers from 1 to a given positive integer num. The input will always be a positive integer greater than 0. Your function should only return the result.

Напишите программу, которая вычисляет сумму всех чисел от 1 до заданного положительного целого числа num. Ввод всегда будет положительным целым числом больше 0. Функция должна возвращать только результат.

https://www.codewars.com/kata/55d24f55d7dd296eb9000030
*/

package main

import "fmt"

func Summation(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(Summation(5))
	fmt.Println(Summation(10))
	fmt.Println(Summation(1))
}
