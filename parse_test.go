package statsp

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in  string
		out []Metric
	}{
		// from Etsy statsd documentation:
		{"gorets:1|c", []Metric{
			Metric{"gorets", Counter, false, 1.0, 0},
		}},
		{"gorets:1|c|@0.1", []Metric{
			Metric{"gorets", Counter, false, 1.0, 0.1},
		}},
		{"glork:320|ms", []Metric{
			Metric{"glork", Timer, false, 320.0, 0},
		}},
		{"gaugor:333|g", []Metric{
			Metric{"gaugor", Guage, false, 333.0, 0},
		}},
		{"gaugor:-10|g", []Metric{
			Metric{"gaugor", Guage, true, -10.0, 0},
		}},
		{"gaugor:+4|g", []Metric{
			Metric{"gaugor", Guage, true, 4.0, 0},
		}},
		{"uniques:765|s", []Metric{
			Metric{"uniques", Set, false, 765.0, 0},
		}},
		{"gorets:1|c\nglork:320|ms\ngaugor:333|g\nuniques:765|s", []Metric{
			Metric{"gorets", Counter, false, 1.0, 0},
			Metric{"glork", Timer, false, 320.0, 0},
			Metric{"gaugor", Guage, false, 333.0, 0},
			Metric{"uniques", Set, false, 765.0, 0},
		}},

		// trailing newlines
		{"gorets:1|c\n", []Metric{
			Metric{"gorets", Counter, false, 1.0, 0},
		}},
		{"gorets:1|c\n\n\n", []Metric{
			Metric{"gorets", Counter, false, 1.0, 0},
		}},

		// based on github.com/b/statsd_spec
		{"glork:320|h", []Metric{
			Metric{"glork", Histogram, false, 320.0, 0},
		}},

		// Names
		{"blah.blah-dash_underscore.0.foo:333|g", []Metric{
			Metric{"blah.blah-dash_underscore.0.foo", Guage, false, 333.0, 0},
		}},
		{"blah.pipe|blah:333|g", []Metric{
			Metric{"blah.pipe|blah", Guage, false, 333.0, 0},
		}},
		{"snowman.\xe2\x98\x83:333|g", []Metric{
			Metric{"snowman.☃", Guage, false, 333.0, 0},
		}},
		{"blah.Iñtërnâtiônàlizætiøn:333|g", []Metric{
			Metric{"blah.Iñtërnâtiônàlizætiøn", Guage, false, 333.0, 0},
		}},

		// Numbers
		{"int:333|g", []Metric{Metric{"int", Guage, false, 333.0, 0}}},
		{"float:333.3|g", []Metric{Metric{"float", Guage, false, 333.3, 0}}},
		{"positive:+333.3|g", []Metric{Metric{"positive", Guage, true, 333.3, 0}}},
		{"negative:-333.3|g", []Metric{Metric{"negative", Guage, true, -333.3, 0}}},
	}

	for i, tst := range tests {
		b := []byte(tst.in)
		metrics, err := Parse(b)

		if err != nil {
			t.Errorf("%d unexpected error %v", i, err)
		}

		if !reflect.DeepEqual(metrics, &tst.out) {
			t.Errorf("%d expected\n%v\ngot\n%v\n", i, tst.out, metrics)
		}

	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name string
		in   string
		err  string
	}{
		{"missing name", ":1|c", "Invalid statsd packet"},
		{"missing :", "gorets1|c", "Invalid statsd packet"},
		{"unknown type", "glork:320|q", "Invalid statsd packet"},
		{"sample rate on non-counter", "gorets:1|g|@0.1", "Invalid statsd packet"},
		{"blank line", "gorets:1|c\n\ngorets:1|c", "Invalid statsd packet"},
	}
	for i, tst := range tests {
		b := []byte(tst.in)
		metrics, err := Parse(b)

		if err == nil || err.Error() != tst.err {
			t.Errorf("%d expected error %s, got %v", i, tst.err, err)
		}
		if metrics != nil {
			t.Errorf("%d expected nil metrics, got\n%v", i, metrics)
		}
	}

}
