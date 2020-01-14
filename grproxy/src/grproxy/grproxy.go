package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	a, _ := url.Parse("http://nginx:80/")
	// 	rprox := httputil.NewSingleHostReverseProxy(a)
	// 	rprox.ServeHTTP(w, r)
	// })

	u, _ := url.Parse("http://nginx:80/")
	http.Handle("/", httputil.NewSingleHostReverseProxy(u))

	gs1, _ := url.Parse("http://gserve1:80/")
	http.Handle("/library1", httputil.NewSingleHostReverseProxy(gs1))

	gs2, _ := url.Parse("http://gserve2:80/")
	http.Handle("/library2", httputil.NewSingleHostReverseProxy(gs2))

	// Start the server
	log.Fatal(http.ListenAndServe(":80", nil))
}
