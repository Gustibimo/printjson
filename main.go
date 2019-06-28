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
	flag.StringVar(&format, "format", "map", "Output format to use (array, map, 2d-array)")
	flag.Parse()

	formats := map[string]formater{
		"map": toMap,
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

func toMap(r io.Reader, args []string) interface{} {
	sc := bufio.NewScanner(r)
	lines := make([]map[string]string, 0)
	for sc.Scan() {
		line := make(map[string]string)
		fields := strings.Fields(sc.Text())

		for i, k := range args {
			if len(fields) <= i {
				break
			}
			// pass field that contain a dash
			if k == "-" {
				continue
			}
			line[k] = fields[i]
		}

		fmt.Print(line)

		lines = append(lines, line)
	}

	return lines
}
