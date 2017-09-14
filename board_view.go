package main

type BoardView struct {
	model   *Board
	xOffset int
	yOffset int
}

func (boardView BoardView) Render() {
	for row := 0; row < boardView.model.width; row++ {
		for col := 0; col < boardView.model.height; col++ {
			cellValue := boardView.model.cells[row][col]
			if cellValue > 0 {
				RenderCell(boardView.xOffset+row, boardView.yOffset+col, 1, 1, tetriminoColor(cellValue))
			}
		}
	}

	renderTetrimino(boardView.xOffset, boardView.yOffset, boardView.model.tetrimino)
}
