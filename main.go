package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"fmt"
)

func makeHandler(target string) func(w http.ResponseWriter, r *http.Request) {
	uri, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(uri)
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	const defaultPort = ":3000"

	http.HandleFunc("/account/", makeHandler("http://account"))
	http.HandleFunc("/contract/", makeHandler("http://contract"))
	http.HandleFunc("/health", healthCheck)
	http.ListenAndServe(defaultPort, nil)
}

