// Copyright 2013 Paul Hammond.
// This software is licensed under the MIT license, see LICENSE.txt for details.

package main

import (
	"fmt"
	"github.com/paulhammond/statsp"
	"time"
)

func main() {
	c := make(chan statsp.Metric)
	go statsp.Listen("127.0.0.1:8125", c)

	for {
		metric := <-c
		fmt.Printf("%s %20s %2s %v\n", time.Now().Format(time.RFC3339), metric.Name, metric.Type, metric.Value)
	}
}
