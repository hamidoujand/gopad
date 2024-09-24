package editor

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func Run() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("init: %w", err)
	}

	//wait for user input
	for {
		printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "GOPAD - A bare bones text editor")
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
