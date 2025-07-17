package middlewares

import (
	"fmt"
	"net/http"
)

type HTTPOptions struct {
	CheckQuery                  bool
	CheckBody                   bool
	CheckBodyOnlyForContentType string
	Whitelist                   []string
}

func Hpp(options HTTPOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HPP Middleware")
			if options.CheckBody && r.Method == http.MethodGet && isCorrectContentType(r, options.CheckBodyOnlyForContentType) {
				filterBodyParams(r, options.Whitelist)
			}
			if options.CheckQuery && r.URL.Query() != nil {
				filterQueryParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	if contentType == "" {
		return true
	}
	return r.Header.Get("Content-Type") == contentType
}

func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Print("Error parsing form: ", err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0])
		}
		if !isWhitelisted(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func isWhitelisted(param string, whitelist []string) bool {
	for _, item := range whitelist {
		if item == param {
			return true
		}
	}
	return false
}

func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()
	for k, v := range query {

		if len(v) > 1 {
			//query.Set(k, v[0])
			query.Set(k, v[len(v)-1])
		}

		if !isWhitelisted(k, whitelist) {
			query.Del(k)
		}

	}
	r.URL.RawQuery = query.Encode()
}
