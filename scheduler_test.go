package main

import (
	"testing"
	"time"
)

func TestIntervalForEvery5Minutes(t *testing.T) {
	var objConfig = &Config{
		Interval: 5 * time.Minute,
		Timeout:  time.Second,
		Hosts: []Host{
			googleHttp,
			googleTelnet,
			bogusHost,
		},
	}

	var result = getInterval(objConfig)
	var expected = "*/300 * * * * *"

	if result != expected {
		t.Fatalf("got %v, expected %v", result, expected)
	}
}

func TestIntervalForEvery10Seconds(t *testing.T) {
	var objConfig = &Config{
		Interval: 5 * time.Second,
		Timeout:  time.Second,
		Hosts: []Host{
			googleHttp,
			googleTelnet,
			bogusHost,
		},
	}

	var result = getInterval(objConfig)
	var expected = "*/5 * * * * *"

	if result != expected {
		t.Fatalf("got %v, expected %v", result, expected)
	}
}
