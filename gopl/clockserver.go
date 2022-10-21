package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Local"
	}
	port := os.Args[1]

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, tz)
	}
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	for {
		t, err := TimeIn(time.Now(), tz)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(c, t.Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
