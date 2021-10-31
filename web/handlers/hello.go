package handlers

import (
	"basics/web/services"
	"fmt"
	"net/http"
)

type HelloHandler struct {
	helloService services.HelloService
}

func NewHelloHander(helloService services.HelloService) *HelloHandler {
	return &HelloHandler{
		helloService: helloService,
	}
}

func (h *HelloHandler) GetHello(w http.ResponseWriter, r *http.Request) {
	greeting := h.helloService.GetGreeting()
	fmt.Fprintf(w, greeting)
}
