package excercises

import (
	"fmt"
	"sync"
	"time"
)

var counter Counter

func RunExerciseWebCrawler() {
	fmt.Println(" * RunExerciseWebCrawler")
	fetchedUrls := FetchedUrls{cache: make(map[string]int)}
	counter = Counter{}
	counter.inc()
	go Crawl("http://golang.org/", 4, fetcher, &fetchedUrls)
	var val int
	for {
		val = counter.readValue()
		if val == 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("done")
}

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) inc() {
	c.mutex.Lock()
	c.value += 1
	c.mutex.Unlock()
}

func (c *Counter) dec() {
	c.mutex.Lock()
	c.value -= 1
	c.mutex.Unlock()
}

func (c *Counter) readValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

type FetchedUrls struct {
	cache map[string]int
	mutex sync.Mutex
}

func (f *FetchedUrls) isInCache(url string) bool {
	f.mutex.Lock()
	_, exists := f.cache[url]
	defer f.mutex.Unlock()
	return exists
}

func (f *FetchedUrls) addToCache(url string) {
	f.mutex.Lock()
	f.cache[url] = 1
	f.mutex.Unlock()
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetchedUrls *FetchedUrls) {
	defer counter.dec()
	if depth <= 0 {
		return
	}
	if fetchedUrls.isInCache(url) {
		fmt.Printf("cached: %s\n", url)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	fetchedUrls.addToCache(url)
	for _, u := range urls {
		counter.inc()
		go Crawl(u, depth-1, fetcher, fetchedUrls)
	}
	return
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
