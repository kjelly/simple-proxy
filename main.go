package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func bypassPath(s string) bool {
	patterns := []string{".css", ".dart", ".js"}
	for i := range patterns {
		if strings.HasSuffix(s, patterns[i]) {
			return true
		}
	}
	return false

}

func hello(w http.ResponseWriter, r *http.Request) {
	path := ""
	if bypassPath(r.URL.Path) {
		path = r.URL.Path
		fmt.Printf("%s\n", path)
	}
	resp, err := http.Get("http://127.0.0.1:8080" + path)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	io.WriteString(w, string(body))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
