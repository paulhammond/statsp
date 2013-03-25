// Copyright 2013 Paul Hammond.
// This software is licensed under the MIT license, see LICENSE.txt for details.

package main

import (
	"fmt"
	"github.com/paulhammond/statsp"
	"time"
)

func main() {
	c := make(chan statsp.Packet)
	cleaner := statsp.NewCleaner()
	go statsp.Listen("127.0.0.1:8125", c, false)

	for {
		packet := <-c
		for i, metric := range *packet.Metrics {
			cleaned := cleaner.Clean(metric)
			newpacket := ""
			if (i == 0) {
				newpacket = "."
			}
			fmt.Printf("%2s %s %20s %2s %.10f %.10f\n", newpacket, packet.Time.Format(time.RFC3339), metric.Name, metric.Type, metric.Value, cleaned.Value)
		}
	}
}
