package main

import "github.com/nsf/termbox-go"

const BORDER_COLOR = termbox.ColorWhite
const BOARD_BG_COLOR = termbox.ColorDefault

var TETRIMINO_COLORS = []termbox.Attribute{
	termbox.ColorGreen,
	termbox.ColorMagenta,
	termbox.ColorRed,
	termbox.ColorCyan,
	termbox.ColorBlue,
	termbox.ColorYellow}
