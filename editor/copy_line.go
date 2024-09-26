package editor

func copyLine(textBuff *[][]rune, copyBuff *[]rune, currentRow *int) {
	copiedLine := make([]rune, len((*textBuff)[*currentRow]))
	copy(copiedLine, (*textBuff)[*currentRow])
	*copyBuff = copiedLine
}
