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