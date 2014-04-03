package main

import (
  "fmt"
  "net/http"
)

var LogChannel = make(chan string)

func main() {
  // Serve assets
  port := ":8080"
  fmt.Printf("Starting server on %s... ", port)
  go func() {
    panic(http.ListenAndServe(port, http.FileServer(http.Dir("public/"))))
  }()
  fmt.Printf("done\n")
  for ;; {
    log_message := <- LogChannel
    fmt.Println(log_message)
  }
}
