package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var ip, port, dialStr string
	if len(os.Args) == 3 {
		ip = os.Args[1]
		port = os.Args[2]
	}
	dialStr = ip + ":" + port
	connect := func() {
		for {
			conn, err := net.Dial("tcp", dialStr)
			if err == nil {
				conn.Close()
			} else {
				log.Fatalf("Dail error", err)
			}
		}
	}
	t1 := time.NewTimer(time.Second * 30)
	ch := make(chan bool)
	go func() {
		<-t1.C
		ch <- true
	}()
	go connect()
	go connect()
	go connect()
	<-ch
}
