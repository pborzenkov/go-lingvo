package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
)

type command struct {
	name string
	desc string
	fn   func(args []string)
}

var commands = make(map[string]*command)

func exit(usage func(), msg interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n\n", msg)
	if usage != nil {
		usage()
	}
	os.Exit(1)
}

func printHeader(w io.Writer) {
	fmt.Printf("Translated by Lingvo (https://developers.lingvolive.com)\n\n")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo <command> [flags]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Lingvo API key must be provided via LINGVO_API_KEY environment variable")
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Commands:")
		fmt.Fprintln(os.Stderr, "")

		cmds := make([]*command, 0, len(commands))
		for _, v := range commands {
			cmds = append(cmds, v)
		}
		sort.Slice(cmds, func(i, j int) bool {
			return cmds[i].name < cmds[j].name
		})

		w := tabwriter.NewWriter(os.Stderr, 0, 8, 0, '\t', 0)
		for _, cmd := range cmds {
			fmt.Fprintf(w, "  %s\t- %s\n", cmd.name, cmd.desc)
		}
		w.Flush()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		exit(flag.Usage, "<command> is required")
	}

	cmd := commands[flag.Arg(0)]
	if cmd == nil {
		exit(flag.Usage, fmt.Sprintf("unsupported command '%s'", flag.Arg(0)))
	}
	cmd.fn(flag.Args()[1:])
}
