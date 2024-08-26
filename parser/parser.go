package parser

import (
	"fmt"
)

type Parser struct {
    tokens []*Token
	stack []*Node
    root *Node
    index int
}

type Node struct {
	Tag        string
	Attributes map[string]string
	Text       string
	Children   []*Node
}

func Parse(xml string) *Node {
    tokens := Lexical(xml)
    parser := &Parser{tokens: tokens, index: 0, root: &Node{}}
    parser.tokens = tokens
    parser.stack = append(parser.stack, parser.root)
    parser.analyze()
    return parser.root
}

func (parser *Parser)analyze() {
	
    for parser.index < len(parser.tokens) && len(parser.stack) > 0 {
        fmt.Println(parser.nextToken().Kind)
        fmt.Printf("%+v \n", parser.stack[len(parser.stack)-1])
        fmt.Printf("%+v \n", parser.stack)
        switch parser.nextToken().Kind {
        case LAB:
            if parser.lookaheadToken(1).Kind == SLASH {
                child := parser.stack[len(parser.stack)-1]
                parser.stack = parser.stack[:len(parser.stack)-1]
                parser.stack[len(parser.stack)-1].Children = append(parser.stack[len(parser.stack)-1].Children, child)
                parser.next()
                parser.next()
                parser.next()
                parser.next()
            } else {
                parser.stackNode()
            }
        case TEXT:
            parser.stack[len(parser.stack)-1].Children = append(
                parser.stack[len(parser.stack)-1].Children,
                &Node{Text: parser.nextToken().Value},
            )
            parser.next()
        }
    }
}

func (parser *Parser)nextToken() *Token {
    return parser.tokens[parser.index]
}

func (parser *Parser)next() int {
    parser.index += 1
    return parser.index
}

func (parser *Parser)lookaheadToken(n int) *Token {
    return parser.tokens[parser.index + n]
}

func (parser *Parser)stackNode() {
    parser.next()
    newNode := &Node{Tag: parser.nextToken().Value, Attributes: map[string]string{}}
    parser.next()
    for parser.nextToken().Kind != RAB {
        if parser.nextToken().Kind == TEXT && parser.lookaheadToken(1).Kind == EQ && parser.lookaheadToken(2).Kind == TEXT {
            newNode.Attributes[parser.nextToken().Value] = parser.lookaheadToken(2).Value
            parser.next()
            parser.next()
            parser.next()
        } else {
            fmt.Println("エラー")
        }
    }
    parser.stack = append(parser.stack, newNode)
    parser.next()
}
