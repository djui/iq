package main

import "github.com/go-ini/ini"

func run(query, file string) (string, error) {
	q, err := parseQuery(query)
	if err != nil {
		return "", err
	}

	b, err := readFile(file)
	if err != nil {
		return "", err
	}

	return runQuery(q, b)
}

func runQuery(q *query, b []byte) (string, error) {
	i, err := ini.Load(b)
	if err != nil {
		return "", err
	}

	return lookup(i, q, ""), nil
}

func lookup(i *ini.File, q *query, fallback string) string {
	s := i.Section(q.Section)
	if s == nil {
		return fallback
	}

	v := s.Key(q.Key)
	if v == nil {
		return fallback
	}

	return v.String()
}
