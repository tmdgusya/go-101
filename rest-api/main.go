package main

import (
	"net/http"
	"restapi/hello"
	"restapi/router"
)

func main() {
	r := router.GetRoute()
	hello.SetUpHelloRouter()
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
