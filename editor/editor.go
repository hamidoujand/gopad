package editor

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func Run() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("init: %w", err)
	}

	var (
		rows       int
		cols       int
		offsetX    int
		offsetY    int
		sourceFile string
	)

	var textBuffer = [][]rune{}
	//fill the text buffer if any file provided.
	if len(os.Args) > 1 {
		sourceFile = os.Args[1]
		var err error
		textBuffer, err = readFile(sourceFile, textBuffer)
		if err != nil {
			return fmt.Errorf("read file: %w", err)
		}
	} else {
		sourceFile = "out.txt"
		//one line by default inside of buffer
		textBuffer = append(textBuffer, []rune{})
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
