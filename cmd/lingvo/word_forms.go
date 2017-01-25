package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"golang.org/x/net/context"

	"github.com/pborzenkov/go-lingvo"
)

func init() {
	commands["get-word-forms"] = &command{
		name: "get-word-forms <word>",
		desc: "get all word forms for <word>",
		fn:   getWordForms,
	}
}

func getWordForms(args []string) {
	fs := flag.NewFlagSet("get-word-forms", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo get-word-forms [flags] <word>")
		fmt.Fprintln(os.Stderr, "")
		fs.PrintDefaults()
	}

	var lang langFlag = langFlag(lingvo.En)
	fs.Var(&lang, "lang", "language of the requested word")
	fs.Parse(args)

	if fs.NArg() != 1 {
		exit(fs.Usage, "invalid number of arguments")
	}

	c := lingvo.NewClient(os.Getenv("LINGVO_API_KEY"))

	wf, err := c.GetWordForms(context.Background(), fs.Arg(0), lingvo.Lang(lang))
	if err != nil {
		exit(nil, err)
	}

	printHeader(os.Stdout)
	for _, f := range wf {
		w := tabwriter.NewWriter(os.Stdout, 10, 8, 2, '\t', 0)
		fmt.Fprintf(w, "%s; %s\n", f.Lexeme, f.PartOfSpeech)
		for _, g := range f.Paradigm.Groups {
			fmt.Fprintf(w, "\n")
			if g.Name != "" {
				fmt.Fprintf(w, "--- %s ---\n", g.Name)
			}
			for _, r := range g.Table {
				for _, c := range r {
					fmt.Fprintf(w, "%s%s\t", c.Prefix, c.Value)
				}
				fmt.Fprintf(w, "\n")
			}
		}
		fmt.Fprintf(w, "\n\n")
		w.Flush()
	}
}
