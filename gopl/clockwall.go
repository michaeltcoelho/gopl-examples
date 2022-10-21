package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

type TZ struct {
	Name   string
	Server string
	Time   string
}

func NewTZ(tz string) *TZ {
	address := strings.Split(tz, "=")
	return &TZ{
		Name:   address[0],
		Server: address[1],
	}
}

type ClockWall map[string]*TZ

func main() {
	tzs := os.Args[1:]
	clockWall := make(ClockWall)

	for _, tzArg := range tzs {
		tz := NewTZ(tzArg)

		clockWall[tz.Name] = tz

		go handleTZ(tz)
	}

	printWall(clockWall)
}

func TimeIn(t time.Time, tz string) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func handleTZ(tz *TZ) {
	for {
		conn, err := net.Dial("tcp", tz.Server)
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 8)
		length, err := conn.Read(buf)
		if err != nil {
			log.Print(err)
			continue

		}
		tz.Time = string(buf[:length])

		conn.Close()

		time.Sleep(1 * time.Second)
	}
}

func printWall(clockWall ClockWall) {
	const tabFormat = "%v\t%v\t%v\t\n"

	var tzs []string
	for k, _ := range clockWall {
		tzs = append(tzs, k)
	}
	sort.Strings(tzs)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	for {
		clearScreen()
		fmt.Fprintf(tw, tabFormat, "Location", "Time", "Server")
		fmt.Fprintf(tw, tabFormat, "--------", "--------", "--------")
		for _, tz := range tzs {
			fmt.Fprintf(tw, tabFormat, clockWall[tz].Name, clockWall[tz].Time, clockWall[tz].Server)
		}
		tw.Flush()
		time.Sleep(1 * time.Second)
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[0;0H")
}
