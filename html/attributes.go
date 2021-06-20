package html

import (
	"fmt"
)

func makeAttribute(name string) func(string) Attribute {
	return func(value string) Attribute {
		return Attribute{Name: name, Value: value}
	}
}

var Src = makeAttribute("src")
var Id = makeAttribute("id")
var Class = makeAttribute("class")
var Charset = makeAttribute("charset")
var Href = makeAttribute("href")
var Value = makeAttribute("value")
var Name = makeAttribute("name")
var Type = makeAttribute("type")
var Action = makeAttribute("action")
var Method = makeAttribute("method")

func (node Node) Media(value string) Node {
	return node.Attribute("media", value)
}

func (node Node) Rel(value string) Node {
	return node.Attribute("rel", value)
}

func (node Node) Content(value string) Node {
	return node.Attribute("content", value)
}

func (node Node) For(value string) Node {
	return node.Attribute("for", value)
}

func (node Node) Method(value string) Node {
	return node.Attribute("method", value)
}

func (node Node) Selected(value string) Node {
	return node.Attribute("selected", value)
}

func (node Node) Align(value string) Node {
	return node.Attribute("align", value)
}

func (node Node) Placeholder(value string) Node {
	return node.Attribute("placeholder", value)
}

func (node Node) Width(value uint) Node {
	return node.Attribute("width", fmt.Sprintf("%v", value))
}

func (node Node) Height(value uint) Node {
	return node.Attribute("height", fmt.Sprintf("%v", value))
}

func (node Node) WidthFloat(value float64) Node {
	return node.Attribute("width", fmt.Sprintf("%v", value))
}

func (node Node) HeightFloat(value float64) Node {
	return node.Attribute("height", fmt.Sprintf("%v", value))
}

func (node Node) Title(value string) Node {
	return node.Attribute("title", value)
}

func (node Node) Controls() Node {
	return node.Attribute("controls", "")
}

func (node Node) Autoplay() Node {
	return node.Attribute("autoplay", "")
}

func (node Node) Label(value string) Node {
	return node.Attribute("label", value)
}

func (node Node) Kind(value string) Node {
	return node.Attribute("kind", value)
}

func (node Node) Srclang(value string) Node {
	return node.Attribute("srclang", value)
}

func (node Node) Default() Node {
	return node.Attribute("default", "")
}

func (node Node) Multiple() Node {
	return node.Attribute("multiple", "")
}

func (node Node) Accept(value string) Node {
	return node.Attribute("accept", value)
}

func (node Node) D(value string) Node {
	return node.Attribute("d", value)
}

func (node Node) Fill(value string) Node {
	return node.Attribute("fill", value)
}

func (node Node) Stroke(value string) Node {
	return node.Attribute("stroke", value)
}
