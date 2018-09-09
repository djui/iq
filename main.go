package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

type query struct {
	Section string
	Key     string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s SECTION.KEY FILE\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		haltOnError(errors.New("missing query argument"))
	} else if flag.NArg() == 1 {
		haltOnError(errors.New("missing file argument"))
	}

	queryArg := flag.Arg(0)
	q, err := parseQuery(queryArg)
	haltOnError(err)

	fileArg := flag.Arg(1)
	b, err := readFile(fileArg)
	haltOnError(err)

	i, err := ini.Load(b)
	haltOnError(err)

	fmt.Println(i.Section(q.Section).Key(q.Key).String())
}

func parseQuery(q string) (*query, error) {
	parts := strings.SplitN(q, ".", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, errors.New("invalid query argument")
	}

	return &query{
		Section: parts[0],
		Key:     parts[1],
	}, nil
}

func readFile(f string) ([]byte, error) {
	if f == "-" {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(f)
}

func haltOnError(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
