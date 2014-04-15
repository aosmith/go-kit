package main

import (
  "gokit"
)

var LogChannel = make(chan string)

func main() {
  web_port := ":8080"
  //api_port := ":8081"
  // Next we use a go routine to serve static assets...
  go gokit.StartStaticServer(web_port, LogChannel)
  //go gokit.StartObjectServer(api_port, LogChannel)
  // This logger loops forever and prints log messages
  // to the console.  (Should be configurable)
  gokit.BlockingLogger(LogChannel)
}
