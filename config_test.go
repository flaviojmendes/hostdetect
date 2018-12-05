package main

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"

	"github.com/andreyvit/diff"
	"gopkg.in/yaml.v2"
)

var textConfig = `timeout: 1s
interval: 1m0s
hosts:
- network: tcp
  address: www.google.com:80
- network: tcp
  address: www.google.com:23
- network: tcp
  address: zebeleo.com.br:80
`

var objConfig = Config{
	Interval: time.Minute,
	Timeout:  time.Second,
	Hosts: []Host{
		googleHttp,
		googleTelnet,
		bogusHost,
	},
}

func TestConfigCanBeSerialized(t *testing.T) {
	d, err := yaml.Marshal(objConfig)
	if err != nil {
		t.Fatalf("error serializing config: %s", err)
	}

	if string(d) != textConfig {
		t.Fatalf("%v", diff.CharacterDiff(textConfig, string(d)))
	}
}

func TestConfigCanBeDeserialized(t *testing.T) {
	var c Config
	yaml.Unmarshal([]byte(textConfig), &c)
	if !reflect.DeepEqual(c, objConfig) {
		t.Fatalf("got: %s expected: %s", c, objConfig)
	}
}

func TestIfItReadConfig(t *testing.T) {
	content, err := ioutil.ReadFile("mock_file")

	if err != nil {
		t.Fatalf("error reading file: mock_file")
	}

	var result = readConfig(content)
	var expected = 3

	if len(result.Hosts) != expected {
		t.Fatalf("got: %v expected: %v", len(result.Hosts), expected)
	}
}

func TestIfFileIsReadableOnceItExists(t *testing.T) {
	var filepath = "mock_file"

	result, err := readConfigFile(filepath)

	if result == nil && err != nil {
		t.Fatalf("error serializing config: %s", filepath)
	}
}

func TestCheckMetricContentForOpenPort(t *testing.T) {
	var host = Host{Network: "tcp", Address: "zebeleu.com.br:80"}
	var check = Check{host: host, state: open}

	var response = check.getMetric()

	var expected = "host_with_failure{host=\"zebeleu.com.br:80\", network=\"tcp\", state=\"0\"} 1"

	if response != expected {
		t.Fatalf("got: %v expected: %v", response, expected)
	}
}

func TestCheckMetricContentForClosedPort(t *testing.T) {
	var host = Host{Network: "tcp", Address: "zebeleu.com.br:80"}
	var check = Check{host: host, state: closed}

	var response = check.getMetric()

	var expected = "host_with_success{host=\"zebeleu.com.br:80\", network=\"tcp\", state=\"1\"} 1"

	if response != expected {
		t.Fatalf("got: %v expected: %v", response, expected)
	}
}

func TestCheckMetricContentForFailedPort(t *testing.T) {
	var host = Host{Network: "tcp", Address: "zebeleu.com.br:80"}
	var check = Check{host: host, state: failure}

	var response = check.getMetric()

	var expected = "host_with_failure{host=\"zebeleu.com.br:80\", network=\"tcp\", state=\"2\"} 1"

	if response != expected {
		t.Fatalf("got: %v expected: %v", response, expected)
	}
}

func TestCheckMetrics(t *testing.T) {
	var host = Host{Network: "tcp", Address: "zebeleu.com.br:80"}
	var check = Check{host: host, state: failure}

	checks = append(checks, check)

	var response = getMetrics()

	var expected = "hosts_with_failures 1"

	if response != expected {
		t.Fatalf("got: %v expected: %v", response, expected)
	}
}
