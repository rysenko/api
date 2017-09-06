package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"fmt"
	"os"
	"strings"
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
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "NET_") {
			pair := strings.Split(e[4:], "=")
			pattern, target := "/"+pair[0]+"/", pair[1]
			fmt.Printf("Route %s to %s\n", pattern, target)
			http.HandleFunc(pattern, makeHandler(target))
		}
	}
	http.HandleFunc("/health", healthCheck)
	fmt.Printf("Listening port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
