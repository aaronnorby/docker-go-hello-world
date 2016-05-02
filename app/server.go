// echoes back 'I am a server'
package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
)

const (
  host string = "localhost"
  port string = ":8000"
)

// var address string = host + port

var count int
var mu = &sync.Mutex{}

func counter() {
  mu.Lock()
  count++
  mu.Unlock()
}

func main() {
  http.HandleFunc("/count", handler)
  log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  counter()
  var nowCount int
  mu.Lock()
  nowCount = count
  mu.Unlock()
  fmt.Fprintf(w, "%s: %d", "I am a go server", nowCount)
}
