package editor

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
)

func displayStatusBar(m mode, sourceFile string, totalLines int, isModified bool, currentRow int, currentCol int, copyBuffLen int, undoBuffLen int, columns int, rows int) {
	var modeStatus string
	var copyStatus string
	var undoStatus string
	var fileStatus string
	var cursorStatus string

	if m == edit {
		modeStatus = "EDIT: "
	}

	if m == view {
		modeStatus = "VIEW: "
	}
	sourceLength := len(sourceFile)
	if sourceLength > 8 {
		sourceLength = 8
	}
	var mod string
	if isModified {
		mod = "Modified"
	} else {
		mod = "Saved"
	}
	fileStatus = fmt.Sprintf("%s - %d Lines %s", sourceFile[:sourceLength], totalLines, mod)
	cursorStatus = fmt.Sprintf("Row %d, Col %d ", currentRow+1, currentCol+1)

	if copyBuffLen > 0 {
		copyStatus = " [COPY]"
	}

	if undoBuffLen > 0 {
		undoStatus = " [UNDO]"
	}

	usedSpace := len(modeStatus) + len(fileStatus) + len(cursorStatus) + len(copyStatus) + len(undoStatus)
	freeSpace := strings.Repeat(" ", columns-usedSpace)
	status := modeStatus + fileStatus + copyStatus + undoStatus + freeSpace + cursorStatus
	printMessage(0, rows, termbox.ColorBlack, termbox.ColorWhite, status)
}
