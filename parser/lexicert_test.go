package parser_test

import (
	"fmt"
	"reflect"
	"testing"
	"xml-parser/parser"
)

func TestTypical(t *testing.T) {
	xml := "<html>konichiha</html>"
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
	tokens := parser.Analicys(xml)
	t.Run("typical", func(t *testing.T) {
		if !reflect.DeepEqual(tokens, target) {
			t.Errorf("%+v : %+v", xml, tokensToString(tokens))
		}
	})
}

func TestAttribute(t *testing.T) {
	xml := "<html key=\"value\" key2=\"va\\\"lue2\">konichiha</html>"
	target := []*parser.Token{
		{Kind: parser.LAB},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.TEXT, Value: "key"},
		{Kind: parser.EQ},
		{Kind: parser.TEXT, Value: "\"value\""},
		{Kind: parser.TEXT, Value: "key2"},
		{Kind: parser.EQ},
		{Kind: parser.TEXT, Value: "\"va\\\"lue2\""},
		{Kind: parser.RAB},
		{Kind: parser.TEXT, Value: "konichiha"},
		{Kind: parser.LAB},
		{Kind: parser.SLASH},
		{Kind: parser.TEXT, Value: "html"},
		{Kind: parser.RAB},
	}
	tokens := parser.Analicys(xml)
	t.Run("typical", func(t *testing.T) {
		if !reflect.DeepEqual(tokens, target) {
			t.Errorf("%+v : %+v", xml, tokensToString(tokens))
		}
	})
}

func TestNewLine(t *testing.T) {
	xml := `
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
	tokens := parser.Analicys(xml)
	t.Run("typical", func(t *testing.T) {
		if !reflect.DeepEqual(tokens, target) {
			t.Errorf("%+v : %+v", xml, tokensToString(tokens))
		}
	})
}

func tokensToString(tokens []*parser.Token) string {
	result := ""
    for _,v :=range tokens {
        result += fmt.Sprintf("%+v", *v)
    }
    return result
}
