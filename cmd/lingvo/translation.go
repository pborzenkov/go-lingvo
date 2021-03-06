package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/pborzenkov/go-lingvo"
)

func init() {
	commands["translate"] = &command{
		name: "translate <word>",
		desc: "translate <word>",
		fn:   translate,
	}
}

func translate(args []string) {
	fs := flag.NewFlagSet("translate", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo translate [flags] <word>")
		fmt.Fprintln(os.Stderr, "")
		fs.PrintDefaults()
	}

	from := lingvo.En
	to := lingvo.Ru
	fs.Var(&from, "from", "source language")
	fs.Var(&to, "to", "target language")
	isCaseSensitive := fs.Bool("case", false, "case sensitive search")
	fs.Parse(args)

	if fs.NArg() != 1 {
		exit(fs.Usage, "invalid number of arguments")
	}

	c := lingvo.NewClient(os.Getenv("LINGVO_API_KEY"))

	as, err := c.Translate(context.Background(), fs.Arg(0), from, to, *isCaseSensitive)
	if err != nil {
		exit(nil, err)
	}

	printHeader(os.Stdout)

	for _, a := range as {
		fmt.Printf("%s\n\n", a.Dictionary)
		fmt.Printf("%s\n", a.Title)
		for _, n := range a.Body {
			if n.Node != lingvo.List {
				continue
			}
			for i, item := range n.Items {
				if item.Markup != nil && item.Markup[0].Node == lingvo.Paragraph {
					fmt.Printf("%d) %s\n", i, extractText(item.Markup))
				}
			}
		}
		fmt.Printf("\n\n")
	}
}
