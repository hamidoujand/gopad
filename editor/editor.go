package editor

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func Run() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("init: %w", err)
	}

	var rows int
	var cols int
	var offsetX int
	var offsetY int

	var textBuffer = [][]rune{
		[]rune("Hello World!"),
	}

	//wait for user input
	for {
		// printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "GOPAD - A bare bones text editor")
		cols, rows = termbox.Size()
		rows-- // one line always for status bar down.

		if cols < 78 {
			cols = 78
		}

		if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
			return fmt.Errorf("clear: %w", err)
		}

		//diplay text buffer buffer
		displayTextBuffer(textBuffer, rows, cols, offsetX, offsetY)

		if err := termbox.Flush(); err != nil {
			return fmt.Errorf("flush: %w", err)
		}
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
			termbox.Close()
			break
		}
	}
	return nil
}
