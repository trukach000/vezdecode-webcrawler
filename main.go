package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
	"vezdecode-webcrawler/crawler"
	"vezdecode-webcrawler/searcher"
	"vezdecode-webcrawler/searcher/entities"

	"github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().Unix())

	//inputFileFlag := flag.String("file", "", "file with links to analyze")
	inputURLFlag := flag.String("url", "", "url to analyze")
	inputFileFlag := flag.String("file", "", "file with urls to analyze")

	recursiveFlag := flag.Bool("r", false, "use recursive search, domain restricted")
	debugFlag := flag.Bool("d", false, "show debug information")

	flag.Parse()

	if *debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	phoneSearch, err := entities.NewSearchPhone()
	if err != nil {
		fmt.Printf("can't create search phone entity: %s\n", err)
		os.Exit(1)
	}

	ipSearch, err := entities.NewSearchIP()
	if err != nil {
		fmt.Printf("can't create search ip entity: %s\n", err)
		os.Exit(1)
	}

	socialSearch, err := entities.NewSearchSocial()
	if err != nil {
		fmt.Printf("can't create search social entity: %s\n", err)
		os.Exit(1)
	}

	emailSearch, err := entities.NewSearchEmail()
	if err != nil {
		fmt.Printf("can't create search email entity: %s\n", err)
		os.Exit(1)
	}

	search := searcher.NewSearcher(phoneSearch, ipSearch, socialSearch, emailSearch)
	cr := crawler.NewCrawler(*recursiveFlag, search)

	res := make([]crawler.AnalyzeResult, 0)
	visited := make([]string, 0)
	if *inputURLFlag != "" {
		// analyze URL
		var err error
		res, visited, err = cr.AnalyzeURL(*inputURLFlag)
		if err != nil {
			fmt.Printf("analyze error: %s\n", err)
			os.Exit(1)
		}
		printResults(visited, res)
	} else if *inputFileFlag != "" {
		IterateFile(*inputFileFlag, cr)
	} else {
		// analyze input HTML data from first argument
		html := flag.Arg(0)
		if html == "" {
			fmt.Println("You should pass html as data for parsing")
			os.Exit(1)
		}
		var err error
		res, err = cr.AnalyzeHtml(html)
		if err != nil {
			fmt.Printf("analyze error: %s", err)
			os.Exit(1)
		}
		printResults(visited, res)
	}

}

func printResults(visited []string, res []crawler.AnalyzeResult) {
	fmt.Printf("RESULT:\n Visited links (%d):\n", len(visited))
	for _, vl := range visited {
		fmt.Printf("ulr: %s\n", vl)
	}

	fmt.Printf("Founded entities:\n")
	for _, entity := range res {
		fmt.Printf("entity name: %s , value: %s\n", entity.EntityName, entity.Value)
	}
}

func IterateFile(filename string, cr *crawler.Crawler) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot read the file: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	results := make([]crawler.AnalyzeResult, 0)
	visited := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		url := scanner.Text()
		res, vis, err := cr.AnalyzeURL(url)
		if err != nil {
			fmt.Printf("analyze error for base URL: %s, err: %s\n", url, err)
		}
		visited = append(visited, vis...)
		results = append(results, res...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printResults(visited, results)

}
