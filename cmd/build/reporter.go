package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/yuin/goldmark/text"
)

type Position struct {
	File   string
	Line   int // starting at 1
	Column int // byte count, starting at 1
}

// String returns a string in one of several forms:
//
//	file:line:column
//	file:line
//	line:column
//	line
//	file
//	-
func (pos Position) String() string {
	s := pos.File
	if pos.Line != 0 {
		if s != "" {
			s += ":"
		}
		s += strconv.Itoa(pos.Line)
		if pos.Column != 0 {
			s += fmt.Sprintf(":%d", pos.Column)
		}
	}
	if s == "" {
		s = "-"
	}
	return s
}

type Diagnostic struct {
	Seg      *text.Segment
	Category string // optional
	Message  string
}

type Reporter struct {
	file       string
	data       []byte
	context    int
	lineStarts []int
}

func NewReporter(file string, data []byte, context int) *Reporter {
	lineStarts := []int{0}
	for {
		i := bytes.IndexByte(data[lineStarts[len(lineStarts)-1]:], '\n')
		if i == -1 {
			break
		}
		lineStarts = append(lineStarts, i+1)
	}

	return &Reporter{
		file:       file,
		data:       data,
		context:    context,
		lineStarts: lineStarts,
	}
}

func (r *Reporter) Report(diag Diagnostic) {
	startLine, found := slices.BinarySearch(r.lineStarts, diag.Seg.Start)
	if !found {
		startLine--
	}
	startColumn := diag.Seg.Start - r.lineStarts[startLine]

	pos := Position{
		File:   r.file,
		Line:   startLine + 1,
		Column: startColumn + 1,
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", pos, diag.Message)

	if r.context > 0 {
		stopLine, found := slices.BinarySearch(r.lineStarts, diag.Seg.Stop)
		if !found {
			stopLine--
		}
		for i := startLine - r.context; i <= stopLine+r.context; i++ {
			if 0 <= i && i < len(r.lineStarts) {
				start := r.lineStarts[i]
				stop := len(r.data) - 1
				if 0 <= i+1 && i+1 < len(r.lineStarts) {
					stop = r.lineStarts[i+1] - 1
				}
				fmt.Fprintf(os.Stderr, "%d\t%s\n", i, r.data[start:stop])
			}
		}
	}
}

func (r *Reporter) Reportf(seg *text.Segment, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	r.Report(Diagnostic{Seg: seg, Message: msg})
}
