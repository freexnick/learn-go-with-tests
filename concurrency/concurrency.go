package concurrency

type WebsiteChecker func(string) bool
type UrlResponses map[string]bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) UrlResponses {
	results := make(UrlResponses)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
