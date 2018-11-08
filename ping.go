package main

import (
	//"fmt"
	"pkg/darwin_amd64"
	"io/ioutil"
	"log"
)

func main() {
	pinger, err := ping.NewPinger("my.chapman.edu")
	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics()

	error := ioutil.WriteFile("data.txt", []byte(printStats(stats)), 0666)
    if error != nil {
        log.Fatal(err)
    }
	//fmt.Printf("%v\n%v Packets Recieved\n%v Packets Sent\n%v Round Trip Times\n%v Standard Deviation for Round Trip Times\n", stats.Addr, stats.PacketsRecv, stats.PacketsSent, stats.Rtts, stats.StdDevRtt) // get send/receive/rtt stats
	//fmt.Printf("%v\n%v\n%v\n%v\n%v\n", stats.Addr, stats.PacketsRecv, stats.PacketsSent, stats.Rtts, stats.StdDevRtt)
}

func printStats(stats Statistics) String {
	return (stats.Addr"\n" +
		stats.PacketsRecv"\n" +
		stats.PacketsSent"\n" +
		stats.Rtts"\n" +
		stats.StdDevRtt"\n")

}
