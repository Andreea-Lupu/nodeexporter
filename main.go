package main

import (
	"os"
	
	"github.com/anuvu/nodeexporter/cli"
)

func main() {
	if err := cli.NewZotExporterCmd().Execute(); err != nil {
		os.Exit(1)
	}
}