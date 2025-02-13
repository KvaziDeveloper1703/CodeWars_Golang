/*
In a small town, the population starts at p0 = 1000 at the beginning of the year. The population increases by 2% per year, and additionally, 50 new inhabitants move to the town each year.
How many years will it take for the town’s population to reach or exceed p = 1200 inhabitants?

Given the parameters:
+ p0 – Initial population;
+ percent – Annual percentage growth;
+ aug – Number of people moving in or out each year;
+ p – Target population.

В небольшом городе начальная численность населения p0 = 1000 человек. Каждый год население увеличивается на 2%, а также в город переезжает 50 новых жителей.
Сколько лет потребуется, чтобы население достигло или превысило p = 1200 человек?

Даны параметры:
+ p0 – начальное население;
+ percent – ежегодный процентный прирост;
+ aug – число жителей, ежегодно прибывающих из города;
+ p – целевая численность населения.

https://www.codewars.com/kata/563b662a59afc2b5120000c6
*/

package main

import (
	"fmt"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	for p0 < p {
		p0 += int(float64(p0)*percent/100) + aug
		years++
	}
	return years
}

func main() {
	p0 := 1500
	percent := 5.0
	aug := 100
	p := 5000
	result := NbYear(p0, percent, aug, p)
	fmt.Println(result)
}
