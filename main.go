package main

import (
	"os"

	"github.com/caioeverest/vulcan/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
