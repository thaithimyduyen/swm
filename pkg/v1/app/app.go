package app

import (
	"fmt"
	"io"
)

// Run is an application entry-point.
func Run(args []string, out io.Writer) (err error) {
	_, err = fmt.Fprintln(out, "Hello Duyen")
	if err != nil {
		return fmt.Errorf("printing: %w", err)
	}

	return nil
}
