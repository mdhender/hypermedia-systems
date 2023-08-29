// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package server

import (
	"log"
	"net/http"
	"unicode"
)

func optionsCors(next http.Handler) http.Handler {
	log.Printf("[server] adding cors middleware\n")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// inject CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// handle CORS
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleBadRunes(next http.Handler) http.Handler {
	log.Printf("[server] adding bad runes middleware\n")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("[server] running bad runes middleware\n")
		// return an error if the URL contains any non-printable runes.
		for _, ch := range r.URL.Path {
			if !unicode.IsPrint(ch) {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
