package main

import (
	"fmt"
	"os"

	"github.com/hamidoujand/gopad/editor"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if err := editor.Run(); err != nil {
		return fmt.Errorf("run: %w", err)
	}
	return nil
}
