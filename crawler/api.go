package crawler

import "github.com/sirupsen/logrus"

func (c *Crawler) AnalyzeURL(url string) (
	result []AnalyzeResult,
	visited []string,
	analyzeErr error,
) {
	visited = make([]string, 0)

	logrus.Debugf("Base url for analyze: %s", url)

	result, analyzeErr = c.analyzeURL(url, url, &visited)
	return result, visited, analyzeErr
}

func (c *Crawler) AnalyzeHtml(pageContent string) (
	result []AnalyzeResult,
	analyzeErr error,
) {
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

	return result, analyzeErr
}
