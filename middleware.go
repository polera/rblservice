package main

import (
	"net/http"
)

func JSONResponse(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
