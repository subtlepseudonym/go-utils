package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"testing"
)

func TestLogPublicIpAddress(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "PRE: ", 0)
	LogPublicIpAddress(logger)

	matched, err := regexp.Match(`PRE: Listening on [0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\n`, buf.Bytes())
	if err != nil {
		t.Fatalf("Error matching logger output to regexp --\n%s", err.Error())
	}
	if !matched {
		t.Fatalf("Logger output didn't match expected pattern --\n%s", buf.String())
	}
}

func TestSimpleHttpResponse(t *testing.T) {
	go func() {
		log.Fatal(http.ListenAndServe(":40000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { SimpleHttpResponse(w, http.StatusOK, "Still alive") })))
	}()

	res, err := http.Get("http://localhost:40000")
	if err != nil {
		t.Fatalf("Error on get request to mocked server --\n%s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Didn't receive expected status code (200 OK) from server: %s", res.Status)
	}

	var simpleJson struct {
		Status string
		Msg    string
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&simpleJson)
	if err != nil {
		t.Fatalf("Error decoding json object in server response --\n%s", err.Error())
	}
	if simpleJson.Status != "200 OK" || simpleJson.Msg != "Still alive" {
		t.Fatalf(`Didn't receive expected content ( {status: "200 OK", msg: "Still alive"} ) in json response: %+v`, simpleJson)
	}
}
