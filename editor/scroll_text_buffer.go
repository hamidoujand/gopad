package editor

func scrollTextBuffer(currentRow int, offsetRow *int, currentCol int, offsetCol *int, totalRows int, totalCols int) {
	if currentRow < *offsetRow {
		*offsetRow = currentRow
	}

	if currentCol < *offsetCol {
		*offsetCol = currentCol
	}

	if currentRow >= *offsetRow+totalRows {
		*offsetRow = currentRow - totalRows + 1
	}

	if currentCol >= *offsetCol+totalCols {
		*offsetCol = currentCol - totalCols + 1
	}
}
