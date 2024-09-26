package editor

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

type mode int

const (
	view mode = 1
	edit mode = 2
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
		currentX   int
		currentY   int
		modified   bool
	)

	var (
		textBuffer = [][]rune{}
		undoBuffer = [][]rune{}
		copyBuffer = []rune{} //single line
	)

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
	mod := view
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

		//scroll
		scrollTextBuffer(currentY, &offsetY, currentX, &offsetX, rows, cols)
		//diplay text buffer buffer
		displayTextBuffer(textBuffer, rows, cols, offsetX, offsetY)
		displayStatusBar(mod, sourceFile, len(textBuffer), modified, currentY, currentX, len(copyBuffer), len(undoBuffer), cols, rows)

		//current cursor pos
		termbox.SetCursor(currentX-offsetX, currentY-offsetY)
		if err := termbox.Flush(); err != nil {
			return fmt.Errorf("flush: %w", err)
		}

		if err := processKeyPress(&currentY, &currentX, &textBuffer, rows, &modified, &mod, sourceFile, &copyBuffer, &undoBuffer); err != nil {
			return fmt.Errorf("processKey: %w", err)
		}
	}
}
