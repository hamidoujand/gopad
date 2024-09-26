package editor

func pasteLine(copyBuff *[]rune, textBuff *[][]rune, currentRow *int, currentCol *int) {
	if len(*copyBuff) == 0 {
		*currentRow++
		*currentCol = 0
	}

	newTextBuff := make([][]rune, len((*textBuff))+1)
	copy(newTextBuff[:*currentRow], (*textBuff)[:*currentRow])
	newTextBuff[*currentRow] = *copyBuff
	copy(newTextBuff[*currentRow+1:], (*textBuff)[*currentRow:])
	*textBuff = newTextBuff
}
