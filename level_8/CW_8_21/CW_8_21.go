/*
Create a function that takes two arguments:
	- salary - an integer;
	- bonus - a boolean.

If bonus is true, the salary should be multiplied by 10, and if bonus is false, the person receives only the base salary.

The function must return the final amount as a string, prefixed with the correct currency symbol: £ for Go.

Создайте функцию, которая принимает два аргумента:
	- salary - целое число;
	- bonus - логическое значение.

Если bonus равен true, зарплата должна быть умножена на 10. Если bonus равен false, человек получает только базовую зарплату.

Функция должна вернуть итоговую сумму в виде строки с соответствующим символом валюты: £ для Go.
*/

package main

import "fmt"

func BonusTime(salary int, bonus bool) string {
	if bonus {
		salary *= 10
	}
	return fmt.Sprintf("£%d", salary)
}