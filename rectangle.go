package main

type Rectangle struct {
	pos    Vector2d
	width  int
	height int
}

func NewRectangle(x int, y int, width int, height int) Rectangle {
	return Rectangle{Vector2d{x, y}, width, height}
}
