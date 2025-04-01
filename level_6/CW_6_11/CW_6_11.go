/*
Given two integers n and m, the task is to:
+ Calculate the first n triangular numbers.
+ Find their least common multiple (LCM).
+ Calculate the first m multiples of the LCM.
+ Return the sum of these multiples.

Examples:
Input:
n = 5
m = 8
Output:
1080

Input:
n = 7
m = 10
Output:
23100

Даны два целых числа n и m. Нужно:
+ Вычислить первые n треугольных чисел.
+ Найти их наименьшее общее кратное (НОК).
+ Вычислить первые m кратных этого НОК.
+ Вернуть сумму этих кратных.

Примеры:
Вход:
n = 5
m = 8
Выход:
1080

Вход:
n = 7
m = 10
Выход:
23100

https://www.codewars.com/kata/566afbfc8595f2e751000040
*/

package main

import (
	"math/big"
)

func findGreatestCommonDivisor(firstNumber, secondNumber *big.Int) *big.Int {
	zero := big.NewInt(0)
	for secondNumber.Cmp(zero) != 0 {
		firstNumber, secondNumber = secondNumber, new(big.Int).Mod(firstNumber, secondNumber)
	}
	return firstNumber
}

func findLeastCommonMultiple(firstNumber, secondNumber *big.Int) *big.Int {
	greatestCommonDivisor := findGreatestCommonDivisor(firstNumber, secondNumber)
	product := new(big.Int).Mul(firstNumber, secondNumber)
	return new(big.Int).Div(product, greatestCommonDivisor)
}

func calculateSumOfFirstMultiplesOfLCMOfTriangularNumbers(numberOfTriangularNumbers, numberOfMultiples int) *big.Int {
	triangularNumbers := make([]*big.Int, numberOfTriangularNumbers)
	for index := 1; index <= numberOfTriangularNumbers; index++ {
		triangularNumbers[index-1] = new(big.Int).Div(new(big.Int).Mul(big.NewInt(int64(index)), big.NewInt(int64(index+1))), big.NewInt(2))
	}

	leastCommonMultiple := triangularNumbers[0]
	for index := 1; index < numberOfTriangularNumbers; index++ {
		leastCommonMultiple = findLeastCommonMultiple(leastCommonMultiple, triangularNumbers[index])
	}

	totalSum := big.NewInt(0)
	for multipleIndex := 1; multipleIndex <= numberOfMultiples; multipleIndex++ {
		currentMultiple := new(big.Int).Mul(leastCommonMultiple, big.NewInt(int64(multipleIndex)))
		totalSum.Add(totalSum, currentMultiple)
	}

	return totalSum
}
