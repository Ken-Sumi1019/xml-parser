package parser_test

import (
	"fmt"
	"testing"
	"xml-parser/parser"
)

func TestParserTypical(t *testing.T) {
	xml := "<html>konichiha</html>"
	node := parser.Parse(xml)

	target := &parser.Node{
		Tag: "html",
        Attributes: map[string]string{},
        Children: []*parser.Node{
            {Text: "konichiha"},
        },
	}
	t.Run("parser typical", func(t *testing.T) {
		if treeToString(node.Children[0],"") != treeToString(target,"") {
			t.Errorf("%+v : %+v", xml, treeToString(target, ""))
		}
	})
}

func TestParserAttribute(t *testing.T) {
	xml := "<html key=\"value\" key2=\"va\\\"lue2\">konichiha</html>"
	node := parser.Parse(xml)

	target := &parser.Node{
		Tag: "html",
        Attributes: map[string]string{"key": "\"value\"", "key2": "\"va\\\"lue2\""},
        Children: []*parser.Node{
            {Text: "konichiha"},
        },
	}
	t.Run("parser attribute", func(t *testing.T) {
		if treeToString(node.Children[0],"") != treeToString(target,"") {
			t.Errorf("%+v : %+v", xml, treeToString(node.Children[0], ""))
		}
	})
}
func treeToString(node *parser.Node, result string) string {
	result += fmt.Sprintf("<Tag: %+v, Attributes: %+v, Text: %+v, Children: [", node.Tag, node.Attributes, node.Text)
	for _, v := range node.Children {
		result = treeToString(v,result)
	}
	result += "]>,\n"
	return result
}
