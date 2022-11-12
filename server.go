package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ln net.Listener) {
	for {
		fmt.Println("Listening for requests...")
		c, err := ln.Accept()
		fmt.Println("Connectection accepted")
		if err != nil {
			fmt.Println(err)
			continue
		}

		for {
		  var msg string
		  err = gob.NewDecoder(c).Decode(&msg)
		  if msg == "close" {
			  c.Close()
			  fmt.Println("Connectection closed")
			  break
		  }
		  if err != nil {
			  fmt.Println(err)
		  } else {
			  fmt.Println("Received", msg)
		  }

		  fmt.Println("33")
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	server(ln)
}
