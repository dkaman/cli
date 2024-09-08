package cli

import (
	"flag"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

type Cmd struct {
	Flags *flag.FlagSet
	Usage string
	Run   func([]string) error

	children map[string]*Cmd
}

var root = &Cmd{}

func Register(name string, cmd *Cmd) {
	path := strings.Split(name, ".")
	node := root
	for idx, el := range path {
		if node.children == nil {
			node.children = make(map[string]*Cmd)
		}
		if idx == len(path)-1 {
			node.children[el] = cmd
			break
		}
		if c, ok := node.children[el]; ok {
			node = c
		} else {
			c := &Cmd{}
			node.children[el] = c
			node = c
		}
	}
}

func Run(args []string) (err error) {
	var ok bool
	var subcmd string
	node := root
	for {
		if len(args) == 0 {
			return fmt.Errorf("subcommands: %s", maps.Keys(node.children))
		}
		subcmd, args = args[0], args[1:]
		node, ok = node.children[subcmd]
		if !ok {
			return fmt.Errorf("unknown subcommand: %s", subcmd)
		}
		if node.Flags != nil {
			node.Flags.Parse(args)
			args = node.Flags.Args()
		}
		if len(node.children) > 0 {
			continue
		}
		if node.Run == nil {
			return fmt.Errorf("misconfigured subcommand")
		}
		return node.Run(args)
	}
}
