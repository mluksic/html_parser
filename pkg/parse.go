package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	linkNodes := findLinkNodes(doc)

	var links []Link
	for _, node := range linkNodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func getText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c)
	}

	return ret
}

func buildLink(node *html.Node) Link {
	var link Link

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
		}

		textArr := strings.Fields(getText(node))
		link.Text = strings.Join(textArr, " ")
	}

	return link
}

func findLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := findLinkNodes(c)
		ret = append(ret, node...)
	}

	return ret
}
