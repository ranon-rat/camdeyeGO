package controllers

import (
	"sync"
	"time"

	"github.com/go-ping/ping"
)

// use this in a go routine
func PingCameras() {

	for {
		var wg sync.WaitGroup
		for IP, ID := range camerasIPs {
			wg.Add(1)
			go func(ID, IP string) {
				defer wg.Done()

				pinger, err := ping.NewPinger(IP)
				if err != nil {

					delete(camerasNodesIceOffer, ID)
					delete(camerasIPs, IP)
					return
				}
				pinger.Count = 3
				pinger.Timeout = time.Duration(time.Second * 5) // i will wait 5 seconds if you have a bad internet its not my problem kek
				err = pinger.Run()                              // Blocks until finished.
				if err != nil {
					delete(camerasNodesIceOffer, ID)
					delete(camerasIPs, IP)
					return
				}
				stats := pinger.Statistics()
				if stats.PacketsRecv == 0 {
					delete(camerasNodesIceOffer, ID)
					delete(camerasIPs, IP)
					return
				}
			}(ID, IP)

		}
		wg.Wait()
		time.Sleep(time.Minute * 30)
	}

}
