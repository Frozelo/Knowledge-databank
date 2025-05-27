// Given an m x n matrix board where each cell is a battleship 'X' or empty '.', return the number of the battleships on board.
// Battleships can only be placed horizontally or vertically on board. 
// In other words, they can only be made of the shape 1 x k (1 row, k columns) or k x 1 (k rows, 1 column), where k can be of any size. At least one horizontal or vertical cell separates between two battleships 
// (i.e., there are no adjacent battleships).

// Input: board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
// Output: 2

// https://leetcode.com/problems/battleships-in-a-board/description

func countBattleships(board [][]byte) int {
    var cnt int
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == '.' {
                continue
            }

            if i > 0 && board[i - 1][j] == 'X' {
                continue
            }

            if j > 0 && board[i][j -1] == 'X' {
                continue
            }

            cnt++
        }
    }

    return cnt
}
