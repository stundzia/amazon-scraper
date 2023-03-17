package goexamples

func Amazon(url string, parse bool) ([]byte, error) {
	p := payload{
		Source: "amazon",
		Url:    url,
		Parse:  parse,
	}

	return doRealtimeRequest(p)
}
