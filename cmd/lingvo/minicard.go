package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pborzenkov/go-lingvo"
)

func init() {
	commands["get-minicard"] = &command{
		name: "get-minicard <word>",
		desc: "get minicard for <word>",
		fn:   getMinicard,
	}
}

func getMinicard(args []string) {
	fs := flag.NewFlagSet("get-minicard", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: lingvo get-minicard [flags] <word>")
		fmt.Fprintln(os.Stderr, "")
		fs.PrintDefaults()
	}

	from := lingvo.En
	to := lingvo.Ru
	fs.Var(&from, "from", "source language")
	fs.Var(&to, "to", "target language")
	sound := fs.String("sound", "", "store word's sound (if available) into the file")
	fs.Parse(args)

	if fs.NArg() != 1 {
		exit(fs.Usage, "invalid number of arguments")
	}

	c := lingvo.NewClient(os.Getenv("LINGVO_API_KEY"))

	m, err := c.GetMinicard(context.Background(), fs.Arg(0), from, to)
	if err != nil {
		exit(nil, err)
	}

	printHeader(os.Stdout)
	fmt.Printf("%s\n", m.Translation.Heading)
	fmt.Printf("%s\n", m.Translation.Translation)
	fmt.Printf("\nSee also: %s\n", strings.Join(m.SeeAlso, ","))

	if *sound == "" {
		return
	}

	if m.Translation.Sound == "" {
		fmt.Printf("\nNo sound available\n")
		return
	}

	snd, err := m.GetSound(context.Background())
	if err != nil {
		exit(nil, err)
	}
	err = ioutil.WriteFile(*sound, snd, 0644)
	if err != nil {
		exit(nil, err)
	}
	fmt.Printf("\nStored sound for '%s' as %s\n", fs.Arg(0), *sound)
}
