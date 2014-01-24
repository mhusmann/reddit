package main

import (
  "log"
  "github.com/mhusmann/reddit"
  "fmt"
)

func main() {
  items, err := reddit.Get("golang")
  if err != nil {
    log.Fatal(err)
  }
  for _, item := range items {
    fmt.Println(item)
  }
}