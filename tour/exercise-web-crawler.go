package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type fetchedURLLock struct {
	mu sync.Mutex
	v  map[string]bool
}

// Value returns the current value of the counter for the given key.
func (c *fetchedURLLock) GetAndSet(key string) bool {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	_, ok := c.v[key]
	c.v[key] = true
	defer c.mu.Unlock()
	return ok
}

var l = fetchedURLLock{v: make(map[string]bool)}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	ok := l.GetAndSet(url)
	if ok {
		fmt.Printf("%s has been fetched\n", url)
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	done := make(chan bool, len(urls))
	for _, u := range urls {
		//fmt.Printf("start to fetch %s\n", u)

		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)

		//Crawl(u, depth-1, fetcher)
	}
	for _, _ = range urls {
		<-done
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	//time.Sleep(10 * time.Second)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
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
