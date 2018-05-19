package main

import (
  "fmt"
  "net/http"
)

func main() {
  // @ "/" endpoint, give this text/html Response TO (http) Request
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
  })

  // a request handler (line 10) cannot accept an http request without an http server listening
  // in this case, listening on port 80 or http://localhost/80
  http.ListenAndServe(":80", nil)
}
