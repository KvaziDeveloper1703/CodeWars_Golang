/*
Write a function dirReduc that takes an array of directions ("NORTH", "SOUTH", "EAST", "WEST") and simplifies the path by removing consecutive opposite directions ("NORTH" + "SOUTH" or "EAST" + "WEST").

Examples:
["NORTH", "SOUTH", "EAST", "WEST"] → []
["NORTH", "WEST", "SOUTH", "EAST"] → ["NORTH", "WEST", "SOUTH", "EAST"]
["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"] → ["WEST"]

Напишите функцию dirReduc, которая принимает массив направлений ("NORTH", "SOUTH", "EAST", "WEST") и упрощает маршрут, удаляя подряд идущие противоположные направления ("NORTH" + "SOUTH" или "EAST" + "WEST").

Примеры:
["NORTH", "SOUTH", "EAST", "WEST"] → []
["NORTH", "WEST", "SOUTH", "EAST"] → ["NORTH", "WEST", "SOUTH", "EAST"]
["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"] → ["WEST"]

https://www.codewars.com/kata/550f22f4d758534c1100025a
*/

package main

import (
	"fmt"
)

func dirReduc(directions []string) []string {

	opposites := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	stack := []string{}

	for _, dir := range directions {
		if len(stack) > 0 && stack[len(stack)-1] == opposites[dir] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, dir)
		}
	}

	return stack
}

func main() {
	fmt.Println(dirReduc([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
	fmt.Println(dirReduc([]string{"NORTH", "WEST", "SOUTH", "EAST"}))
	fmt.Println(dirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}))
}
