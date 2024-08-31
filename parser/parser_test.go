package parser_test

import (
	"fmt"
	"html-parser/parser"
	"testing"
)

func TestParserTypical(t *testing.T) {
	html := `
		<html>
			konichiha
		</html>`
	node := parser.Parse(html)

	target := &parser.Node{
		Tag: "html",
        Attributes: map[string]string{},
        Children: []*parser.Node{
            {Text: "konichiha"},
        },
	}
	t.Run("parser typical", func(t *testing.T) {
		if treeToString(node.Children[0],"") != treeToString(target,"") {
			t.Errorf("%+v : %+v", html, treeToString(target, ""))
		}
	})
}

func TestParserAttribute(t *testing.T) {
	html := `<html key="value" key2="va\"lue2">konichiha</html>`
	node := parser.Parse(html)

	target := &parser.Node{
		Tag: "html",
        Attributes: map[string]string{"key": "value", "key2": "va\"lue2"},
        Children: []*parser.Node{
            {Text: "konichiha"},
        },
	}
	t.Run("parser attribute", func(t *testing.T) {
		if treeToString(node.Children[0],"") != treeToString(target,"") {
			t.Errorf("%+v : %+v", html, treeToString(node.Children[0], ""))
		}
	})
}

func TestParserDistortedHtml(t *testing.T) {
	html := `
		<html key="value" key2="va\"lue2">
			<div>
				konichiha
		</html>`
	node := parser.Parse(html)

	target := &parser.Node{
		Tag: "html",
        Attributes: map[string]string{"key": "value", "key2": "va\"lue2"},
        Children: []*parser.Node{
			{
				Tag: "div",
				Children: []*parser.Node{
					{Text: "konichiha"},
				},
			},
        },
	}
	t.Run("parser attribute", func(t *testing.T) {
		if treeToString(node.Children[0],"") != treeToString(target,"") {
			t.Errorf("%+v : %+v", html, treeToString(node.Children[0], ""))
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
