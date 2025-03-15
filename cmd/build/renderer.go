package main

import (
	"bytes"
	"log/slog"
	"strings"

	"github.com/yuin/goldmark/ast"
	extensionast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Renderer is based on [github.com/yuin/goldmark/renderer/html.Renderer].
type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindDocument, r.renderDocument)
	reg.Register(ast.KindHeading, r.renderHeading)
	reg.Register(ast.KindBlockquote, r.renderBlockquote)
	reg.Register(ast.KindCodeBlock, r.renderCodeBlock)
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	reg.Register(ast.KindHTMLBlock, r.renderHTMLBlock)
	reg.Register(ast.KindList, r.renderList)
	reg.Register(ast.KindListItem, r.renderListItem)
	reg.Register(ast.KindParagraph, r.renderParagraph)
	reg.Register(ast.KindTextBlock, r.renderTextBlock)
	reg.Register(ast.KindThematicBreak, r.renderThematicBreak)

	reg.Register(ast.KindAutoLink, r.renderAutoLink)
	reg.Register(ast.KindCodeSpan, r.renderCodeSpan)
	reg.Register(ast.KindEmphasis, r.renderEmphasis)
	reg.Register(ast.KindImage, r.renderImage)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindRawHTML, r.renderRawHTML)
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindString, r.renderString)
}

func (r *Renderer) renderDocument(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Heading)
		_, _ = w.WriteString(strings.Repeat("=", n.Level))
		_ = w.WriteByte(' ')
	} else {
		if node.NextSibling() != nil {
			_ = w.WriteByte('\n')
		}
		_ = w.WriteByte('\n')
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderBlockquote(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderBlockquote")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderCodeBlock")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderFencedCodeBlock")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderHTMLBlock")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderList")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderListItem")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if node.NextSibling() != nil {
			_ = w.WriteByte('\n')
		}
		_ = w.WriteByte('\n')
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTextBlock")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderThematicBreak")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderAutoLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderAutoLink")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("#raw")
		_ = w.WriteByte('(')
		_ = w.WriteByte('"')
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			v := c.(*ast.Text).Value(source)
			if bytes.HasSuffix(v, []byte("\n")) {
				strWrite(w, v[:len(v)-1])
				strWrite(w, []byte(" "))
			} else {
				strWrite(w, v)
			}
		}
		_ = w.WriteByte('"')
		return ast.WalkSkipChildren, nil
	} else {
		_ = w.WriteByte(')')
		return ast.WalkContinue, nil
	}
}

func (r *Renderer) renderEmphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Emphasis)
		fn := "#emph"
		if n.Level == 2 {
			fn = "#strong"
		}
		_, _ = w.WriteString(fn)
		_ = w.WriteByte('[')
	} else {
		_ = w.WriteByte(']')
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderImage")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderLink")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderRawHTML")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Text)
		if n.IsRaw() {
			unsafeWrite(w, n.Value(source))
		} else {
			contentWrite(w, n.Value(source))
			if n.HardLineBreak() {
				_, _ = w.WriteString(" \\\n")
			} else if n.SoftLineBreak() {
				_ = w.WriteByte('\n')
			}
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderString(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderString")
	return ast.WalkContinue, nil
}

// TableRenderer is based on [github.com/yuin/goldmark/extension.TableHTMLRenderer].
type TableRenderer struct{}

func NewTableRenderer() *TableRenderer {
	return &TableRenderer{}
}

func (r *TableRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(extensionast.KindTable, r.renderTable)
	reg.Register(extensionast.KindTableHeader, r.renderTableHeader)
	reg.Register(extensionast.KindTableRow, r.renderTableRow)
	reg.Register(extensionast.KindTableCell, r.renderTableCell)
}

func (r *TableRenderer) renderTable(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTable")
	return ast.WalkContinue, nil
}

func (r *TableRenderer) renderTableHeader(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTableHeader")
	return ast.WalkContinue, nil
}

func (r *TableRenderer) renderTableRow(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTableRow")
	return ast.WalkContinue, nil
}

func (r *TableRenderer) renderTableCell(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTableCell")
	return ast.WalkContinue, nil
}

// StrikethroughRenderer is based on [github.com/yuin/goldmark/extension.StrikethroughHTMLRenderer].
type StrikethroughRenderer struct{}

func NewStrikethroughRenderer() *StrikethroughRenderer {
	return &StrikethroughRenderer{}
}

func (r *StrikethroughRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(extensionast.KindStrikethrough, r.renderStrikethrough)
}

func (r *StrikethroughRenderer) renderStrikethrough(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderStrikethrough")
	return ast.WalkContinue, nil
}

// TaskCheckBoxRenderer is based on [github.com/yuin/goldmark/extension.TaskCheckBoxHTMLRenderer].
type TaskCheckBoxRenderer struct{}

func NewTaskCheckBoxRenderer() renderer.NodeRenderer {
	return &TaskCheckBoxRenderer{}
}

func (r *TaskCheckBoxRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(extensionast.KindTaskCheckBox, r.renderTaskCheckBox)
}

func (r *TaskCheckBoxRenderer) renderTaskCheckBox(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTaskCheckBox")
	return ast.WalkContinue, nil
}

func unsafeWrite(w util.BufWriter, p []byte) {
	_, _ = w.Write(p)
}

func contentWrite(w util.BufWriter, p []byte) {
	l := 0
	r := len(p)
	for i := 0; i < r; i++ {
		switch p[i] {
		case
			'*',  // at word boundaries (strong) and inside "/*" and "*/" (comment)
			'_',  // at word boundaries (emph)
			'`',  // likely everywhere (raw)
			':',  // inside "http://" and "https://" (link) and between the term opening slash and term closing colon (term)
			'<',  // when adjacent to text (label)
			'>',  // when adjacent to text (label)
			'@',  // almost everywhere (ref)
			'=',  // at line start (heading)
			'-',  // at line start (list) and inside "--" and "---" (symbols)
			'+',  // at line start (enum)
			'.',  // when following line start and unescaped digits (enum)
			'/',  // at line start (terms) and inside "//" (comment)
			'$',  // likely everywhere (math)
			'\\', // likely everywhere (linebreak, escape)
			'\'', // likely everywhere (smartquote)
			'"',  // likely everywhere (smartquote)
			'~',  // likely everywhere (symbols)
			'#',  // likely everywhere (scripting)
			'[',  // in square bracket markup (markup)
			']':  // in square bracket markup (markup)
			_, _ = w.Write(p[l:i])
			_ = w.WriteByte('\\')
			l = i
		}
	}
	_, _ = w.Write(p[l:r])
}

func strWrite(w util.BufWriter, p []byte) {
	l := 0
	r := len(p)
	for i := 0; i < r; i++ {
		switch p[i] {
		case '\\', '"':
			_, _ = w.Write(p[l:i])
			_ = w.WriteByte('\\')
			l = i
		case '\n':
			_, _ = w.Write(p[l:i])
			_, _ = w.WriteString("\\n")
			l = i + 1
		case '\r':
			_, _ = w.Write(p[l:i])
			_, _ = w.WriteString("\\r")
			l = i + 1
		case '\t':
			_, _ = w.Write(p[l:i])
			_, _ = w.WriteString("\\t")
			l = i + 1
		}
	}
	_, _ = w.Write(p[l:r])
}
