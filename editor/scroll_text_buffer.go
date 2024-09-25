package editor

func scrollTextBuffer(currentRow int, offsetRow *int, currentCol int, offsetCol *int, totalRows int, totalCols int) {
	offsetR := *offsetRow
	offsetC := *offsetCol
	if currentRow < offsetR {
		*offsetRow = currentRow
	}

	if currentCol < offsetC {
		*offsetCol = currentCol
	}

	if currentRow >= offsetR+totalRows {
		*offsetRow = currentRow - totalRows + 1
	}

	if currentCol >= offsetC+totalCols {
		*offsetCol = currentCol - totalCols + 1
	}
}
