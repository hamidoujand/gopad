package editor

func pushBuffer(textBuff *[][]rune, undoBuff *[][]rune) {
	copiedUndoBuff := make([][]rune, len(*textBuff))
	copy(copiedUndoBuff, *textBuff)
	*undoBuff = copiedUndoBuff
}

func pullBuffer(undoBuff *[][]rune, textBuff *[][]rune) {
	if len(*undoBuff) == 0 {
		return
	}
	*textBuff = *undoBuff
}
