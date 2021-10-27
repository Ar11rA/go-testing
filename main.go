package main

import (
	"basics/async"
	"fmt"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

func main() {
	fmt.Println("hi")
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	fmt.Println(async.FetchQuotes(client))
}
