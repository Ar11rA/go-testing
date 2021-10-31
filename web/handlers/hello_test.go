package handlers_test

import (
	"basics/web/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHelloService struct{}

func (m mockHelloService) GetGreeting() string {
	return "mock hello"
}

func TestHelloHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello/", nil)
	response := httptest.NewRecorder()
	mockHelloServiceStruct := mockHelloService{}
	helloHandler := handlers.NewHelloHander(mockHelloServiceStruct)
	handler := http.HandlerFunc(helloHandler.GetHello)
	handler.ServeHTTP(response, request)

	if response.Body.String() != "mock hello" {
		t.Fatalf("Non-expected response body %v:\tbody: %v", "mock hello", response.Body.String())
	}
	if response.Code != 200 {
		t.Fatalf("Non-expected status code %v:\tbody: %v", "200", response.Code)
	}
}
