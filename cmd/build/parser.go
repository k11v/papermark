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

type ImageBlockASTTransformer struct{}

func NewImageBlockASTTransformer() *ImageBlockASTTransformer {
	return &ImageBlockASTTransformer{}
}

func (b *ImageBlockASTTransformer) Transform(doc *ast.Document, reader text.Reader, pc parser.Context) {
	imageBlockParagraphs := make([]*ast.Paragraph, 0)

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if n.Kind() == ast.KindParagraph {
				if n.ChildCount() == 1 {
					if n.FirstChild().Kind() == ast.KindImage {
						imageBlockParagraphs = append(imageBlockParagraphs, n.(*ast.Paragraph))
					}
				}
			}
		}
		return ast.WalkContinue, nil
	})

	for _, n := range imageBlockParagraphs {
		imageBlock := NewImageBlock()
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			imageBlock.AppendChild(imageBlock, c)
		}
		n.Parent().ReplaceChild(n.Parent(), n, imageBlock)
	}
}
