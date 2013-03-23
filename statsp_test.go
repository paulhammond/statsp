package statsp

import (
	"testing"
)

func TestMetricTypeString(t *testing.T) {
	var s string

	s = Counter.String()
	if s != "c" {
		t.Errorf("expected counter string to be 'c', got '%s'", s)
	}

	s = Timer.String()
	if s != "ms" {
		t.Errorf("expected timer string to be 'ms', got '%s'", s)
	}

	s = Guage.String()
	if s != "g" {
		t.Errorf("expected gauge string to be 'g', got '%s'", s)
	}

	s = Histogram.String()
	if s != "h" {
		t.Errorf("expected histogram string to be 'h', got '%s'", s)
	}

	s = Set.String()
	if s != "s" {
		t.Errorf("expected set string to be 's', got '%s'", s)
	}
}
