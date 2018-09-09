package main

import (
	"errors"
	"strings"
)

type query struct {
	Section string
	Key     string
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
