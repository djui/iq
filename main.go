package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
	fileArg := flag.Arg(1)

	v, err := run(queryArg, fileArg)
	haltOnError(err)

	fmt.Println(v)
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
