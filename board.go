package main

import (
	"math/rand"
	"time"
)

type Board struct {
	width         int
	height        int
	cells         [][]int
	tetrimino     *Tetrimino
	nextTetrimino *Tetrimino
}

func NewBoard(width int, height int) *Board {
	b := Board{width: width, height: height}
	array2d := make([][]int, width)
	for i := range array2d {
		array2d[i] = make([]int, height)
	}
	b.cells = array2d
	rand.Seed(time.Now().UnixNano())
	b.tetrimino = CreateRandomTetrimino()
	b.nextTetrimino = CreateRandomTetrimino()
	return &b
}

func (board *Board) Rotate() {
	newPositions := board.tetrimino.calculateRotatedBlockPositions()
	updateOutOfBoardTetriminoPos(board, newPositions)
	tPos := board.tetrimino.position

	collision := false
	for _, pos := range newPositions {
		if board.cells[tPos.x+pos.x][tPos.y+pos.y] > 0 {
			collision = true
			break
		}
	}
	if !collision {
		board.tetrimino.applyRotation(newPositions)
	}
}

func updateOutOfBoardTetriminoPos(board *Board, rotatedBlocks [4]Vector2d) {
	minX := 0
	maxX := board.width - 1
	minY := 0
	maxY := board.height - 1
	for i := 0; i < 4; i++ {
		x := rotatedBlocks[i].x + board.tetrimino.position.x
		y := rotatedBlocks[i].y + board.tetrimino.position.y
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	if minX < 0 {
		board.tetrimino.position.x = board.tetrimino.position.x - minX
	}
	if maxX > board.width-1 {
		board.tetrimino.position.x = board.tetrimino.position.x - (maxX - (board.width - 1))
	}
	if minY < 0 {
		board.tetrimino.position.y = board.tetrimino.position.y - minY
	}
	if maxY > board.height-1 {
		board.tetrimino.position.y = board.tetrimino.position.y - (maxY - (board.height - 1))
	}
}

func (board *Board) MoveDown() bool {
	collisionFunc := func(x int, y int) bool {
		targetX := board.tetrimino.position.x + x
		targetY := board.tetrimino.position.y + y + 1
		return targetY >= board.height || board.cells[targetX][targetY] > 0
	}
	collides := board.tetrimino.CheckCollision(collisionFunc)
	if collides {
		return false
	}
	board.tetrimino.MoveDown()
	return true
}

func (board *Board) MoveRight() bool {
	collisionFunc := func(x int, y int) bool {
		targetX := board.tetrimino.position.x + x + 1
		targetY := board.tetrimino.position.y + y
		return targetX >= board.width || board.cells[targetX][targetY] > 0
	}
	collides := board.tetrimino.CheckCollision(collisionFunc)
	if collides {
		return false
	}
	board.tetrimino.MoveRight()
	return true
}

func (board *Board) MoveLeft() bool {
	collisionFunc := func(x int, y int) bool {
		targetX := board.tetrimino.position.x + x - 1
		targetY := board.tetrimino.position.y + y
		return targetX < 0 || board.cells[targetX][targetY] > 0
	}
	collides := board.tetrimino.CheckCollision(collisionFunc)
	if collides {
		return false
	}
	board.tetrimino.MoveLeft()
	return true
}

func (board *Board) FallDown() {
	for board.MoveDown() {
	}
}

func (board *Board) StartNewTurn() int {
	t := board.tetrimino
	t.forEachCell(func(x int, y int, cellValue int) {
		if cellValue > 0 {
			board.cells[t.position.x+x][t.position.y+y] = cellValue
		}
	})

	board.tetrimino = board.nextTetrimino
	board.nextTetrimino = CreateRandomTetrimino()
	return board.clearLines()
}

func (board *Board) clearLines() int {
	fullLines := make([]int, 0)
	for y := 0; y < board.height; y++ {
		fullLine := true
		for x := 0; x < board.width; x++ {
			if board.cells[x][y] == 0 {
				fullLine = false
				break
			}
		}
		if fullLine {
			fullLines = append(fullLines, y)
		}
	}
	for _, line := range fullLines {
		for i := 0; i < board.width; i++ {
			board.cells[i][line] = 0
			for j := line; j > 0; j-- {
				board.cells[i][j] = board.cells[i][j-1]
				board.cells[i][j-1] = 0
			}
		}
	}

	return len(fullLines)
}

func (board *Board) isGameOver() bool {
	for x := 0; x < board.width; x++ {
		if board.cells[x][0] > 0 {
			return true
		}
	}
	return false
}
