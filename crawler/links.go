package crawler

import (
	"vezdecode-webcrawler/utils"

	"golang.org/x/net/html"
)

func (cr *Crawler) pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !utils.Contains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = cr.pageLinks(links, c)
	}
	return links
}
