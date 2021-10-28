package async

import (
	"testing"
)

func TestCheckLink_Valid(t *testing.T) {
	channel := make(chan string)
	go CheckLink("https://www.google.com", channel)
	res := <-channel
	if res != "The website https://www.google.com is up" {
		t.Errorf("val is wrong")
	}
}

func TestCheckLink_Invalid(t *testing.T) {
	channel := make(chan string)
	go CheckLink("http://abc", channel)
	res := <-channel
	if res != "The website http://abc is down" {
		t.Errorf("val is wrong")
	}
}
