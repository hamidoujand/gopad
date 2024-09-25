package editor

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func getKey() (termbox.Event, error) {
	var keyEvent termbox.Event
	event := termbox.PollEvent()
	switch event.Type {
	case termbox.EventKey:
		keyEvent = event
	case termbox.EventError:
		return termbox.Event{}, fmt.Errorf("pollEvent: %w", event.Err)
	}

	return keyEvent, nil
}
