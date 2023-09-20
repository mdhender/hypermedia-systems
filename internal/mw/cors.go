// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package mw

import (
	"log"
	"net/http"
	"strings"
)

// note: we could use the go-chi/cors middleware
//
//cors.Handler(cors.Options{
//	AllowedOrigins:   []string{"https://*", "http://*"},
//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
//	ExposedHeaders:   []string{"Link"},
//	AllowCredentials: false,
//	MaxAge:           300, // Maximum value not ignored by any of major browsers
//}))

// CORS will inject CORS headers on any OPTIONS request
func CORS(next http.Handler) http.Handler {
	log.Printf("[middleware] adding cors pre-flight middleware\n")

	allowHeaders := strings.Join([]string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"Content-Length",
		"Content-Type",
		"X-CSRF-Token",
	}, ", ")
	allowMethods := strings.Join([]string{
		"DELETE",
		"GET",
		"HEAD",
		"OPTIONS",
		"POST",
		"PUT",
	}, ", ")
	maxAge := "300" // max age not ignored by any of major browsers

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// inject CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		w.Header().Set("Access-Control-Max-Age", maxAge)

		// if we get the pre-flight request, return immediately
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
