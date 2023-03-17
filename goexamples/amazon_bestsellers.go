package goexamples

func AmazonBestsellers(query string, domain string, startPage, categoryID int, parse bool) ([]byte, error) {
	p := payload{
		Source:    "amazon_bestsellers",
		Query:     query,
		Domain:    domain,
		StartPage: startPage,
		Context: []contextEntry{{
			Key:   "category_id",
			Value: categoryID,
		}},
		Parse: parse,
	}

	return doRealtimeRequest(p)
}
