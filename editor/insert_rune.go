package editor

import "github.com/nsf/termbox-go"

func insertRune(event termbox.Event, textBuff [][]rune, currentRow int, currentCol *int) {
	line := make([]rune, len(textBuff[currentRow])+1)
	copy(line[:*currentCol], textBuff[currentRow][:*currentCol])
	if event.Key == termbox.KeySpace {
		line[*currentCol] = ' '
	} else if event.Key == termbox.KeyTab {
		line[*currentCol] = '	'
	} else {
		line[*currentCol] = event.Ch
	}
	copy(line[*currentCol+1:], textBuff[currentRow][*currentCol:])
	textBuff[currentRow] = line
	*currentCol++
}
