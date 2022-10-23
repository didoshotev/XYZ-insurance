package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func BaseMiddlewareHandler(handler http.Handler) http.Handler {
	// adapter func
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before req
		fmt.Println("before handler middleware")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Println("after handler middleware; %s", time.Since(start))
		// after req
	})
}
