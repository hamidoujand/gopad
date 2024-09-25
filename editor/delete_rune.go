package editor

func deleteRune(textBuff *[][]rune, currentRow *int, currentCol *int) {
	if *currentCol > 0 {
		*currentCol--
		delete_line := make([]rune, len((*textBuff)[*currentRow])-1)
		copy(delete_line[:*currentCol], (*textBuff)[*currentRow][:*currentCol])
		copy(delete_line[*currentCol:], (*textBuff)[*currentRow][*currentCol+1:])
		(*textBuff)[*currentRow] = delete_line
	} else if *currentRow > 0 {
		append_line := make([]rune, len((*textBuff)[*currentRow]))
		copy(append_line, (*textBuff)[*currentRow][*currentCol:])
		newTextBuff := make([][]rune, len((*textBuff))-1)
		copy(newTextBuff[:*currentRow], (*textBuff)[:*currentRow])
		copy(newTextBuff[*currentRow:], (*textBuff)[*currentRow+1:])
		(*textBuff) = newTextBuff
		*currentRow--
		*currentCol = len((*textBuff)[*currentRow])
		insert_line := make([]rune, len((*textBuff)[*currentRow])+len(append_line))
		copy(insert_line[:len((*textBuff)[*currentRow])], (*textBuff)[*currentRow])
		copy(insert_line[len((*textBuff)[*currentRow]):], append_line)
		(*textBuff)[*currentRow] = insert_line
	}
}
