package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/pborzenkov/go-lingvo"
)

func init() {
	commands["get-article"] = &command{
		name: "get-article <word>",
		desc: "get article for <word>",
		fn:   getArticle,
	}
}

func getArticle(args []string) {
	fs := flag.NewFlagSet("get-article", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo get-article [flags] <word>")
		fmt.Fprintln(os.Stderr, "")
		fs.PrintDefaults()
	}

	from := lingvo.En
	to := lingvo.Ru
	fs.Var(&from, "from", "source language")
	fs.Var(&to, "to", "target language")
	dict := fs.String("dict", "Universal", "dictionary to get the article from")
	fs.Parse(args)

	if fs.NArg() != 1 {
		exit(fs.Usage, "invalid number of arguments")
	}

	c := lingvo.NewClient(os.Getenv("LINGVO_API_KEY"))

	a, err := c.GetArticle(context.Background(), fs.Arg(0), *dict, from, to)
	if err != nil {
		exit(nil, err)
	}

	printHeader(os.Stdout)
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

func extractText(n []*lingvo.ArticleNode) string {
	text := ""

	if n == nil {
		return text
	}

	for _, c := range n {
		if c.Markup != nil {
			text += extractText(c.Markup)
		}
		text += c.Text
	}
	return text
}
