package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("Listen failed, %s", err)
	}
	count := 0
	t1 := time.NewTimer(time.Second * 30)
	go func() {
		<-t1.C
		fmt.Println("Accepted connections", count)
		os.Exit(0)
	}()

	for {
		conn, err := ln.Accept()
		if err == nil {
			count++
		} else {
			log.Fatalf("Accept failed", err)
		}
		conn.Close()
	}
}
