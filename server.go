package main

import (
  "gokit"
)

var LogChannel = make(chan string)

func main() {
  port := ":8080" // config
  // Next we use a go routine to serve static assets...
  go gokit.StartStaticServer(port, LogChannel)
  // This logger loops forever and prints log messages
  // to the console.  (Should be configurable)
  gokit.BlockingLogger(LogChannel)
}
