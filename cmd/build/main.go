package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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
	source, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	converter := goldmark.New(goldmark.WithExtensions(extension.GFM))
	err = converter.Convert(source, &buf)
	if err != nil {
		return err
	}

	typst := exec.Command("typst", "compile", "-", outputFile)
	typst.Stdin = &buf
	typst.Stdout = os.Stdout
	typst.Stderr = os.Stderr
	return typst.Run()
}
