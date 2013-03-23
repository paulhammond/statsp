package statsp

import (
	"strconv"
)

type Metric struct {
	Name       string
	Type       MetricType
	Value      float64
	SampleRate float64
}

type MetricType int

const (
	_       = iota
	Counter = MetricType(iota)
	Timer
	Guage
	Histogram
	Set
)

var metric_type_lookup = map[string]MetricType{
	"c":  Counter,
	"ms": Timer,
	"g":  Guage,
	"h":  Histogram,
	"s":  Set,
}

func (m MetricType) String() string {
	for s, i := range metric_type_lookup {
		if i == m {
			return s
		}
	}
	return strconv.Itoa(int(m))
}
