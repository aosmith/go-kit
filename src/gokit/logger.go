package gokit

import (
  "fmt"
)

func BlockingLogger(logChannel chan string) {
  for ;; {
    log_message := <- logChannel
    fmt.Println(log_message)
  }
}
