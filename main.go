package main

import (
  "fmt"
  "net/http"
)

func routeHandle(rw http.ResponseWriter, r*http.Request) {
  fmt.Fprintf(rw, "Hello %s", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", routeHandle)
  http.ListenAndServe(":8080", nil)
}
