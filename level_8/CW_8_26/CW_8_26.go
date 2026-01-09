/*
Create a function close_compare that takes three parameters: a, b, and an optional parameter margin.

The function should determine whether a is less than, close to, or greater than b, and return a corresponding value.

Rules:
	- If a is close to b, return 0;
	- For this task, a is considered close to b if margin is greater than or equal to |a - b|;
	- Otherwise:
		- If a < b, return -1;
		- If a > b, return 1.
	- If margin is not provided, treat it as 0;
	- You may assume that margin ≥ 0.

Examples:
Input: a = 3, b = 5, margin = 3
Output: 0

Input: a = 3, b = 5, margin = 0
Output: -1

Создайте функцию close_compare, которая принимает три параметра: a, b и необязательный параметр margin.

Функция должна определить, является ли a меньше, близким к или больше b, и вернуть соответствующее значение.

Правила:
	- Если a близко к b, верните 0;
	- В рамках этой задачи a считается близким к b, если margin больше либо равен |a − b|;
	- В противном случае:
		- Если a < b, верните -1;
		- Если a > b, верните 1;
	- Если параметр margin не передан, считать его равным 0;
	- Можно считать, что margin ≥ 0.

Примеры:
Ввод: a = 3, b = 5, margin = 3
Вывод: 0

Ввод: a = 3, b = 5, margin = 0
Вывод: -1
*/

package main

import "math"

func CloseCompare(a, b, margin float64) int {
	if math.Abs(a-b) <= margin {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}