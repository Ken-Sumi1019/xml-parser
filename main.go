package main

import (
	"xml-parser/tree"
)

type mountain struct {
	root *tree.Node
}

func Parse() *mountain {
	return &mountain{}
}
