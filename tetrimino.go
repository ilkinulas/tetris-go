package main

import (
	"math/rand"
	"strings"
	"unicode"
)

const (
	I = iota
	O
	T
	S
	Z
	J
	L
)

var pivots = [] Vector2d{{1, 1}, {1, 1}, {2, 1}, {2, 2}, {1, 2}, {2, 1}, {1, 1}}
var positions = [] Vector2d{{3, -1}, {3, -1}, {3, -1}, {3, -1}, {3, -1}, {3, 0}, {3, 0}}
var tetriminos = map[int]string{
	I: `....
	    IIII
	    ....
	    ....`,
	O: `....
	    .OO.
	    .OO.
	    ....`,
	T: `....
	    .TTT
	    ..T.
	    ....`,
	S: `....
	    ..SS
	    .SS.
	    ....`,
	Z: `....
	    ZZ..
	    .ZZ.
	    ....`,
	J: `..J.
	    ..J.
	    .JJ.
	    ....`,
	L: `.L..
	    .L..
	    .LL.
	    ....`,
}

func newTetrimino(tetriminoType int) *Tetrimino {
	cellData := tetriminos[tetriminoType]
	cellData = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, cellData)
	t := &Tetrimino{
		tetriminoType: tetriminoType,
		cells:         parseCellData(cellData, tetriminoType),
		position:      positions[tetriminoType],
		pivot:         pivots[tetriminoType],
	}
	t.calculateBounds()
	return t
}

func parseCellData(cellData string, tetriminoType int) [4][4]int {
	cells := [4][4]int{}
	for i := range cellData {
		x := i % 4
		y := i / 4
		if cellData[i] == '.' {
			cells[x][y] = 0
		} else {
			cells[x][y] = tetriminoType + 1
		}
	}
	return cells
}

type Tetrimino struct {
	tetriminoType int
	cells         [4][4]int
	position      Vector2d
	pivot         Vector2d
	bounds        Rectangle
}

func CreateRandomTetrimino() *Tetrimino {
	tetriminoType := rand.Intn(7)
	tetrimino := newTetrimino(tetriminoType)
	tetrimino.bounds = tetrimino.calculateBounds()
	return tetrimino
}

func (t *Tetrimino) blockPositions() [4]Vector2d {
	positions := [4]Vector2d{}
	index := 0
	t.forEachCell(func(x int, y int, cellValue int) {
		if cellValue > 0 {
			positions[index] = Vector2d{x, y}
			index++
		}
	})
	return positions
}

func (t *Tetrimino) forEachCell(f func(x int, y int, cellValue int)) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			f(x, y, t.cells[x][y])
		}
	}
}

func (t *Tetrimino) calculateRotatedBlockPositions() [4]Vector2d {
	if t.tetriminoType == O {
		return t.blockPositions()
	}
	positions := [4]Vector2d{}
	index := 0
	t.forEachCell(func(x int, y int, cellValue int) {
		if cellValue > 0 {
			relPosToPivot := Vector2d{x, y}.subtract(t.pivot)
			//TODO buraya aciklama ekle. rotation matris carpimi
			positions[index] = t.pivot.subtract(Vector2d{-relPosToPivot.y, relPosToPivot.x})
			index++
		}
	})

	collides := false
	for _, pos := range positions {
		if pos.y < 0 {
			collides = true
		}
	}
	if collides {
		for i := 0; i < 4; i++ {
			positions[i].y = positions[i].y + 1
		}
	}
	return positions
}

func (t *Tetrimino) applyRotation(rotatedPositions [4]Vector2d) {
	t.forEachCell(func(x int, y int, cellValue int) {
		t.cells[x][y] = 0
	})
	for _, pos := range rotatedPositions {
		t.cells[pos.x][pos.y] = t.tetriminoType + 1
	}
	t.bounds = t.calculateBounds()
}

func (t *Tetrimino) calculateBounds() Rectangle {
	minX := 1000
	minY := 1000
	maxX := 0
	maxY := 0
	t.forEachCell(func(x int, y int, cellValue int) {
		if cellValue > 0 {
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
	})

	return NewRectangle(minX, minY, maxX-minX+1, maxY-minY+1)
}

func (t *Tetrimino) MoveDown() {
	t.position.y = t.position.y + 1
}

func (t *Tetrimino) MoveLeft() {
	t.position.x = t.position.x - 1
}

func (t *Tetrimino) MoveRight() {
	t.position.x = t.position.x + 1
}

type CollisionFunction func(x int, y int) bool

func (t *Tetrimino) CheckCollision(collisionFunction CollisionFunction) bool {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if t.cells[x][y] > 0 {
				if collisionFunction(x, y) {
					return true
				}
			}
		}
	}
	return false
}
