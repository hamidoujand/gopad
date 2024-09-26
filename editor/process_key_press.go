package editor

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func processKeyPress(currentRow *int, currentCol *int, textBuff *[][]rune, totalRows int, isModified *bool, currentMode *mode, filename string) error {
	keyEvent, err := getKey()
	if err != nil {
		return fmt.Errorf("getKey: %w", err)
	}

	if keyEvent.Key == termbox.KeyEsc {
		*currentMode = view

	} else if keyEvent.Ch != 0 {
		//handle chars
		if *currentMode == edit {
			insertRune(keyEvent, *textBuff, *currentRow, currentCol)
			*isModified = true
		} else {
			switch keyEvent.Ch {
			case 'q':
				termbox.Close()
				os.Exit(0)
			case 'e':
				*currentMode = edit
			case 'w':
				if err := writeFile(filename, textBuff, isModified); err != nil {
					return fmt.Errorf("writefile: %w", err)
				}
			}
		}
	} else {
		switch keyEvent.Key {
		case termbox.KeyEnter:
			if *currentMode == edit {
				insertLine(textBuff, currentCol, currentRow)
				*isModified = true
			}
		case termbox.KeyTab:
			if *currentMode == edit {
				insertRune(keyEvent, *textBuff, *currentRow, currentCol)
				*isModified = true
			}
		case termbox.KeySpace:
			if *currentMode == edit {
				insertRune(keyEvent, *textBuff, *currentRow, currentCol)
				*isModified = true
			}
		case termbox.KeyBackspace:
			if *currentMode == edit {
				deleteRune(textBuff, currentRow, currentCol)
				*isModified = true
			}
		case termbox.KeyBackspace2:
			if *currentMode == edit {
				deleteRune(textBuff, currentRow, currentCol)
				*isModified = true
			}
		case termbox.KeyHome:
			*currentCol = 0
		case termbox.KeyEnd:
			*currentCol = len((*textBuff)[*currentRow])
		case termbox.KeyPgup:
			if *currentRow-(totalRows/4) > 0 {
				*currentRow -= (totalRows / 4)
			}
		case termbox.KeyPgdn:
			if *currentRow+(totalRows/4) < len((*textBuff))-1 {
				*currentRow += (totalRows / 4)
			}
		case termbox.KeyArrowUp:
			if *currentRow != 0 {
				*currentRow--
			}
		case termbox.KeyArrowDown:
			if *currentRow < len((*textBuff))-1 {
				*currentRow++
			}
		case termbox.KeyArrowLeft:
			if *currentCol != 0 {
				*currentCol--
			} else if *currentRow > 0 {
				*currentRow--
				*currentCol = len((*textBuff)[*currentRow])
			}
		case termbox.KeyArrowRight:
			if *currentCol < len((*textBuff)[*currentRow]) {
				*currentCol++
			} else if *currentRow < len((*textBuff))-1 {
				*currentRow++
				*currentCol = 0
			}
		}
		if *currentCol > len((*textBuff)[*currentRow]) {
			*currentCol = len((*textBuff)[*currentRow])
		}
	}

	return nil
}
