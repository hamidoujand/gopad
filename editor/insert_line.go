package editor

func insertLine(textBuff *[][]rune, currentCol *int, currentRow *int) {
	rightLine := make([]rune, len((*textBuff)[*currentRow][*currentCol:]))
	copy(rightLine, (*textBuff)[*currentRow][*currentCol:])
	leftLine := make([]rune, len((*textBuff)[*currentRow][:*currentCol]))
	copy(leftLine, (*textBuff)[*currentRow][:*currentCol])
	(*textBuff)[*currentRow] = leftLine
	*currentRow++
	*currentCol = 0
	newTextBuff := make([][]rune, len((*textBuff))+1)
	copy(newTextBuff, (*textBuff)[:*currentRow])
	newTextBuff[*currentRow] = rightLine
	copy(newTextBuff[*currentRow+1:], (*textBuff)[*currentRow:])
	*textBuff = newTextBuff
}
