package html

func makeNodeFunc(name string) func(...Attribute) func(...Node) Node {
	return func(attributes ...Attribute) func(...Node) Node {
		node := Node{nodeType: name}
		for _, attribute := range attributes {
			node.Attributes = append(node.Attributes, attribute)
		}
		return func(children ...Node) Node {
			node.children = children
			return node
		}
	}
}

var Html = makeNodeFunc("html")
var Head = makeNodeFunc("head")
var Title = makeNodeFunc("title")
var Body = makeNodeFunc("body")
var Meta = makeNodeFunc("meta")
var Span = makeNodeFunc("span")
var P = makeNodeFunc("p")
var Div = makeNodeFunc("div")
var Nav = makeNodeFunc("nav")
var Style = makeNodeFunc("style")
var Script = makeNodeFunc("script")
var H1 = makeNodeFunc("h1")
var H2 = makeNodeFunc("h2")
var H3 = makeNodeFunc("h3")
var H4 = makeNodeFunc("h4")
var H5 = makeNodeFunc("h5")
var Svg = makeNodeFunc("svg")
var Path = makeNodeFunc("path")
var Table = makeNodeFunc("table")
var Thead = makeNodeFunc("thead")
var Tbody = makeNodeFunc("tbody")
var Tr = makeNodeFunc("tr")
var Th = makeNodeFunc("th")
var Td = makeNodeFunc("td")
var Button = makeNodeFunc("button")
var A = makeNodeFunc("a")
var Input = makeNodeFunc("input")
var Img = makeNodeFunc("img")
var I = makeNodeFunc("i")
var Link = makeNodeFunc("link")
var Ul = makeNodeFunc("ul")
var Ol = makeNodeFunc("ol")
var Li = makeNodeFunc("li")
var Hr = makeNodeFunc("hr")
var Form = makeNodeFunc("form")
var Label = makeNodeFunc("label")
var Pre = makeNodeFunc("pre")
var Select = makeNodeFunc("select")
var Option = makeNodeFunc("option")
var Textarea = makeNodeFunc("textarea")
var Video = makeNodeFunc("video")
var Source = makeNodeFunc("source")
var Track = makeNodeFunc("track")
var Footer = makeNodeFunc("footer")
