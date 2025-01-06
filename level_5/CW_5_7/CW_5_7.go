/*
Pete wants to bake cakes but needs help figuring out how many he can make with the ingredients he has. Write a function cakes() that takes two arguments:
+ recipe (an object representing the required amounts of each ingredient);
+ available (an object representing the available amounts of each ingredient).

The function should return the maximum number of cakes Pete can bake (an integer). If an ingredient is missing in available, consider it as 0.

Examples:
Input: cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200})
Output: 2

Пит любит печь торты, но не умеет считать, сколько он может сделать с учетом имеющихся ингредиентов. Напишите функцию cakes(), которая принимает два аргумента:
+ recipe (объект, представляющий требуемое количество каждого ингредиента);
+ available (объект, представляющий доступное количество каждого ингредиента).

Функция должна возвращать максимальное количество тортов, которые Пит сможет испечь (целое число). Если какого-то ингредиента нет в available, считать его равным 0.

Примеры:
Ввод: cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200})
Вывод: 2

https://www.codewars.com/kata/525c65e51bf619685c000059
*/

package main

import (
	"fmt"
)

func cakes(recipe, available map[string]int) int {
	maxCakes := -1

	for ingredient, amountNeeded := range recipe {
		amountAvailable := available[ingredient]
		if amountAvailable < amountNeeded {
			return 0
		}

		numCakes := amountAvailable / amountNeeded
		if maxCakes == -1 || numCakes < maxCakes {
			maxCakes = numCakes
		}
	}

	return maxCakes
}

func main() {
	recipe := map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	available := map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}
	fmt.Println(cakes(recipe, available))

	recipe = map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	available = map[string]int{"flour": 500, "sugar": 200}
	fmt.Println(cakes(recipe, available))

	recipe = map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	available = map[string]int{"sugar": 200, "eggs": 1}
	fmt.Println(cakes(recipe, available))
}
