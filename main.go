package main

import (
	"log"
	"github.com/volker48/touchstone/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
