package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	pinger, err := ping.NewPinger("112.12.11.0")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Timeout = time.Duration(time.Second * 5)
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats, stats.PacketLoss, stats.PacketsSent, stats.PacketsRecv)
}
