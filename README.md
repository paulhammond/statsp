# StatsP

This is a [Go][go] parser for the [statsd][statsd] [wire protocol][protocol].

Unlike most other implementations it does **not** perform aggregation or send
the values to a metrics store, it just parses values and does a small amount
of normalization. If you need aggregation or storage a [github search for
"statsd"](https://github.com/search?q=statsd) will list many alternatives.

## Installation

Install with `go get`:

    go get github.com/paulhammond/statsp

## Usage

StatsP will parse a slice of bytes and return a slice of Metric values:

    b := []byte("gorets:1|c")
    metrics, err := statsp.Parse(b)
    fmt.Println(metrics[0].Name)       // "gorets"
    fmt.Println(metrics[0].Value)      // 1.0
    fmt.Println(metrics[0].Type)       // statsp.Counter

The most common usage is to read StatsD data directly from the network; a
basic server implementation is provided that sends received metrics on a
channel.

    c := make(chan statsp.Packet)
    go statsp.Listen("127.0.0.1:8125", c)
    for {
      packet := <-c
      // do something with the metric here
    }

Even though StatsP doesn't perform aggregation, the Gauge type allows for both
relative and absolute values, so needs some processing. `statsd.Cleaner`
provides an implementation of this:

    cleaner := statsp.NewCleaner()

    b := []byte("foo:1|g\nfoo:+1|g\nfoo:-1|g\nfoo:-1|g\n")
    metrics, err := statsp.Parse(b)

    for _, metric := range(metrics) {
      cleaned := cleaner.Clean(metric)
      fmt.Println(metric.value, cleaned.value)
    }
    // Output:
    // 1.0 1.0
    // 1.0 2.0
    // -1.0 1.0
    // -1.0 0.0

An example of using both is provided in [statsp-example](statsp-example/statsp-example.go)

## References

  * [go][go]
  * [statsd][statsd]
  * [stats protocol][protocol]
  * [metricsd][metricsd]

[go]: http://golang.org/
[statsd]: https://github.com/etsy/statsd
[protocol]: https://github.com/b/statsd_spec
[metricsd]: https://github.com/mojodna/metricsd

## License

Copyright (c) 2013 Paul Hammond. gocollectd is available under the MIT
license, see [LICENSE.txt](LICENSE.txt) for details