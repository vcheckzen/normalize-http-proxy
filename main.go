package main

import (
	"os"

	"github.com/vcheckzen/normalize-http-proxy/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
