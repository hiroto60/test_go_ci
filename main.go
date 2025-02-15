package main

import (
	"net/http"

	"github.com/hiroto60/test_dolce_ci/gen/greetconnect"
	"github.com/hiroto60/test_dolce_ci/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	greeter := handler.NewGreetHandler()
	path, handler := greetconnect.NewGreeterHandler(greeter)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8083",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
