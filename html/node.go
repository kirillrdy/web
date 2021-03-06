package html

import (
	golang_html "html"
	"io"
	"strings"
)

type Node struct {
	nodeType   string
	Attributes []Attribute
	children   []Node
	text       string
}

//TODO pass buffer for even faster writing
func (node Node) attributesAsString(result *strings.Builder) {
	//TODO test this theory
	for i := range node.Attributes {
		//Note this is done for performance
		result.WriteString(" ")
		result.WriteString(node.Attributes[i].Name)
		result.WriteString("=\"")
		result.WriteString(node.Attributes[i].Value)
		result.WriteString("\"")
	}
}

func (node Node) WriteTo(writer io.Writer) (int64, error) {
	written, err := io.WriteString(writer, node.String())
	return int64(written), err
}

func (node Node) String() string {
	var builder strings.Builder
	node.writeToBuffer(&builder)
	return builder.String()
}

func (node Node) writeToBuffer(buffer *strings.Builder) {

	//TODO peformance bench this 'if' vs 'if len' vs no if
	if node.nodeType == "html" {
		buffer.WriteString("<!DOCTYPE html>")
	}

	if node.nodeType == "" {
		buffer.WriteString(node.text)
		return
	}

	buffer.WriteString("<")
	buffer.WriteString(node.nodeType)
	node.attributesAsString(buffer)
	buffer.WriteString(">")

	if len(node.children) > 0 {
		for i := range node.children {
			node.children[i].writeToBuffer(buffer)
		}
	} else {
		buffer.WriteString(node.text)
	}

	buffer.WriteString("</")
	buffer.WriteString(node.nodeType)
	buffer.WriteString(">")

}

func Text(text string) Node {
	return Node{text: golang_html.EscapeString(text)}
}

func TextUnsafe(text string) Node {
	return Node{text: text}
}

// Text will excape any input as html
func (node Node) Text(text string) Node {
	node.text = golang_html.EscapeString(text)
	return node
}

// Same as Text() but doesn't escape html
func (node Node) TextUnsafe(text string) Node {
	node.text = text
	return node
}

func (node *Node) Append(children ...Node) {
	node.children = append(node.children, children...)
}

func (node Node) Children(children ...Node) Node {
	node.children = append(node.children, children...)
	return node
}

func (node Node) Attribute(attributeName, value string) Node {
	node.Attributes = append(node.Attributes, Attribute{Name: attributeName, Value: value})
	return node
}
