package main

import (
	"github.com/TualatinX/blockchain-go/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
