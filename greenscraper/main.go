package main

import (
	"fmt"
	"github.com/cncf-tags/tag-env-tooling/greenscraper/cmd"
	"regexp"
	"sync"
)

func main() {
	keywords, err := cmd.ReadLinesFromFile("keywords.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to read keywords from file: %v", err))
	}

	urls, err := cmd.ReadLinesFromFile("urls.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to read URLs from file: %v", err))
	}
	// Read exclusions from exclusions.txt
	exclusions, err := cmd.ReadLinesFromFile("exclusions.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to read exclusions from file: %v", err))
	}

	titleRegex := regexp.MustCompile(".*'>(.*?)<span class=\"vs\">.*")
	keywordRegexes := make([]*regexp.Regexp, len(keywords))
	for i, keyword := range keywords {
		keywordRegexes[i] = regexp.MustCompile(keyword)
	}

	const concurrentLimit = 5
	sem := make(chan struct{}, concurrentLimit)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}
		go func(u string) {
			cmd.ProcessURL(u, keywordRegexes, exclusions, titleRegex, &wg)
			<-sem
		}(url)
	}

	wg.Wait()
}
