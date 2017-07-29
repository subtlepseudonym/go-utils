package http

import (
	"bytes"
	"log"
	"regexp"
	"testing"
)

func TestLogPublicIpAddress(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "PRE: ", 0)
	LogPublicIpAddress(logger)

	matched, err := regexp.Match(`PRE: Listening on [0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`, buf.Bytes())
	if err != nil {
		t.Fatalf("Error matching logger output to regexp --\n%s", err.Error())
	}
	if !matched {
		t.Fatalf("Logger output didn't match expected pattern --\n%s", buf.String())
	}
}
