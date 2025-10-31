package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time
		start := time.Now()
		fmt.Printf("Request Entered | PATH: %v | Method %v \n", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)

		fmt.Printf("Request completed  | PATH: %v | Method %v | Time : %v",
			r.URL.Path, r.Method, time.Since(start))

		//end time
	})
}
func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}
