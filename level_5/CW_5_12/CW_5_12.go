/*
In this kata, you are required to write a function that identifies peaks in a numeric array.

A peak is a value that is greater than the values immediately before and after it.

Write a function that returns the positions and values of all peaks in a given integer array.

The result must be returned as an object with two properties:
	- pos - an array containing the indices of the peaks;
	- peaks - an array containing the corresponding peak values.

If the array contains no peaks, the function must return an object with two empty arrays.

Rules and Conditions:
	- All input arrays contain valid integers and may be empty;
	- Input validation is not required;
	- The first and last elements of the array are not considered peaks, as they do not have neighbors on both sides.

A plateau is a sequence of equal values.
	- A plateau forms a peak only if the values before and after the plateau are smaller;
	- If a plateau is a peak, return only the position and value of the first element of the plateau;
	- If the plateau does not descend afterward, it is not considered a peak.

В этом задании необходимо написать функцию, которая находит пики (локальные максимумы) в числовом массиве.

Пик - это элемент массива, значение которого больше значений элементов, находящихся непосредственно слева и справа от него.

Напишите функцию, которая возвращает позиции и значения всех пиков в массиве целых чисел.

Результат должен быть возвращён в виде объекта с двумя полями:
	- pos - массив индексов пиков;
	- peaks - массив соответствующих значений.

Если в массиве нет пиков, функция должна вернуть объект с двумя пустыми массивами.

Правила и условия:
	- Все входные массивы содержат корректные целые числа и могут быть пустыми;
	- Проверка входных данных не требуется;
	- Первый и последний элементы массива не считаются пиками, так как у них нет двух соседей.

Плато - это последовательность одинаковых значений.
	- Плато считается пиком, только если значения до и после него меньше;
	- Если плато является пиком, необходимо вернуть позицию и значение первого элемента плато;
	- Если после плато нет снижения, пик не фиксируется.
*/

package main

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	result := PosPeaks{
		Pos:   []int{},
		Peaks: []int{},
	}

	if len(array) < 3 {
		return result
	}

	potentialPeakPosition := -1

	for index := 1; index < len(array); index++ {
		if array[index] > array[index-1] {
			potentialPeakPosition = index
		} else if array[index] < array[index-1] && potentialPeakPosition != -1 {
			result.Pos = append(result.Pos, potentialPeakPosition)
			result.Peaks = append(result.Peaks, array[potentialPeakPosition])
			potentialPeakPosition = -1
		}
	}

	return result
}