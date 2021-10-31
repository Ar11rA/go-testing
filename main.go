package main

import (
	"basics/web"
	"fmt"
)

func main() {
	// fmt.Println("Heimdall HTTP client example")
	// timeout := 1000 * time.Millisecond
	// client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	// fmt.Println(async.FetchQuotes(client))
	// async.FetchCharacterById(client, 1)

	fmt.Println()
	fmt.Println("Web server")
	web.Start()
}
