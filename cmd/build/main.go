package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var outputFileFlag = flag.String("o", "", "output file")

func main() {
	flag.Parse()

	outputFile := *outputFileFlag
	if outputFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "error: empty output file flag\n")
		os.Exit(1)
	}

	inputFile := flag.Arg(0)
	if inputFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "error: empty input file arg\n")
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		_, _ = fmt.Fprint(os.Stderr, "error: extra args\n")
		os.Exit(1)
	}

	err := run(outputFile, inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(outputFile, inputFile string) error {
	typstCmd := exec.Command("typst", "compile", inputFile, outputFile)
	typstCmd.Stdout = os.Stdout
	typstCmd.Stderr = os.Stderr
	return typstCmd.Run()
}
