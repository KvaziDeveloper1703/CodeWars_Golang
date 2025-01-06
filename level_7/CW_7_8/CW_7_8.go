/*
Write a function that takes a DNA string as input and returns its complementary strand. In DNA, the symbols "A" and "T" are complements of each other, as are "C" and "G". The input will always be a valid DNA string, and it will never be empty.

Напишите функцию, которая принимает строку, представляющую одну цепь ДНК, и возвращает её комплементарную цепь. В ДНК символы "A" и "T" являются комплементарными, так же как "C" и "G". Входные данные всегда будут корректной строкой ДНК и никогда не будут пустыми.

https://www.codewars.com/kata/554e4a2f232cdd87d9000038
*/

package main

import (
	"fmt"
	"strings"
)

func GetComplementaryDNA(dna string) string {
	var complement strings.Builder

	for _, nucleotide := range dna {
		switch nucleotide {
		case 'A':
			complement.WriteRune('T')
		case 'T':
			complement.WriteRune('A')
		case 'C':
			complement.WriteRune('G')
		case 'G':
			complement.WriteRune('C')
		default:
			panic("Invalid DNA nucleotide")
		}
	}

	return complement.String()
}

func main() {
	fmt.Println(GetComplementaryDNA("ATCG"))
	fmt.Println(GetComplementaryDNA("GATTACA"))
	fmt.Println(GetComplementaryDNA("CGCG"))
}
