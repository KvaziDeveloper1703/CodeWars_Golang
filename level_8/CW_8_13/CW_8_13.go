/*
Given an array of sheep, where some sheep may be missing, write a function that counts and returns the number of sheep present in the array. A true value means the sheep is present.

Example:
[True,  True,  True,  False,
 True,  True,  True,  True,
 True,  False, True,  False,
 True,  False, False, True,
 True,  True,  True,  True,
 False, False, True,  True]

The correct output should be 17.

Дан массив овец, где некоторые могут отсутствовать. Напиши функцию, которая подсчитывает и возвращает количество присутствующих овец. Значение true означает, что овца присутствует.

Пример:
[True,  True,  True,  False,
 True,  True,  True,  True,
 True,  False, True,  False,
 True,  False, False, True,
 True,  True,  True,  True,
 False, False, True,  True]

Правильный ответ: 17.

https://www.codewars.com/kata/54edbc7200b811e956000556
*/

package main

import (
	"fmt"
)

func CountSheep(sheep []bool) int {
	count := 0
	for _, present := range sheep {
		if present {
			count++
		}
	}
	return count
}

func main() {
	sheep := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	result := CountSheep(sheep)
	fmt.Println(result)
}
