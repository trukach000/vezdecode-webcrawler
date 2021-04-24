package crawler

import "vezdecode-webcrawler/searcher"

type Crawler struct {
	recursive bool
	search    *searcher.Searcher
}

func NewCrawler(recursive bool, search *searcher.Searcher) *Crawler {
	return &Crawler{
		recursive: recursive,
		search:    search,
	}
}
