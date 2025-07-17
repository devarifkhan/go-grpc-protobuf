package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	fmt.Println("Compression Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---Compression Middleware start---")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		// Set the header to indicate that the response is compressed
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer func(gz *gzip.Writer) {
			err := gz.Close()
			if err != nil {

			}
		}(gz)

		// Wrap the ResponseWriter to capture the status code
		w = &gzipResponseWriter{ResponseWriter: w, Writter: gz}

		next.ServeHTTP(w, r)
		fmt.Println("---Compression Middleware end---")
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writter *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writter.Write(b)
}
