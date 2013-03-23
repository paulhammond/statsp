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
	cleaner := statsp.NewCleaner()
	go statsp.Listen("127.0.0.1:8125", c)

	for {
		metric := <-c
		cleaned := cleaner.Clean(metric)
		fmt.Printf("%s %20s %2s %.10f %.10f\n", time.Now().Format(time.RFC3339), metric.Name, metric.Type, metric.Value, cleaned.Value)
	}
}
