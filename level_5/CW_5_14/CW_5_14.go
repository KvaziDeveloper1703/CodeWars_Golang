/*
You are given a string provided by the user. The string is guaranteed not to be null / nil / None, so you do not need to check for null values.

Your task is to determine whether the string is alphanumeric.

A string is considered alphanumeric if it meets all of the following conditions:
	- It contains at least one character.
	- It consists only of:
		- uppercase Latin letters (A-Z);
		- lowercase Latin letters (a-z);
		- digits (0-9).
	- It does not contain:
		- whitespaces;
		- underscores (_);
		- any other special characters.

Return True if the string is alphanumeric, otherwise return False.

Вам дана строка, введённая пользователем. Гарантируется, что строка не равна null / nil / None, поэтому проверку на null выполнять не нужно.

Необходимо определить, является ли строка алфавитно-цифровой.
	- Строка считается алфавитно-цифровой, если выполняются все следующие условия:
	- Строка содержит хотя бы один символ:
	- Строка состоит только из:
		- латинских букв в верхнем регистре (A-Z);
		- латинских букв в нижнем регистре (a-z);
		- цифр (0-9).
	- В строке отсутствуют:
		- пробелы;
		- символ подчёркивания (_);
		- любые другие специальные символы.

Верните True, если строка является алфавитно-цифровой, иначе верните False.
*/

package main

func alphanumeric(inputString string) bool {
	if len(inputString) == 0 {
		return false
	}

	for _, character := range inputString {
		if !(character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' || character >= '0' && character <= '9') {
			return false
		}
	}

	return true
}