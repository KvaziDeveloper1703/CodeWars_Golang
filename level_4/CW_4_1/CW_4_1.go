/*
You are given a sequence u defined as follows:
	- The first element is u(0) = 1;
	- For every number x in u, the numbers y = 2x + 1 and z = 3x + 1 must also be in u;
	- No other numbers belong to the sequence;
	- The sequence is sorted in ascending order and contains no duplicates.

Example of the beginning of the sequence: u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]

Write a function dbl_linear(n) that returns the n-th element of the ordered sequence u.

Задана последовательность u, определённая следующим образом:
	- Первый элемент: u(0) = 1;
	- Для каждого числа x из u в последовательности также должны быть числа: y = 2x + 1 и z = 3x + 1;
	- Других чисел в последовательности нет;
	- Последовательность упорядочена по возрастанию и не содержит повторов.

Начало последовательности: u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]

Реализуйте функцию dbl_linear(n), которая возвращает n-й элемент упорядоченной последовательности u.
*/

package kata

func DblLinear(n int) int {
	sequence := make([]int, n+1)
	sequence[0] = 1

	indexForTwo := 0
	indexForThree := 0

	for currentPosition := 1; currentPosition <= n; currentPosition++ {
		valueFromTwo := 2*sequence[indexForTwo] + 1
		valueFromThree := 3*sequence[indexForThree] + 1

		if valueFromTwo < valueFromThree {
			sequence[currentPosition] = valueFromTwo
			indexForTwo++
		} else if valueFromTwo > valueFromThree {
			sequence[currentPosition] = valueFromThree
			indexForThree++
		} else {
			sequence[currentPosition] = valueFromTwo
			indexForTwo++
			indexForThree++
		}
	}

	return sequence[n]
}