package parser

type Parser struct {
    tokens []*Token
	stack []*Node
}

type Node struct {
	tag        string
	attributes map[string]string
	text       string
	children   []*Node
}

func (parser *Parser)Parse(xml string) *Node {
    tokens := Lexical(xml)
    parser.tokens = tokens
}
