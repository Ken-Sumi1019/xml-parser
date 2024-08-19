package parser_test

import (
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
			t.Errorf("%+v : %+v", xml, tokens)
		}
	})
}
