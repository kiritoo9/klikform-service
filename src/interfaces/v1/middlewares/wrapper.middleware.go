package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Apply(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
