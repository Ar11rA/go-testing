package services_test

import (
	"basics/web/services"
	"testing"
)

func TestGetGreeting(t *testing.T) {
	helloService := services.NewHelloService()
	res := helloService.GetGreeting()
	if res != "wello" {
		t.Fatalf("Non-expected response body %v:\tbody: %v", "wello", res)
	}
}
