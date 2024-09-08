package main

import (
	"flag"
	"fmt"

	"github.com/dkaman/cli"
)

var username string
var name string

func init() {
	{
		fs := flag.NewFlagSet("list", flag.ContinueOnError)
		fs.StringVar(&username, "username", "", "folder name")
		cli.Register("collections.folders", &cli.Cmd{
			Flags: fs,
		})
	}

	{
		cli.Register("collections.folders.list", &cli.Cmd{
			Run: Run,
		})
	}

	{
		fs := flag.NewFlagSet("list", flag.ContinueOnError)
		fs.StringVar(&name, "name", "", "folder name")
		cli.Register("collections.folders.create", &cli.Cmd{})
	}

}

func Run(args []string) error {
	fmt.Printf("listing folder: %s\n", args)
	return nil
}
