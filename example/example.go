package main

import (
	"fmt"
	"os"

	"github.com/dkaman/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Printf("failed to run cmd: %s\n", err)
		os.Exit(1)
	}
}
