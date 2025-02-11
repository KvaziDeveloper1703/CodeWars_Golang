/*
Nathan loves cycling. Since he knows it's important to stay hydrated, he drinks 0.5 litres of water per hour of cycling.
You are given the time in hours, and your task is to return the number of litres Nathan will drink, rounded down to the nearest whole number.

Натан любит кататься на велосипеде. Поскольку он знает, как важно поддерживать водный баланс, он выпивает 0.5 литра воды за каждый час езды.
Тебе дано время в часах, и твоя задача — вычислить, сколько литров воды выпьет Натан, округлив результат вниз до целого числа.

https://www.codewars.com/kata/582cb0224e56e068d800003c
*/

package main

import (
	"fmt"
	"math"
)

func Litres(givenTime float64) int {
	return int(math.Floor(givenTime * 0.5))
}

func main() {
	time := 6.7
	litres := Litres(time)
	fmt.Println(litres)
}
