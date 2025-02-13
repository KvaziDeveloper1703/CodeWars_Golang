/*
A bus is moving through the city, picking up and dropping off passengers at each bus stop.
You are given a list of integer pairs. Each pair consists of:
+ First element → the number of people who get on the bus.
+ Second element → the number of people who get off the bus.

Your task is to return the number of people still on the bus after the last bus stop (after processing the last pair in the array).
Even though it is the last stop, the bus might not be empty, meaning some passengers could still be inside (maybe sleeping).

Автобус движется по городу, высаживая и подбирая пассажиров на каждой остановке.
Вам дан список пар чисел. Каждая пара состоит из:
+ Первый элемент → количество людей, которые садятся в автобус.
+ Второй элемент → количество людей, которые выходят из автобуса.

Ваша задача — вернуть количество людей, оставшихся в автобусе после последней остановки (после обработки последней пары в массиве).
Даже после последней остановки автобус может быть не пуст, так как некоторые пассажиры могут всё ещё находиться внутри (может быть, они уснули).

https://www.codewars.com/kata/5648b12ce68d9daa6b000099
*/

package main

import "fmt"

func Number(busStops [][2]int) int {
	people := 0
	for _, stop := range busStops {
		people += stop[0]
		people -= stop[1]
	}
	return people
}

func main() {
	stops := [][2]int{{10, 0}, {3, 5}, {5, 8}}
	result := Number(stops)
	fmt.Println(result)
}
