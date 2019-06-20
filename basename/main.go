package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	usage = "strip directory and suffix from filenames"
)

var multiple bool
var suffix string
var zero bool

func init() {
	const (
		a_usage = "support multiple arguments and treat each as a NAME"
		s_usage = "remove a trailing `SUFFIX`; implies -a"
		z_usage = "end each output line with NUL, not newline"
	)
	flag.BoolVar(&multiple, "a", false, a_usage)
	flag.BoolVar(&multiple, "multiple", false, a_usage)
	flag.StringVar(&suffix, "s", "", s_usage)
	flag.StringVar(&suffix, "suffix", "", s_usage)
	flag.BoolVar(&zero, "z", false, z_usage)
	flag.BoolVar(&zero, "zero", false, z_usage)
}

func main() {
	// Note: -s implies -a
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "missing operand\n")
		os.Exit(1)
	}
	if suffix != "" {
		multiple = true
	}
	if flag.NArg() == 2 && suffix == "" {
		suffix = flag.Arg(1)
	}
	if flag.NArg() > 2 && suffix == "" && !multiple {
		fmt.Fprintf(os.Stderr, "extra operand: %s\n", flag.Arg(2))
		os.Exit(1)
	}

	for _, arg := range flag.Args() {
		var newline string
		var basename string
		base := strings.Split(arg, "/")
		basename = base[len(base)-1]

		if suffix != "" {
			basename = strings.TrimSuffix(basename, suffix)
		}

		switch zero {
		case true:
			newline = "\x00"
		default:
			newline = "\n"
		}

		fmt.Printf("%s%s", basename, newline)
		if flag.NArg() > 1 && !multiple {
			break
		}
	}
}
