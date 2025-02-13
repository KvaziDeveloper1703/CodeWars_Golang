/*
Ben has a simple idea to make a profit: he buys something and sells it again. Of course, there’s no profit if he buys and sells at the same price. Instead, he wants to buy at the lowest price and sell at the highest price.
Write a function that returns both the minimum and maximum numbers from a given list/array.

У Бена есть простая идея для получения прибыли: купить товар дешевле и продать дороже. Конечно, если цена покупки и продажи одинаковая, прибыли не будет. Поэтому он хочет купить по самой низкой цене и продать по самой высокой.
Напишите функцию, которая возвращает минимальное и максимальное число из переданного списка/массива.

https://www.codewars.com/kata/559590633066759614000063
*/

package main

import (
	"fmt"
)

func MinMax(givenArray []int) [2]int {
	min, max := givenArray[0], givenArray[0]

	for _, number := range givenArray {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	return [2]int{min, max}
}

func main() {
	myArray := []int{1, 2, 3, 4, 5}
	result := MinMax(myArray)
	fmt.Println(result)
}
