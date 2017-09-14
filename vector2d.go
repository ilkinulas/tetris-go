package main

type Vector2d struct {
	x int
	y int
}

func (v Vector2d) add(other Vector2d) Vector2d {
	return Vector2d{v.x + other.x, v.y + other.y}
}

func (v Vector2d) subtract(other Vector2d) Vector2d {
	return Vector2d{v.x - other.x, v.y - other.y}
}
