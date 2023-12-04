package main

import (
	"os"

	"github.com/thaithimyduyen/swm/pkg/v1/app"
)

func main() {
	err := app.Run(os.Args, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
}
