package main

import (
	"log/slog"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var KindImageBlock = ast.NewNodeKind("ImageBlock")

type ImageBlock struct {
	ast.BaseBlock
}

func NewImageBlock() *ImageBlock {
	return &ImageBlock{}
}

func (n *ImageBlock) Kind() ast.NodeKind {
	return KindImageBlock
}

func (n *ImageBlock) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

type ImageBlockParagraphTransformer struct{}

func NewImageBlockParagraphTransformer() *ImageBlockParagraphTransformer {
	return &ImageBlockParagraphTransformer{}
}

func (b *ImageBlockParagraphTransformer) Transform(node *ast.Paragraph, reader text.Reader, pc parser.Context) {
	slog.Error("unimplemented ImageBlockParagraphTransformer")
}
