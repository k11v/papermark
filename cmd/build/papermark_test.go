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
		{Markdown: "foo  \nbaz\n", WantTypst: "foo \\\nbaz\n\n"},
		{Markdown: "*foo  \nbar*\n", WantTypst: "#emph[foo \\\nbar]\n\n"},
		{Markdown: "`code  \nspan`\n", WantTypst: "#raw(\"code   span\")\n\n"},
		{Markdown: "<a href=\"foo  \nbar\">\n", WantTypst: "\n\n"},
		{Markdown: "foo  \n", WantTypst: "foo\n\n"},
		{Markdown: "### foo  \n", WantTypst: "=== foo\n\n"},

		// Soft line breaks
		{Markdown: "foo\nbaz\n", WantTypst: "foo\nbaz\n\n"},
		{Markdown: "foo \n baz\n", WantTypst: "foo\nbaz\n\n"},
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
