package async

import (
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
}

func FetchQuotes(client HttpClient) string {
	res, err := client.Get("http://api.quotable.io/random/", nil)
	if err != nil {
		return ""
	}

	// Heimdall returns the standard *http.Response object
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
