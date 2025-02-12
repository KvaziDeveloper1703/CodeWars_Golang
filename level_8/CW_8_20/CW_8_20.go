/*
Timmy & Sarah think they are in love, but in their town, they will only know for sure after picking a flower each.
If one flower has an even number of petals and the other has an odd number of petals, it means they are in love.
Write a function that takes the number of petals from each flower and returns true if they are in love, and false if they aren't.

Тимми и Сара думают, что влюблены, но в их городе это можно проверить только с помощью цветов.
Если у одного цветка четное количество лепестков, а у другого — нечетное, значит, они влюблены.
Напиши функцию, которая принимает количество лепестков у каждого цветка и возвращает true, если они влюблены, и false, если нет.

https://www.codewars.com/kata/555086d53eac039a2a000083
*/

package main

import (
	"fmt"
)

func InLove(flower_1, flower_2 int) bool {
	return (flower_1%2 == 0 && flower_2%2 != 0) || (flower_1%2 != 0 && flower_2%2 == 0)
}

func main() {
	flower_1 := 4
	flower_2 := 7
	result := InLove(flower_1, flower_2)
	fmt.Println(result)
}
