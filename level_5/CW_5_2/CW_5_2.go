/*
Write a function that takes a non-negative integer (representing seconds) as input and returns the time formatted as a human-readable string in the format HH:MM:SS:
HH: Hours, padded to 2 digits (range: 00 to 99)
MM: Minutes, padded to 2 digits (range: 00 to 59)
SS: Seconds, padded to 2 digits (range: 00 to 59)

The input will not exceed 359999 seconds, which corresponds to the maximum time 99:59:59.

Напишите функцию, которая принимает неотрицательное целое число (количество секунд) и возвращает строку с временем в формате, понятном человеку: HH:MM:SS:
HH: Часы, дополненные до двух цифр (диапазон: 00 до 99)
MM: Минуты, дополненные до двух цифр (диапазон: 00 до 59)
SS: Секунды, дополненные до двух цифр (диапазон: 00 до 59)

Входное значение не превышает 359999 секунд, что соответствует максимальному времени 99:59:59.

https://www.codewars.com/kata/52685f7382004e774f0001f7
*/

package main

import (
	"fmt"
)

func FormatTime(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func main() {
	fmt.Println(FormatTime(0))
	fmt.Println(FormatTime(59))
	fmt.Println(FormatTime(60))
	fmt.Println(FormatTime(3599))
	fmt.Println(FormatTime(3600))
	fmt.Println(FormatTime(86399))
	fmt.Println(FormatTime(359999))
}
