package main

import (
	"fmt"
	"os"

	"github.com/dlbarduzzi/sentinel"
)

func main() {
	app := sentinel.New()

	if err := app.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s\n", err)
		os.Exit(1)
	}
}
