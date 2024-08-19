package main

type mountain struct {
	tag        string
	attributes map[string]string
	text       string
	children   []*mountain
}

func Parse() *mountain {
	return &mountain{}
}

func (parent *mountain) AddChild(child *mountain) {
	parent.children = append(parent.children, child)
}

func (parent *mountain) Search(equalFn func(*mountain) bool) *mountain {
	return nil
}

func main() {}
