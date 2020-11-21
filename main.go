package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if strings.TrimSpace(path) == "" {
		http.Error(w, "tinyrev usage: tinyrev-host/<upstream-url>", http.StatusBadRequest)
		return
	}

	upstream, err := url.Parse("https://" + r.URL.Path[1:])
	if err != nil {
		http.Error(w, "tinyrev usage: tinyrev-host/<upstream-url>", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(upstream.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
