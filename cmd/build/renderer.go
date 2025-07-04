package main

import (
	"bytes"
	_ "embed"
	"log/slog"
	"strconv"
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

//go:embed template.typ
var templateBytes []byte

func (r *Renderer) renderDocument(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		unsafeWrite(w, templateBytes)
		_, _ = w.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Heading)
		_, _ = w.WriteString(strings.Repeat("=", n.Level))
		_, _ = w.WriteRune(' ')
	} else {
		_, _ = w.WriteRune('\n')
		if node.NextSibling() != nil {
			_, _ = w.WriteRune('\n')
		}
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
	if entering {
		n := node.(*ast.FencedCodeBlock)
		_, _ = w.WriteString("#")
		_, _ = w.WriteString("figure")
		_, _ = w.WriteString("(\n")

		// TODO: Get caption from attributes.
		_, _ = w.WriteString("caption: ")
		_, _ = w.WriteString(`"`)
		strWrite(w, []byte("Lorem ipsum."))
		_, _ = w.WriteString(`"`)
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString("raw")
		_, _ = w.WriteString("(")

		_, _ = w.WriteString("block: ")
		_, _ = w.WriteString("true")
		_, _ = w.WriteString(", ")

		if lang := n.Language(source); lang != nil {
			_, _ = w.WriteString("lang: ")
			_, _ = w.WriteString(`"`)
			strWrite(w, lang)
			_, _ = w.WriteString(`"`)
			_, _ = w.WriteString(", ")
		}

		_, _ = w.WriteString(`"`)
		for i := 0; i < n.Lines().Len(); i++ {
			l := n.Lines().At(i)
			strWrite(w, l.Value(source))
		}
		_, _ = w.WriteString(`"`)

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";\n")

		_, _ = w.WriteString("#")
		_, _ = w.WriteString("label")
		_, _ = w.WriteString("(")

		// TODO: Get label from attributes.
		_, _ = w.WriteString(`"`)
		strWrite(w, []byte("lorem"))
		_, _ = w.WriteString(`"`)

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";\n")
		if node.NextSibling() != nil {
			_, _ = w.WriteString("\n")
		}
		return ast.WalkSkipChildren, nil
	} else {
		return ast.WalkContinue, nil
	}
}

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderHTMLBlock")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.List)

		if n.IsOrdered() {
			_, _ = w.WriteString("#enum")
		} else {
			_, _ = w.WriteString("#list")
		}
		_, _ = w.WriteString("(\n")

		_, _ = w.WriteString("tight: ")
		_, _ = w.WriteString(strconv.FormatBool(n.IsTight))
		_, _ = w.WriteString(",\n")

		if n.IsOrdered() && n.Start != 1 {
			_, _ = w.WriteString("start: ")
			_, _ = w.WriteString(strconv.Itoa(n.Start))
			_, _ = w.WriteString(",\n")
		}
	} else {
		_, _ = w.WriteRune(')')
		_, _ = w.WriteString(";\n")
		if node.NextSibling() != nil {
			_, _ = w.WriteRune('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteRune('[')
		fc := n.FirstChild()
		if fc != nil {
			if _, ok := fc.(*ast.TextBlock); !ok {
				_, _ = w.WriteRune('\n')
			}
		}
	} else {
		_, _ = w.WriteRune(']')
		_, _ = w.WriteString(",\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteRune('\n')
		if node.NextSibling() != nil {
			_, _ = w.WriteRune('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if n.NextSibling() != nil && n.FirstChild() != nil {
			_, _ = w.WriteRune('\n')
		}
	}
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
		_, _ = w.WriteRune('(')
		_, _ = w.WriteString("block: false, ")
		_, _ = w.WriteRune('"')
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			v := c.(*ast.Text).Value(source)
			if bytes.HasSuffix(v, []byte("\n")) {
				strWrite(w, v[:len(v)-1])
				strWrite(w, []byte(" "))
			} else {
				strWrite(w, v)
			}
		}
		_, _ = w.WriteRune('"')
		_, _ = w.WriteRune(')')
		return ast.WalkSkipChildren, nil
	} else {
		_, _ = w.WriteRune(';')
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
		_, _ = w.WriteRune('[')
	} else {
		_, _ = w.WriteRune(']')
		_, _ = w.WriteRune(';')
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Image)
		_, _ = w.WriteString("#")
		_, _ = w.WriteString("image")
		_, _ = w.WriteString("(")

		_, _ = w.WriteString(`"`)
		strWrite(w, n.Destination)
		_, _ = w.WriteString(`"`)

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";")
		return ast.WalkSkipChildren, nil
	} else {
		return ast.WalkContinue, nil
	}
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
				_, _ = w.WriteRune('\n')
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
	if entering {
		n := node.(*extensionast.Table)
		_, _ = w.WriteString("#")
		_, _ = w.WriteString("figure")
		_, _ = w.WriteString("(\n")

		// TODO: Get caption from attributes.
		_, _ = w.WriteString("caption: ")
		_, _ = w.WriteString(`"`)
		strWrite(w, []byte("Lorem ipsum."))
		_, _ = w.WriteString(`"`)
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString("table")
		_, _ = w.WriteString("(\n")

		_, _ = w.WriteString("columns: ")
		_, _ = w.WriteString("(")
		for i := 0; i < len(n.Alignments); i++ {
			if i != 0 {
				_, _ = w.WriteString(", ")
			}
			_, _ = w.WriteString("auto")
		}
		_, _ = w.WriteString(")")
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString("align: ")
		_, _ = w.WriteString("(")
		for i, a := range n.Alignments {
			if i != 0 {
				_, _ = w.WriteString(", ")
			}
			switch a {
			case extensionast.AlignLeft:
				_, _ = w.WriteString("left")
			case extensionast.AlignRight:
				_, _ = w.WriteString("right")
			case extensionast.AlignCenter:
				_, _ = w.WriteString("center")
			default:
				_, _ = w.WriteString("auto")
			}
		}
		_, _ = w.WriteString(")")
		_, _ = w.WriteString(",\n")
	} else {
		_, _ = w.WriteString(")")
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";\n")

		_, _ = w.WriteString("#")
		_, _ = w.WriteString("label")
		_, _ = w.WriteString("(")

		// TODO: Get label from attributes.
		_, _ = w.WriteString(`"`)
		strWrite(w, []byte("lorem"))
		_, _ = w.WriteString(`"`)

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";\n")
		if node.NextSibling() != nil {
			_, _ = w.WriteRune('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *TableRenderer) renderTableHeader(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("table.header")
		_, _ = w.WriteString("(")
	} else {
		_, _ = w.WriteString(")")
		_, _ = w.WriteString(",\n")
	}
	return ast.WalkContinue, nil
}

func (r *TableRenderer) renderTableRow(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString(",\n")
	}
	return ast.WalkContinue, nil
}

func (r *TableRenderer) renderTableCell(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if node.PreviousSibling() != nil {
			_, _ = w.WriteString(", ")
		}
		_, _ = w.WriteString("[")
	} else {
		_, _ = w.WriteString("]")
	}
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

func NewTaskCheckBoxRenderer() *TaskCheckBoxRenderer {
	return &TaskCheckBoxRenderer{}
}

func (r *TaskCheckBoxRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(extensionast.KindTaskCheckBox, r.renderTaskCheckBox)
}

func (r *TaskCheckBoxRenderer) renderTaskCheckBox(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	slog.Error("unimplemented renderTaskCheckBox")
	return ast.WalkContinue, nil
}

type ImageBlockRenderer struct{}

func NewImageBlockRenderer() *ImageBlockRenderer {
	return &ImageBlockRenderer{}
}

func (r *ImageBlockRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindImageBlock, r.renderImageBlock)
}

func (r *ImageBlockRenderer) renderImageBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("#")
		_, _ = w.WriteString("figure")
		_, _ = w.WriteString("(\n")

		// TODO: Get caption from attributes.
		_, _ = w.WriteString("caption: ")
		_, _ = w.WriteString(`"`)
		strWrite(w, []byte("Lorem ipsum."))
		_, _ = w.WriteString(`"`)
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString("[")
	} else {
		_, _ = w.WriteString("]")
		_, _ = w.WriteString(",\n")

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";\n")

		_, _ = w.WriteString("#")
		_, _ = w.WriteString("label")
		_, _ = w.WriteString("(")

		// TODO: Get label from attributes.
		_, _ = w.WriteString(`"`)
		strWrite(w, []byte("lorem"))
		_, _ = w.WriteString(`"`)

		_, _ = w.WriteString(")")
		_, _ = w.WriteString(";\n")
		if n.NextSibling() != nil {
			_, _ = w.WriteString("\n")
		}
	}
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
			_, _ = w.WriteRune('\\')
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
			_, _ = w.WriteRune('\\')
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
