/*
Snakes and Ladders is an ancient Indian board game that has become a classic worldwide. It is played by two or more players on a board consisting of numbered squares from 1 to 100. Some squares are connected by ladders, which move the player up the board, and snakes, which move the player down the board.

Your task is to create a class called SnakesLadders.

The test cases will repeatedly call the method: play(die1, die2).
	- die1 and die2 are integers between 1 and 6;
	- A player moves forward by the sum of die1 and die2;
	- The method is called independently for each turn and must manage the game state internally.

Rules:
	- There are two players;
	- Both players start on square 0;
	- Player 1 always starts;
	- Players take turns alternating;
	- If a player rolls a double, they take another turn;
	- The board is numbered from 1 to 100.

Ladders:
	- If a player lands exactly on the bottom of a ladder, they move up to the top of that ladder;
	- This happens even if the player rolled a double.

Snakes:
	- If a player lands exactly on the head of a snake, they slide down to the tail of the snake;
	- This also applies even if the player rolled a double.

Winning the Game:
	- A player must land exactly on square 100 to win;
	- If the roll would move the player past 100, they bounce back.
	- If a player rolls a double and lands on 100 with no remaining moves, they win immediately and do not take another turn.

Return Values:
	- If a player wins: "Player n Wins!";
	- If the game has already ended and another move is attempted: "Game over!";
	- Otherwise, after a valid move: "Player n is on square x". Where: n is the current player (1 or 2), x is the square the player is on.

«Змеи и лестницы» - это древняя индийская настольная игра, которая сегодня считается классикой во всём мире. В игре используется поле с пронумерованными клетками от 1 до 100. На некоторых клетках расположены лестницы, которые перемещают игрока вверх, и змеи, которые отправляют игрока вниз.

Необходимо создать класс SnakesLadders.

Тесты будут вызывать метод: play(die1, die2).
	- die1 и die2 - целые числа от 1 до 6;
	- Игрок перемещается на сумму значений двух кубиков;
	- Метод вызывается для каждого хода и должен самостоятельно хранить состояние игры.

Правила:
	- В игре участвуют два игрока;
	- Оба начинают на клетке 0;
	- Игрок 1 ходит первым;
	- Ходы выполняются по очереди;
	- Если выпал дубль, игрок получает дополнительный ход;
	- Игровое поле нумеруется от 1 до 100.

Лестницы:
	- Если игрок попадает точно на начало лестницы, он сразу поднимается на её верх;
	- Это происходит даже если выпал дубль.

Змеи:
	- Если игрок попадает точно на голову змеи, он скатывается вниз к её хвосту;
	- Правило действует даже при дубле.

Победа:
	- Для победы необходимо попасть точно на клетку 100;
	- Если выпало слишком большое число, игрок отскакивает назад.

Если игрок выбрасывает дубль и попадает на 100 без лишних шагов, он выигрывает сразу и не ходит снова.

Возвращаемые значения:
	- Если игрок выиграл: "Player n Wins!";
	- Если игра уже закончена и совершается новый ход: "Game over!".

В остальных случаях: "Player n is on square x". Где: n - номер текущего игрока (1 или 2), x - номер клетки, на которой он находится.
*/

package main

import "fmt"

type SnakesLadders struct {
	playerPositions  [2]int
	currentPlayer    int
	isGameOver       bool
	boardTransitions map[int]int
}

func (game *SnakesLadders) NewGame() {
	game.playerPositions = [2]int{0, 0}
	game.currentPlayer = 0
	game.isGameOver = false

	game.boardTransitions = map[int]int{
		2:  38,
		7:  14,
		8:  31,
		15: 26,
		21: 42,
		28: 84,
		36: 44,
		51: 67,
		71: 91,
		78: 98,
		87: 94,

		16: 6,
		46: 25,
		49: 11,
		62: 19,
		64: 60,
		74: 53,
		89: 68,
		92: 88,
		95: 75,
		99: 80,
	}
}

func (game *SnakesLadders) Play(firstDie int, secondDie int) string {
	if game.isGameOver {
		return "Game over!"
	}

	playingPlayer := game.currentPlayer
	rollSum := firstDie + secondDie
	newPosition := game.playerPositions[playingPlayer] + rollSum

	if newPosition > 100 {
		newPosition = 100 - (newPosition - 100)
	}

	if destination, exists := game.boardTransitions[newPosition]; exists {
		newPosition = destination
	}

	game.playerPositions[playingPlayer] = newPosition

	if newPosition == 100 {
		game.isGameOver = true
		return fmt.Sprintf("Player %d Wins!", playingPlayer+1)
	}

	result := fmt.Sprintf(
		"Player %d is on square %d",
		playingPlayer+1,
		newPosition,
	)

	if firstDie != secondDie {
		game.currentPlayer = 1 - game.currentPlayer
	}

	return result
}