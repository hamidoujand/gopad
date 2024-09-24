package editor

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string, textBuff [][]rune) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			textBuff = append(textBuff, []rune{})
			return textBuff, nil
		} else {
			return nil, fmt.Errorf("open file[%s]: %w", filename, err)
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		textBuff = append(textBuff, []rune{})

		for _, ch := range line {
			textBuff[lineNumber] = append(textBuff[lineNumber], ch)
		}
		lineNumber++
	}
	if lineNumber == 0 {
		textBuff = append(textBuff, []rune{})
	}
	return textBuff, nil
}
