package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var (
	outputFileFlag = flag.String("o", "", "output file")
	sourceFileFlag = flag.String("s", "", "source file")
)

func main() {
	flag.Parse()

	outputFile := *outputFileFlag
	if outputFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "error: empty output file flag\n")
		os.Exit(1)
	}

	sourceFile := *sourceFileFlag

	inputFile := flag.Arg(0)
	if inputFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "error: empty input file arg\n")
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		_, _ = fmt.Fprint(os.Stderr, "error: extra args\n")
		os.Exit(1)
	}

	err := run(outputFile, sourceFile, inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(outputFile, sourceFile, inputFile string) error {
	source, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	converter := NewPapermark()
	err = converter.Convert(source, &buf)
	if err != nil {
		return err
	}

	if sourceFile != "" {
		err = os.WriteFile(sourceFile, buf.Bytes(), 0o600)
		if err != nil {
			return err
		}
	}

	typst := exec.Command("typst", "compile", "-", outputFile)
	typst.Stdin = &buf
	typst.Stdout = os.Stdout
	typst.Stderr = os.Stderr
	return typst.Run()
}

func NewPapermark() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,         // https://github.github.com/gfm/#autolinks-extension-
			&TableExtension{},         // https://github.github.com/gfm/#tables-extension-
			&StrikethroughExtension{}, // https://github.github.com/gfm/#strikethrough-extension-
			&TaskCheckBoxExtension{},  // https://github.github.com/gfm/#task-list-items-extension-
			&ImageBlockExtension{},
			// TODO: Math.
			// TODO: Footnotes (https://github.blog/changelog/2021-09-30-footnotes-now-supported-in-markdown-fields/).
			// TODO: Wikilinks.
			// TODO: Attributes.
			// TODO: YAML metadata.
		),
		goldmark.WithRenderer(
			renderer.NewRenderer(renderer.WithNodeRenderers(
				util.Prioritized(NewRenderer(), 1000)),
			),
		),
	)
}

// TableExtension is based on [github.com/yuin/goldmark/extension.Table].
type TableExtension struct{}

func (e *TableExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithParagraphTransformers(
			util.Prioritized(extension.NewTableParagraphTransformer(), 200),
		),
		parser.WithASTTransformers(
			util.Prioritized(extension.NewTableASTTransformer(), 0),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTableRenderer(), 500),
	))
}

// StrikethroughExtension is based on [github.com/yuin/goldmark/extension.Strikethrough].
type StrikethroughExtension struct{}

func (e *StrikethroughExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(extension.NewStrikethroughParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewStrikethroughRenderer(), 500),
	))
}

// TaskCheckBoxExtension is based on [github.com/yuin/goldmark/extension.TaskList].
type TaskCheckBoxExtension struct{}

func (e *TaskCheckBoxExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(extension.NewTaskCheckBoxParser(), 0),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTaskCheckBoxRenderer(), 500),
	))
}

type ImageBlockExtension struct{}

func (e *ImageBlockExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithParagraphTransformers(
			util.Prioritized(NewImageBlockParagraphTransformer(), 200),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewImageBlockRenderer(), 500),
	))
}
