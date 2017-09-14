package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	keyboardEvents := make(chan termbox.Key)
	go handleKeyPress(keyboardEvents)

	game := NewGame()

gameLoop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		game.Update()
		game.Render()
		termbox.Flush()

		select {
		case key := <-keyboardEvents:
			switch key {
			case termbox.KeyEsc:
				break gameLoop
			default:
				game.OnKeyPress(key)

			}
		default:
			time.Sleep(40 * time.Millisecond)
			continue
		}
	}
}

func handleKeyPress(keyboardEvents chan termbox.Key) {
	for {
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			keyboardEvents <- event.Key
		}
	}
}
