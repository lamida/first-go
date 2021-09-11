// TODO
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type cache struct {
	c  map[string]bool
	mu sync.Mutex
}

func (c *cache) Seen(url string) bool {
	// c.mu.Lock()
	// defer c.mu.Unlock()
	if c.c == nil {
		c.c = make(map[string]bool)
	}
	if _, ok := c.c[url]; ok {
		return true
	}
	c.c[url] = true
	return false
}

func Crawl(url string, depth int, fetcher Fetcher, cache cache) {
	if depth <= 0 {
		return
	}

	if cache.Seen(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q]n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, cache)
	}
	return
}

func main() {
	done := make(chan bool, 1)
	go func() {
		fmt.Println("Workiing")
		Crawl("https://golang.org/", 4, fetcher, cache{})
		fmt.Println("dddone")
		done <- true
	}()
	// x, y, z := fetcher.Fetch("http://golang.org/")
	// fmt.Println(x, y, z)
	<-done
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	url  []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.url, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
