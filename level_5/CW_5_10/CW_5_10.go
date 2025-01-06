/*
Write a function first_non_repeating_letter that takes a string as input and returns the first character that does not repeat anywhere in the string.

Rules:
Upper- and lowercase letters are treated as the same character, but the function should return the character in its original case.
If all characters in the string are repeated, return an empty string ("").
The function should handle any Unicode characters.

Examples:
Input: 'stress' → Output: 't' (only 't' is non-repeating)
Input: 'sTreSS' → Output: 'T' ('T' is non-repeating, preserving its case)
Input: 'aabbcc' → Output: "" (all characters are repeated)

Напишите функцию first_non_repeating_letter, которая принимает строку и возвращает первый символ, который нигде в строке не повторяется.

Правила:
Заглавные и строчные буквы считаются одним и тем же символом, но функция должна возвращать символ в его исходном регистре.
Если все символы в строке повторяются, вернуть пустую строку ("").
Функция должна обрабатывать любые символы Юникода.

Примеры:
Ввод: 'stress' → Вывод: 't' (только 't' не повторяется)
Ввод: 'sTreSS' → Вывод: 'T' ('T' не повторяется, сохраняется регистр)
Ввод: 'aabbcc' → Вывод: "" (все символы повторяются)

https://www.codewars.com/kata/52bc74d4ac05d0945d00054e
*/