package editor

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

func printMessage(col int, row int, fg termbox.Attribute, bg termbox.Attribute, msg string) {
	for _, ch := range msg {
		termbox.SetCell(col, row, ch, fg, bg)
		//move col to after number of bytes in rune
		col += utf8.RuneLen(ch)
	}
}
