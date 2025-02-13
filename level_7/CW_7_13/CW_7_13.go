/*
In a factory, a printer prints labels for boxes. For a specific type of box, the printer must use colors represented by letters from 'a' to 'm'.

The colors used by the printer are recorded in a control string. A "good" control string consists only of letters from 'a' to 'm'.
For example, "aaabbbbhaijjjm" is a valid control string.

However, sometimes errors occur due to a lack of colors or technical malfunctions, resulting in a "bad" control string.
For instance, "aaaxbbbbyyhwawiwjjjwwm" contains letters outside the range 'a' to 'm', making it a bad control string.

You need to write a function printer_error that takes a control string as input and returns the error rate as a fraction (string format).
The numerator is the count of incorrect letters (those outside 'a' to 'm'), and the denominator is the total length of the control string.
The fraction must not be simplified.

На фабрике принтер печатает этикетки для коробок. Для одного вида коробок принтер должен использовать только цвета, обозначенные буквами от 'a' до 'm'.

Использованные цвета записываются в контрольную строку.
Корректная ("хорошая") строка состоит только из букв от 'a' до 'm'.
Например, "aaabbbbhaijjjm" – это правильная строка.

Иногда из-за нехватки цветов или технических сбоев происходит ошибка, и формируется "плохая" контрольная строка.
Например, "aaaxbbbbyyhwawiwjjjwwm" содержит буквы за пределами 'a'—'m', значит, это ошибочная строка.

Нужно написать функцию printer_error, которая принимает контрольную строку и возвращает долю ошибок в виде дроби (строка).
Числитель — количество неправильных букв (вне диапазона 'a'—'m'), знаменатель — общая длина строки.
Дробь не упрощать.

https://www.codewars.com/kata/56541980fa08ab47a0000040
*/

package main

import (
	"fmt"
)

func PrinterError(givenString string) string {
	total := len(givenString)
	errors := 0

	for _, character := range givenString {
		if character < 'a' || character > 'm' {
			errors++
		}
	}

	return fmt.Sprintf("%d/%d", errors, total)
}

func main() {
	myString := "aaaxbbbbyyhwawiwjjjwwm"
	fmt.Println(PrinterError(myString))
}
