package main

import (
	"github.com/nsf/termbox-go"
)

type Game struct {
	boardWidth  int
	boardHeight int
	board       *Board
	boardView   *BoardView
	model       *GameModel
	frameCount  int
}

const (
	w = 10
	h = 20
)

func NewGame() *Game {
	game := Game{
		boardWidth:  w,
		boardHeight: h,
	}
	game.board = NewBoard(w, h)
	game.model = &GameModel{targetNumberOfLinesToLevelUp: linesPerLevel}
	game.boardView = &BoardView{game.board, 1, 1}
	return &game
}

func (game *Game) Restart() {
	game.board = NewBoard(w, h)
	game.model = &GameModel{}
	game.boardView = &BoardView{game.board, 1, 1}
}

func (game *Game) Render() {
	RenderCell(0, 0, game.boardWidth+2, game.boardHeight+2, BORDER_COLOR)
	RenderCell(1, 1, game.boardWidth, game.boardHeight, BOARD_BG_COLOR)

	game.boardView.Render()

	RenderScoreboard(game.boardWidth*2+5, 8, game.model)

	RenderNextTetrimino(game.boardWidth+2, 0, game.board.nextTetrimino)
	if game.model.IsGameOver {
		RenderGameOver()
	}
}

func (game *Game) Update() {
	game.frameCount++
	if game.frameCount >= framesPerBlock(game) {
		game.frameCount = 0
		if !game.model.IsGameOver {
			game.MoveDown()
		}
	}
}

func framesPerBlock(game *Game) int {
	frames := 15 - (game.model.Level * 2)
	if frames < 1 {
		frames = 1
	}
	return frames
}

func (game *Game) OnKeyPress(key termbox.Key) {
	if game.model.IsGameOver {
		game.Restart()
		return
	}

	switch key {
	case termbox.KeyArrowDown:
		game.MoveDown()
		break
	case termbox.KeyArrowUp:
		game.board.Rotate()
		break
	case termbox.KeyArrowRight:
		game.board.MoveRight()
		break
	case termbox.KeyArrowLeft:
		game.board.MoveLeft()
		break
	case termbox.KeySpace:
		game.board.FallDown()
		break
	}
}

func (game *Game) MoveDown() {
	moved := game.board.MoveDown()
	if !moved {
		clearedLines := game.board.StartNewTurn()
		game.model.updateGameModel(clearedLines)
		if game.board.isGameOver() {
			game.model.IsGameOver = true
		}
	}
}
