/*
You are given:
	- a two-dimensional array arr containing non-negative integers;
	- a target value val;
	- a starting coordinate coord in the form.

Your task is to return an iterable (depending on the programming language) containing all coordinates that:
	- contain the value val;
	- are connected to the starting coordinate through cells with the same value.

A coordinate is considered connected if there exists a path from the starting coordinate using only cells with value val, where each step moves:
	- horizontally;
	- vertically;
	- or diagonally.

Вам даны:
	- двумерный массив arr, содержащий неотрицательные целые числа;
	- искомое значение val;
	- начальная координата coord в формате.

Необходимо вернуть итерируемую коллекцию, содержащую все координаты, которые:
	- содержат значение val;
	- связаны с начальной координатой через клетки с тем же значением.

Координаты считаются связанными, если существует путь от начальной клетки, проходящий только через клетки со значением val, при этом разрешены переходы:
	- по горизонтали;
	- по вертикали;
	- по диагонали.
*/

package kata

func ConnectedValues(arr [][]int, val int, coord [2]int) [][2]int {
	result := make([][2]int, 0)

	rows := len(arr)
	if rows == 0 {
		return result
	}
	columns := len(arr[0])

	startRow := coord[0]
	startColumn := coord[1]

	if startRow < 0 || startRow >= rows || startColumn < 0 || startColumn >= columns {
		return result
	}

	if arr[startRow][startColumn] != val {
		return result
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	queue := make([][2]int, 0)
	queue = append(queue, coord)
	visited[startRow][startColumn] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentRow := current[0]
		currentColumn := current[1]

		result = append(result, current)

		for _, direction := range directions {
			newRow := currentRow + direction[0]
			newColumn := currentColumn + direction[1]

			if newRow >= 0 && newRow < rows &&
				newColumn >= 0 && newColumn < columns &&
				!visited[newRow][newColumn] &&
				arr[newRow][newColumn] == val {

				visited[newRow][newColumn] = true
				queue = append(queue, [2]int{newRow, newColumn})
			}
		}
	}

	return result
}