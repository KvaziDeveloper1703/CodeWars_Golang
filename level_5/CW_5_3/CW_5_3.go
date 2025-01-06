/*
Write a function that takes three integers as input, representing RGB color values (red, green, and blue). The function should return a hexadecimal color code string.
+ Each RGB value must be between 0 and 255.
+ If an input value is outside this range, round it to the nearest valid value (0 or 255).
+ The returned hexadecimal string must always be 6 characters long (e.g., "FFFFFF" or "9400D3").

Examples:
Input: 255, 255, 255 → Output: "FFFFFF"
Input: 255, 255, 300 → Output: "FFFFFF"
Input: 0, 0, 0 → Output: "000000"
Input: 148, 0, 211 → Output: "9400D3"

Напишите функцию, которая принимает три целых числа, представляющих значения RGB (красный, зелёный и синий), и возвращает строку в виде шестизначного шестнадцатеричного кода цвета.
+ Каждое значение RGB должно быть в диапазоне от 0 до 255.
+ Если входное значение выходит за этот диапазон, оно должно быть округлено к ближайшему допустимому значению (0 или 255).
+ Возвращаемая строка в шестнадцатеричном формате должна всегда содержать ровно 6 символов (например, "FFFFFF" или "9400D3").

Примеры:
Ввод: 255, 255, 255 → Вывод: "FFFFFF"
Ввод: 255, 255, 300 → Вывод: "FFFFFF"
Ввод: 0, 0, 0 → Вывод: "000000"
Ввод: 148, 0, 211 → Вывод: "9400D3"

https://www.codewars.com/kata/513e08acc600c94f01000001
*/

package main

import (
	"fmt"
)

func Clamp(value int) int {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return value
}

func RGBToHex(r, g, b int) string {
	r = Clamp(r)
	g = Clamp(g)
	b = Clamp(b)

	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

func main() {
	fmt.Println(RGBToHex(255, 255, 255))
	fmt.Println(RGBToHex(255, 255, 300))
	fmt.Println(RGBToHex(0, 0, 0))
	fmt.Println(RGBToHex(148, 0, 211))
	fmt.Println(RGBToHex(-20, 500, 128))
}
