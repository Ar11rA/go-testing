package async

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockHttpClient struct{}

func (m mockHttpClient) Get(url string, headers http.Header) (*http.Response, error) {
	t := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("123")),
	}

	return &t, nil
}

func TestFetchQuotes(t *testing.T) {
	m := mockHttpClient{}
	res := FetchQuotes(m)
	if res != "123" {
		t.Errorf("res should be 123, but is %v", res)
	}
}
