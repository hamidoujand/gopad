package editor

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string, textBuff *[][]rune, isModified *bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for row, line := range *textBuff {
		newLine := "\n"
		if row == len(*textBuff)-1 {
			newLine = ""
		}
		if _, err := writer.WriteString(string(line)); err != nil {
			return fmt.Errorf("writeString: %w", err)
		}
		writer.WriteString(newLine)
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush: %w", err)
	}
	*isModified = false

	return nil
}
