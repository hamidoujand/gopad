package editor

func cutLine(textBuff *[][]rune, copyBuff *[]rune, currentRow *int, currentCol *int) {
	copyLine(textBuff, copyBuff, currentRow)
	if *currentRow >= len(*textBuff) || len(*textBuff) < 2 {
		return
	}

	nexTextBuff := make([][]rune, len(*textBuff)-1)
	copy(nexTextBuff[:*currentRow], (*textBuff)[:*currentRow])
	copy(nexTextBuff[*currentRow:], (*textBuff)[*currentRow+1:])
	*textBuff = nexTextBuff

	if *currentRow > 0 {
		*currentRow--
		*currentCol = 0
	}
}
