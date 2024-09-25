package editor

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func processKeyPress(currentRow int, textBuffLen int) (int, error) {
	keyEvent, err := getKey()
	if err != nil {
		return 0, fmt.Errorf("getKey: %w", err)
	}

	if keyEvent.Key == termbox.KeyEsc {
		termbox.Close()
		os.Exit(0)

	} else if keyEvent.Ch != 0 {
		//handle chars
	} else {
		switch keyEvent.Key {
		case termbox.KeyArrowUp:
			if currentRow != 0 {
				currentRow--
				return currentRow, nil
			}
		case termbox.KeyArrowDown:
			if currentRow < textBuffLen-1 {
				currentRow++
				return currentRow, nil
			}
		}
	}

	return currentRow, nil
}
