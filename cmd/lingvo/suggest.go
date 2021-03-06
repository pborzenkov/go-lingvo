package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/pborzenkov/go-lingvo"
)

func init() {
	commands["suggest"] = &command{
		name: "suggest <word>",
		desc: "get all word forms for <word>",
		fn:   suggest,
	}
}

func suggest(args []string) {
	fs := flag.NewFlagSet("suggest", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo suggest [flags] <word>")
		fmt.Fprintln(os.Stderr, "")
		fs.PrintDefaults()
	}

	from := lingvo.En
	to := lingvo.Ru
	fs.Var(&from, "from", "source language")
	fs.Var(&to, "to", "target language")
	fs.Parse(args)

	if fs.NArg() != 1 {
		exit(fs.Usage, "invalid number of arguments")
	}

	c := lingvo.NewClient(os.Getenv("LINGVO_API_KEY"))

	s, err := c.Suggest(context.Background(), fs.Arg(0), from, to)
	if err != nil {
		exit(nil, err)
	}

	printHeader(os.Stdout)
	for _, w := range s {
		fmt.Printf("%s\n", w)
	}
}
