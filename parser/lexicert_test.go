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
		&parser.Token{Kind: parser.LAB},
		&parser.Token{Kind: parser.TEXT, Value: "html"},
		&parser.Token{Kind: parser.RAB},
		&parser.Token{Kind: parser.TEXT, Value: "konichiha"},
		&parser.Token{Kind: parser.LAB},
		&parser.Token{Kind: parser.SLASH},
		&parser.Token{Kind: parser.TEXT, Value: "html"},
		&parser.Token{Kind: parser.RAB},
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
		&parser.Token{Kind: parser.LAB},
		&parser.Token{Kind: parser.TEXT, Value: "html"},
		&parser.Token{Kind: parser.TEXT, Value: "key"},
		&parser.Token{Kind: parser.EQ},
		&parser.Token{Kind: parser.TEXT, Value: "\"value\""},
		&parser.Token{Kind: parser.TEXT, Value: "key2"},
		&parser.Token{Kind: parser.EQ},
		&parser.Token{Kind: parser.TEXT, Value: "\"va\\\"lue2\""},
		&parser.Token{Kind: parser.RAB},
		&parser.Token{Kind: parser.TEXT, Value: "konichiha"},
		&parser.Token{Kind: parser.LAB},
		&parser.Token{Kind: parser.SLASH},
		&parser.Token{Kind: parser.TEXT, Value: "html"},
		&parser.Token{Kind: parser.RAB},
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
