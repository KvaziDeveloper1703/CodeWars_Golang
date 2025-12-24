/*
The rules of the Rock, Paper, Scissors game are as follows:
	- Rock beats Scissors;
	- Scissors beat Paper;
	- Paper beats Rock;
	- If both players choose the same move, the result is a draw.

You will be given the valid moves of two players.

Your task is to determine the outcome of the game and return:
	- "Player 1 won!" if player 1 wins;
	- "Player 2 won!" if player 2 wins;
	- "Draw!" if the game ends in a draw.

Правила игры «Камень, ножницы, бумага»:
	- Камень побеждает ножницы;
	- Ножницы побеждают бумагу;
	- Бумага побеждает камень;
	- Если оба игрока выбирают одинаковый ход - ничья.

Вам будут переданы корректные ходы двух игроков.

Необходимо определить результат игры и вернуть:
	- "Player 1 won!", если победил первый игрок;
	- "Player 2 won!", если победил второй игрок;
	- "Draw!", если результат - ничья.
*/

package main

func Rps(playerOneMove, playerTwoMove string) string {
	if playerOneMove == playerTwoMove {
		return "Draw!"
	}

	if (playerOneMove == "rock" && playerTwoMove == "scissors") || (playerOneMove == "scissors" && playerTwoMove == "paper") || (playerOneMove == "paper" && playerTwoMove == "rock") {
		return "Player 1 won!"
	}

	return "Player 2 won!"
}