package tree

type Node struct {
	tag        string
	attributes map[string]string
	text       string
	children   []*Node
}

func (parent *Node) AddChild(child *Node) {
	parent.children = append(parent.children, child)
}

func (parent *Node) Search(equalFn func(*Node) bool) *Node {
	return nil
}
