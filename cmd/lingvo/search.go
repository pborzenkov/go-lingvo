package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/pborzenkov/go-lingvo"
)

func init() {
	commands["search"] = &command{
		name: "search <word>",
		desc: "search dictionaries for <word>",
		fn:   search,
	}
}

func search(args []string) {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo search [flags] <word>")
		fmt.Fprintln(os.Stderr, "")
		fs.PrintDefaults()
	}

	from := lingvo.En
	to := lingvo.Ru
	zone := lingvo.SearchAuto
	fs.Var(&from, "from", "source language")
	fs.Var(&to, "to", "target language")
	fs.Var(&zone, "zone", "search zone")
	start := fs.Int("start", 0, "starting position")
	pageSize := fs.Int("pageSize", 1, "page size")
	fs.Parse(args)

	if fs.NArg() != 1 {
		exit(fs.Usage, "invalid number of arguments")
	}

	c := lingvo.NewClient(os.Getenv("LINGVO_API_KEY"))

	r, err := c.Search(context.Background(), fs.Arg(0), from, to, zone, *start, *pageSize)
	if err != nil {
		exit(nil, err)
	}

	printHeader(os.Stdout)
	for _, a := range r.Items {
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
	}
}
