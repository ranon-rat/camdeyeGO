package main

import "github.com/ranon-rat/camdeyeGO/clientNode/src/connection"

func main() {
	//defer settingMDNS.Setup().Close()
	//netScan.Setup()
	connection.ConnectToCameras()

}
