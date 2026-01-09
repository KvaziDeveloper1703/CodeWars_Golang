/*
Your task is to return the multiplication table for a given number.
The input number is always an integer from 1 to 10.
The result should be returned as a string, where each row represents one line of the multiplication table.

Requirements:
	- The table must include multipliers from 1 to 10;
	- Each row must follow this format: multiplier * number = result;
	- Use \n to separate rows;
	- Do not add a trailing newline at the end of the string.

Ваша задача - вернуть таблицу умножения для заданного числа.
Входное число всегда является целым числом от 1 до 10.
Результат должен быть возвращён в виде строки, где каждая строка представляет одну строку таблицы умножения.

Требования:
	- Таблица должна содержать множители от 1 до 10;
	- Каждая строка должна иметь формат: множитель * число = результат;
	- Для перехода на новую строку используйте символ \n;
	- В конце строки не должно быть лишнего переноса строки.
*/

package main

import "strconv"

func MultiTable(number int) string {
	result := ""
	for i := 1; i <= 10; i++ {
		line := strconv.Itoa(i) + " * " +
			strconv.Itoa(number) + " = " +
			strconv.Itoa(i*number)
		if i < 10 {
			line += "\n"
		}
		result += line
	}
	return result
}