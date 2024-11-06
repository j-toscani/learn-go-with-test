package concurrency

type WbsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WbsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// Pass url as a parameter so that each call has it´s own copy
		// Otherwise all goroutines would share the same url
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
