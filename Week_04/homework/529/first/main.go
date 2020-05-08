package main

func main() {

}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	if board[click[0]][click[1]] != 'E' {
		return board
	}

	r, c := len(board), len(board[0])
	count := mineCount(board, click, r, c)

	if count > 0 {
		board[click[0]][click[1]] = byte(count + 48)
		return board
	}

	board[click[0]][click[1]] = 'B'
	for _, d := range directs() {
		if x, y := click[0]+d[0], click[1]+d[1]; x >= 0 && x < r && y >= 0 && y < c && board[x][y] == 'E' {
			updateBoard(board, []int{x, y})
		}
	}

	return board
}

func mineCount(board [][]byte, click []int, r, c int) int {
	count := 0
	for _, d := range directs() {
		if x, y := click[0]+d[0], click[1]+d[1]; x >= 0 && x < r && y >= 0 && y < c && (board[x][y] == 'M' || board[x][y] == 'X') {
			count++
		}
	}

	return count
}

func directs() [][]int {
	return [][]int{{1, -1}, {1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, 0}, {-1, 1}, {-1, -1}}
}
