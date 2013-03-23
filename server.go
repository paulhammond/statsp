package statsp

import (
	"log"
	"net"
)

// Listen creates a UDP server that parses statsd data into metrics and
// sends them over a channel.
func Listen(addr string, c chan Metric) {
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
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println("error: Failed to recieve packet", err)
		} else {
			metrics, err := Parse(buf[0:n])
			if err != nil {
				log.Println("error: Failed to recieve packet", err)
			}
			if metrics != nil {
				for _, p := range *metrics {
					c <- p
				}
			}
		}
	}
}
