package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"testing"
)

// TestGetPublicIP retrieves this computer's public IP and ensures that it matches
// an IP regular expression
// NOTE: This test will fail if it cannot reach the address it's using to retrieve
// your IP. That address may change periodically and is hardcoded in source
func TestGetPublicIP(t *testing.T) {
	ip, err := GetPublicIP(http.DefaultClient)
	if err != nil {
		t.Fatalf("GetPublicIP failed: %s", err)
	}

	matched, err := regexp.Match(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`, []byte(ip))
	if err != nil {
		t.Fatalf("Regex matching failed: %s", err)
	}
	if !matched {
		t.Errorf("IP doesn't match regex: %s", ip)
		t.Fail()
	}
}

// TestLogPublicIpAddress retrieves this computer's public IP and ensures that it matches
// an IP regular expression embedded within the expected log message
// NOTE: This test will fail if it cannot reach the address it's using to retrieve
// your IP. That address may change periodically and is hardcoded in source
func TestLogPublicIpAddress(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "PRE: ", 0)
	LogPublicIpAddress(logger, 8080)

	matched, err := regexp.Match(`PRE: Listening on [0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}:8080\n`, buf.Bytes())
	if err != nil {
		t.Fatalf("Error matching logger output to regexp --\n%s", err)
	}
	if !matched {
		t.Errorf("Logger output didn't match expected pattern --\n%s", buf.String())
		t.Fail()
	}
}

// TestSimpleJSONResponse sets up a mocked server that uses SimpleJSONResponse to return
// a message. A request is made to the mocked server and the response is checked to
// ensure that it matches the expected output of SimpleJSONResponse
func TestSimpleJSONResponse(t *testing.T) {
	go func() {
		err := http.ListenAndServe(":50000", http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				SimpleJSONResponse(w, http.StatusOK, "Still alive")
			}))
		log.Fatal(err)
	}()

	res, err := http.Get("http://localhost:50000")
	if err != nil {
		t.Fatalf("Error on get request to mocked server --\n%s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Didn't receive expected status code (200 OK) from server: %s", res.Status)
		t.Fail()
	}

	var simpleJson struct {
		Status string
		Msg    string
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&simpleJson)
	if err != nil {
		t.Fatalf("Error decoding json object in server response --\n%s", err)
	}
	if simpleJson.Msg != "Still alive" {
		t.Errorf(`Didn't receive expected content ( {status: "200 OK", msg: "Still alive"} ) in json response: %+v`, simpleJson)
		t.Fail()
	}
}

// TestSimpleHttpResponse sets up a mocked server that uses SimpleHttpResponse to return
// a message. A request is made to the mocked server and the response is checked to
// ensure that it matches the expected output of SimpleHttpResponse
func TestSimpleHttpResponse(t *testing.T) {
	go func() {
		err := http.ListenAndServe(":40000", http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				SimpleHttpResponse(w, http.StatusOK, "Still alive")
			}))
		log.Fatal(err)
	}()

	res, err := http.Get("http://localhost:40000")
	if err != nil {
		t.Fatalf("Error on get request to mocked server --\n%s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Didn't receive expected status code (200 OK) from server: %s", res.Status)
		t.Fail()
	}

	var simpleJson struct {
		Status string
		Msg    string
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&simpleJson)
	if err != nil {
		t.Fatalf("Error decoding json object in server response --\n%s", err)
	}
	if simpleJson.Msg != "Still alive" {
		t.Errorf(`Didn't receive expected content ( {status: "200 OK", msg: "Still alive"} ) in json response: %+v`, simpleJson)
		t.Fail()
	}
}
