package parser

import "strings"

type Token struct {
	Kind  uint8
	Value string
}

var (
	TEXT  uint8 = 0
	EQ    uint8 = 1 // =
	LAB   uint8 = 2 // <
	RAB   uint8 = 3 // >
	SLASH uint8 = 4 // /
)

func Lexical(xml string) []*Token {
	lex := lexicer{text: xml, index: 0}
	tokens := []*Token{}

	for lex.index < len(xml) {
		switch lex.nextChar() {
		case '<':
			tokens = append(tokens, &Token{Kind: LAB})
			lex.next()
		case '>':
			tokens = append(tokens, &Token{Kind: RAB})
			lex.next()
		case '=':
			tokens = append(tokens, &Token{Kind: EQ})
			lex.next()
		case '/':
			tokens = append(tokens, &Token{Kind: SLASH})
			lex.next()
		case ' ', '\n', '\r':
			lex.next()
		case '\'', '"':
			text := lex.takeOutTextInQuotes()
			tokens = append(tokens, &Token{Kind: TEXT, Value: text})
		default:
			text := lex.takeOutText()
			tokens = append(tokens, &Token{Kind: TEXT, Value: text})
		}
	}

	return tokens
}

type lexicer struct {
	text  string
	index int
}

func (lex *lexicer) nextChar() uint8 {
	return lex.text[lex.index]
}

func (lex *lexicer) next() int {
	lex.index += 1
	return lex.index
}

func (lex *lexicer) takeOutText() string {
	firstIdx := lex.index
	for lex.nextChar() != '<' && lex.nextChar() != '>' && lex.nextChar() != ' ' && lex.nextChar() != '=' {
		lex.next()
	}
	return strings.Trim(lex.text[firstIdx:lex.index], "\n ")
}

func (lex *lexicer) takeOutTextInQuotes() string {
	firstIdx := lex.index
	quote := lex.nextChar()
	lex.next()
	for lex.nextChar() != quote {
		if lex.nextChar() == '\\' {
			lex.next()
		}
		lex.next()
	}
	lex.next()
	return lex.text[firstIdx:lex.index]
}
