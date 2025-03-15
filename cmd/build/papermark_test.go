package main

import (
	"strings"
	"testing"
)

func TestPapermark(t *testing.T) {
	md := NewPapermark()

	tests := []struct {
		Markdown  string
		WantTypst string
	}{
		// Hard line breaks
		{Markdown: "foo  \nbaz\n", WantTypst: "foo \\\nbaz\n"},
		{Markdown: "*foo  \nbar*\n", WantTypst: "#emph[foo \\\nbar]\n"},
		{Markdown: "`code  \nspan`\n", WantTypst: "#raw(\"code   span\")\n"},
		{Markdown: "<a href=\"foo  \nbar\">\n", WantTypst: "\n"},
		{Markdown: "foo  \n", WantTypst: "foo\n"},
		{Markdown: "### foo  \n", WantTypst: "=== foo\n"},

		// Soft line breaks
		{Markdown: "foo\nbaz\n", WantTypst: "foo\nbaz\n"},
		{Markdown: "foo \n baz\n", WantTypst: "foo\nbaz\n"},

		// If expression wasn't terminated with ';' and '.' wasn't escaped,
		// ".body" would be interpreted as part of the expression.
		{Markdown: "*foo*.body\n", WantTypst: "#emph[foo];\\.body\n"},
	}

	for _, tt := range tests {
		t.Run(tt.Markdown, func(t *testing.T) {
			b := new(strings.Builder)
			err := md.Convert([]byte(tt.Markdown), b)
			if err != nil {
				t.Fatalf("got %v err", err)
			}
			if got, want := b.String(), tt.WantTypst; got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}
