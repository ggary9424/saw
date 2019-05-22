package main

import (
	"os"

	"saw/cmd"
)

func main() {
	if err := cmd.SawCommand.Execute(); err != nil {
		os.Exit(0)
	}
}
