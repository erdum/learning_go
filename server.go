package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ln net.Listener) {
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		var msg string
		err = gob.NewDecoder(c).Decode(&msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Received", msg)
		}
		c.Close()
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	server(ln)
}
