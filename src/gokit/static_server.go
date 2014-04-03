package gokit

import (
  "net/http"
)

func StartStaticServer(port string, log chan string) {
  log <- "Starting server on" + port + "..."
  panic(http.ListenAndServe(port, http.FileServer(http.Dir("public/"))))
  log <- "done\n"
}
