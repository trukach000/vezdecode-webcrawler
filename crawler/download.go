package crawler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	gourl "net/url"
	"time"
	"vezdecode-webcrawler/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func (c *Crawler) GetAndParseURL(url string) (*html.Node, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	r, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't get web page by url: %s, error: %s", url, err)
	}
	defer r.Body.Close()

	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("can't parse by url: %s, error: %s", url, err)
	}
	return b, err
}

func (c *Crawler) renderNode(n *html.Node) (string, error) {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		return "", fmt.Errorf("can't render node: %s", err)
	}
	return buf.String(), nil
}

func (c *Crawler) analyzeURL(
	url string,
	baseurl string,
	visited *[]string,
) (
	result []AnalyzeResult,
	analyzeErr error,
) {
	result = make([]AnalyzeResult, 0)

	logrus.Debugf("Getting url: %s", url)

	page, err := c.GetAndParseURL(url)
	if err != nil {
		return result, fmt.Errorf("can't get and pase URL: %s", err)
	}

	pageContent, err := c.renderNode(page)
	if err != nil {
		return result, fmt.Errorf("can't render page from URL: %s, err:: %s", url, err)
	}

	searchResult := c.search.SearchAll(pageContent)

	for k, ra := range searchResult {
		for _, r := range ra {
			result = append(result, AnalyzeResult{
				EntityName: k,
				Value:      r,
			})
			logrus.Debugf("Search result: (entity: %s, value:%s)", k, r)

		}
	}
	*visited = append(*visited, url)
	if c.recursive {
		links := c.pageLinks(nil, page)
		baseUrlObj, err := gourl.Parse(baseurl)
		if err != nil {
			return result, fmt.Errorf("can't parse baseurl: %s, err: %s", baseurl, err)
		}
		for _, link := range links {
			linkObj, err := gourl.Parse(link)
			if err != nil {
				log.Printf("can't parse url: %s, err: %s", url, err)
				continue
			}
			if !utils.Contains(*visited, link) && linkObj.Hostname() == baseUrlObj.Hostname() {
				logrus.Debugf("analyze nested link: %s", link)
				r, err := c.analyzeURL(link, baseurl, visited)
				if err != nil {
					log.Printf("error during analyze links: %s", err)
					continue
				}
				result = append(result, r...)
			}
		}
	}

	return result, nil
}
