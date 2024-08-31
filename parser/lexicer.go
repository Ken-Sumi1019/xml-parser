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

func Lexical(html string) []*Token {
	lex := lexicer{text: html, index: 0}
	tokens := []*Token{}

	for lex.index < len(html) {
		switch lex.nextChar() {
		case '<':
			if lex.lookaheadChar(1) == '!' && lex.lookaheadChar(2) == '-' && lex.lookaheadChar(3) == '-' {
				lex.skipComment()
			} else {
				tokens = append(tokens, &Token{Kind: LAB})
				lex.next()
			}
		case '>':
			tokens = append(tokens, &Token{Kind: RAB})
			lex.next()
		case '=':
			tokens = append(tokens, &Token{Kind: EQ})
			lex.next()
		case '/':
			tokens = append(tokens, &Token{Kind: SLASH})
			lex.next()
		// https://triple-underscore.github.io/infra-ja.html#ascii-whitespace
		case ' ', '\n', '\r', '\t', '\f':
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

func (lex *lexicer) lookaheadChar(n int) uint8 {
	return lex.text[lex.index+n]
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
	return strings.Trim(lex.text[firstIdx:lex.index], "\n \r\t\f")
}

func (lex *lexicer) takeOutTextInQuotes() string {
	quote := lex.nextChar()
	resultSlice := make([]byte, 0, 128)
	lex.next()
	for lex.nextChar() != quote {
		if lex.nextChar() == '\\' {
			lex.next()
		}
		resultSlice = append(resultSlice, lex.nextChar())
		lex.next()
	}
	lex.next()
	return string(resultSlice)
}

func (lex *lexicer) skipComment() {
	for {
		if lex.nextChar() == '-' {
			if lex.lookaheadChar(1) == '-' && lex.lookaheadChar(2) == '>' {
				lex.next()
				lex.next()
				lex.next()
				break
			}
		}
		lex.next()
	}
}
