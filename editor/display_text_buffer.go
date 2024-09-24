package editor

import "github.com/nsf/termbox-go"

func displayTextBuffer(buff [][]rune, maxRows int, maxCols int, offsetX int, offsetY int) {
	var col int
	var row int

	for row = 0; row < maxRows; row++ {
		textBufferRow := row + offsetY
		for col = 0; col < maxCols; col++ {
			textBufferCol := col + offsetX
			if textBufferRow >= 0 && textBufferRow < len(buff) &&
				textBufferCol < len(buff[textBufferRow]) {
				if buff[textBufferRow][textBufferCol] != '\t' {
					termbox.SetChar(col, row, buff[textBufferRow][textBufferCol])
				} else {
					termbox.SetCell(col, row, ' ', termbox.ColorDefault, termbox.ColorGreen)
				}
			} else if row+offsetY > len(buff)-1 {
				//print * when file empty
				termbox.SetCell(0, row, '*', termbox.ColorBlue, termbox.ColorDefault)
			}
		}
		termbox.SetChar(col, row, '\n')
	}
}
