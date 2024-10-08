package parser_test

import (
	"fmt"
	"html-parser/parser"
	"reflect"
	"testing"
)

func TestTypical(t *testing.T) {
	html := `<html>konichiha</html>`
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func TestAttribute(t *testing.T) {
	html := `<html key="value" key2="va\"lue2">konichiha</html>`
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.TEXT, Value: "key"},
		{Kind: parser.EQ},
		{Kind: parser.TEXT, Value: "value"},
		{Kind: parser.TEXT, Value: "key2"},
		{Kind: parser.EQ},
		{Kind: parser.TEXT, Value: "va\"lue2"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func TestNewLine(t *testing.T) {
	html := `
    <html>
        konichiha
    </html>
    `
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func TestAsciiWhitespace(t *testing.T) {
	html := "<html>      \n\t\t\fkonichiha\n\t\t\f \n </html>"
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func TestNest(t *testing.T) {
	html := `
    <html>
        <body>
            konichiha
        </body>
    </html>
    `
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "body"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "body"},
		{Kind: parser.RAB},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func TestNestWithAttribute(t *testing.T) {
	html := `
    <html key="value" key2="va\"lue2">
        <body>
            konichiha
        </body>
    </html>
    `
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.TEXT, Value: "key"},
		{Kind: parser.EQ},
		{Kind: parser.TEXT, Value: "value"},
		{Kind: parser.TEXT, Value: "key2"},
		{Kind: parser.EQ},
		{Kind: parser.TEXT, Value: "va\"lue2"},
		{Kind: parser.RAB},
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "body"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "body"},
		{Kind: parser.RAB},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func TestComment(t *testing.T) {
	html := `
    <html>
		<!--My favorite operators are
        > and <!-->konichiha
    </html>
    `
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Lexical(html)
	if !reflect.DeepEqual(tokens, target) {
		t.Errorf("%+v : %+v", html, tokensToString(tokens))
	}
}

func tokensToString(tokens []*parser.Token) string {
	result := ""
	for _, v := range tokens {
		result += fmt.Sprintf("%+v", *v)
	}
	return result
}
