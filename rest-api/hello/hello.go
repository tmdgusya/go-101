package hello

import (
	"net/http"
	"restapi/router"
)

func SetUpHelloRouter() {
	var route = router.GetRoute()

	route.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello!, World!"))
		if err != nil {
			return
		}
	})
}
