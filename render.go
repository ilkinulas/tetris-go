package main

import (
	"strconv"
	"github.com/nsf/termbox-go"
)

func renderTetrimino(xOffset int, yOffset int, t *Tetrimino) {
	t.forEachCell(func(x int, y int, value int) {
		if value > 0 {
			RenderCell(xOffset+t.position.x+x, yOffset+t.position.y+y, 1, 1, tetriminoColor(value))
		}
	})
}

func RenderNextTetrimino(xOffset int, yOffset int, t *Tetrimino) {
	RenderBorder(xOffset+13, yOffset, 6, 6)
	RenderText("NEXT", xOffset+15, yOffset)
	t.forEachCell(func(x int, y int, value int) {
		if value > 0 {
			RenderCell(xOffset+x+1, yOffset+y+1, 1, 1, termbox.ColorWhite)
		}
	})
}

func tetriminoColor(value int) termbox.Attribute {
	return TETRIMINO_COLORS[value%len(TETRIMINO_COLORS)]
}

func RenderScoreboard(x, y int, model *GameModel) {
	RenderBorder(x, y, 6, 3)
	RenderText("SCORE", x+3, y)
	RenderText(strconv.Itoa(model.Score), x+3, y+1)

	RenderBorder(x, y+3, 6, 3)
	RenderText("LEVEL", x+3, y+3)
	RenderText(strconv.Itoa(model.Level), x+3, y+4)

	RenderBorder(x, y+6, 6, 3)
	RenderText("LINES", x+3, y+6)
	RenderText(strconv.Itoa(model.NumLines), x+3, y+7)
}

func fillRect(x, y, w, h int, ch rune, fgColor termbox.Attribute, bgColor termbox.Attribute) {
	realX := x * 2
	for i := 0; i < w*2; i = i + 2 {
		for j := 0; j < h; j++ {
			termbox.SetCell(realX+i, y+j, ch, fgColor, bgColor)
			termbox.SetCell(realX+i+1, y+j, ch, fgColor, bgColor)
		}
	}
}

func RenderCell(x, y, w, h int, color termbox.Attribute) {
	fillRect(x, y, w, h, ' ', color, color)
}

func RenderBorder(x, y, w, h int) {
	w = w * 2
	termbox.SetCell(x, y, '┌', termbox.ColorDefault, termbox.ColorDefault)
	fill(x+1, y, w-2, 1, '─')
	termbox.SetCell(x+w-1, y, '┐', termbox.ColorDefault, termbox.ColorDefault)
	fill(x, y+1, 1, h-2, '│')
	fill(x+w-1, y+1, 1, h-2, '│')
	termbox.SetCell(x, y+h-1, '└', termbox.ColorDefault, termbox.ColorDefault)
	fill(x+1, y+h-1, w-2, 1, '─')
	termbox.SetCell(x+w-1, y+h-1, '┘', termbox.ColorDefault, termbox.ColorDefault)
}

func fill(x, y, w, h int, r rune) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			termbox.SetCell(x+i, y+j, r, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func RenderText(text string, x, y int) {
	for index, ch := range text {
		termbox.SetCell(x+index, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func RenderGameOver() {
	fillRect(1, 6, 10, 8, ' ', termbox.ColorWhite, termbox.ColorWhite)
	RenderText("                ", 4, 7)
	RenderText("   GAME OVER    ", 4, 8)
	RenderText("                ", 4, 9)
	RenderText(" press any key  ", 4, 10)
	RenderText(" to continue    ", 4, 11)
	RenderText("                ", 4, 12)

}
