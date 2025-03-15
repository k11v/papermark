package main

import (
	"github.com/yuin/goldmark/ast"
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
