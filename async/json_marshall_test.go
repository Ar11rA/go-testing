package async

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockHttpClientTwo struct{}

// mock exact JSON object this way
func (m mockHttpClientTwo) Get(url string, headers http.Header) (*http.Response, error) {
	t := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("{\"name\": \"aritra\"}")),
	}
	if strings.Contains(url, "1") {
		return &t, nil
	}
	return nil, errors.New("mock Error")
}

func TestFetchCharacterById_Success(t *testing.T) {
	var client mockHttpClientTwo
	char := FetchCharacterById(client, 1)
	if char.Name != "aritra" {
		t.Errorf("val is wrong but it is %v", char.Name)
	}
}

// panic testing
func TestFetchCharacterById_Failure(t *testing.T) {
	var client mockHttpClientTwo
	defer func() { recover() }()

	FetchCharacterById(client, 2)

	t.Errorf("Did not throw error")
}
