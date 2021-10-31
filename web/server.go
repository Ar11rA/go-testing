package web

import (
	"fmt"
	"net/http"
)

func Start() {
	fmt.Println("Server started at 8080")
	err := http.ListenAndServe(":8080", Router())
	if err != nil {
		panic(err)
	}
}
