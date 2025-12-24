/*
Write a function that checks whether a given string is a palindrome, ignoring letter case.
A palindrome is a word, number, phrase, or any sequence of characters that reads the same forwards and backwards, for example: madam or racecar.

Напишите функцию, которая проверяет, является ли заданная строка палиндромом, без учёта регистра символов.
Палиндром - это слово, число, фраза или любая последовательность символов, которая читается одинаково слева направо и справа налево, например: madam или racecar.
*/

package main

import "strings"

func IsPalindrome(input string) bool {
	lowered := strings.ToLower(input)

	left := 0
	right := len(lowered) - 1

	for left < right {
		if lowered[left] != lowered[right] {
			return false
		}
		left++
		right--
	}

	return true
}