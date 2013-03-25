package statsp

import (
	"log"
	"net"
	"time"
)

type Packet struct {
	Metrics *[]Metric
	Addr    net.Addr
	Time    time.Time
}

// Listen creates a UDP server that parses statsd data into metrics and
// sends them over a channel.
func Listen(addr string, c chan Packet, clean bool) {
	cleaner := NewCleaner()
	laddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalln("fatal: failed to resolve address", err)
	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatalln("fatal: failed to listen", err)
	}
	for {
		buf := make([]byte, 1452)
		n, raddr, err := conn.ReadFrom(buf[:])
		t := time.Now().UTC()
		if err != nil {
			log.Println("error: Failed to recieve packet", err)
		} else {
			metrics, err := Parse(buf[0:n])
			if err != nil {
				log.Println("error: Failed to recieve packet", err)
			}
			if metrics != nil {
				var p Packet
				if clean {
					cleaned := cleaner.CleanMetrics(*metrics)
					p = Packet{&cleaned, raddr, t}
				} else {
					p = Packet{metrics, raddr, t}
				}

				c <- p
			}
		}
	}
}
