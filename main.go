package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type formater func(io.Reader, []string) interface{}

func main() {
	var format string
	flag.StringVar(&format, "format", "2d-array", "Output format to use (array, map, 2d-array)")
	flag.Parse()

	formats := map[string]formater{
		"array":    toArray,
		"2d-array": to2dArray,
	}

	f, ok := formats[format]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown format '%s'\n", format)
		return
	}

	out := f(os.Stdin, flag.Args())

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(out)
}

func toArray(r io.Reader, args []string) interface{} {
	sc := bufio.NewScanner(r)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func to2dArray(r io.Reader, args []string) interface{} {
	sc := bufio.NewScanner(r)
	lines := make([][]string, 0)
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		lines = append(lines, parts)
	}
	return lines
}
