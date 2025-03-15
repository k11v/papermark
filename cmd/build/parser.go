package main

import (
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

func (b *ImageBlockParagraphTransformer) Transform(n *ast.Paragraph, reader text.Reader, pc parser.Context) {
	if n.ChildCount() == 1 {
		c := n.FirstChild()
		if c.Kind() == ast.KindImage {
			imageBlock := NewImageBlock()
			imageBlock.AppendChild(imageBlock, c)
			n.Parent().ReplaceChild(n.Parent(), n, imageBlock)
		}
	}
}
