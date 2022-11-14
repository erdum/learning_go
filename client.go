package main

import (
	"fmt"
	"encoding/gob"
	"net"
)

func client() {
  c, err := net.Dial("tcp", "127.0.0.1:9999")
  if err != nil {
    fmt.Println(err)
    return
  }

  for {  	
    var msg string
    fmt.Scanln(&msg)
    fmt.Println("Sending", msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
      fmt.Println(err)
    }
    if msg == "close" {
      break
    }
  }

  c.Close()
}

func main() {
	client()
}
