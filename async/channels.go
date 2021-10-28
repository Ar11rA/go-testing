package async

import (
	"fmt"
	"net/http"
)

func CheckLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("The website %v is down", url)
	}
	c <- fmt.Sprintf("The website %v is up", url)
}
