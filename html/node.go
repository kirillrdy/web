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

func (node Node) attributesAsString(result *strings.Builder) {
	for i := range node.Attributes {
		result.WriteString(" ")
		result.WriteString(node.Attributes[i].Name)
		result.WriteString("=\"")
		result.WriteString(node.Attributes[i].Value)
		result.WriteString("\"")
	}
}

func (node Node) attributesToWriter(writer io.Writer) (int, error) {
	for i := range node.Attributes {
		// TODO is single Write call faster with fmt.Sprintf ???
		//TODO return value of WriteString
		io.WriteString(writer, " ")
		io.WriteString(writer, node.Attributes[i].Name)
		io.WriteString(writer, "=\"")
		io.WriteString(writer, node.Attributes[i].Value)
		io.WriteString(writer, "\"")
	}
	return 0, nil
}

func (node Node) WriteTo(writer io.Writer) (int64, error) {
	n, err := node.writeToWriter(writer)
	return int64(n), err
}

func (node Node) String() string {
	var builder strings.Builder
	// TODO replace with using the writeToWritier
	node.writeToWriter(&builder)
	return builder.String()
}

// TODO return values of io.WriteString
func (node Node) writeToWriter(writer io.Writer) (int, error) {

	//TODO peformance bench this 'if' vs 'if len' vs no if
	if node.nodeType == "html" {
		io.WriteString(writer, "<!DOCTYPE html>")
	}

	if node.nodeType == "" {
		return io.WriteString(writer, node.text)
	}

	io.WriteString(writer, "<")
	io.WriteString(writer, node.nodeType)
	node.attributesToWriter(writer)
	io.WriteString(writer, ">")

	if len(node.children) > 0 {
		for i := range node.children {
			node.children[i].writeToWriter(writer)
		}
	} else {
		io.WriteString(writer, node.text)
	}

	io.WriteString(writer, "</")
	io.WriteString(writer, node.nodeType)
	io.WriteString(writer, ">")

	return 0, nil
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
