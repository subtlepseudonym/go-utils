package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// This is the address used to retrieve your public IP
const ipAPIAddress = "https://api.ipify.org?format=json"

type ipResponse struct {
	IP string `json:"ip"`
}

// GetPublicIP retrieves the public IP address of the computer where this function is run
// This function will return an error if there is no internet connection and, as of v2.0.0,
// depends upon ipify.org
func GetPublicIP(client *http.Client) (string, error) {
	req, err := http.NewRequest(http.MethodGet, ipAPIAddress, nil)
	if err != nil {
		return "", errors.Wrap(err, "build request failed")
	}

	c := http.DefaultClient
	if client != nil {
		c = client
	}

	res, err := c.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "do request failed")
	}

	if res == nil {
		return "", errors.New("response is nil")
	}
	decoder := json.NewDecoder(res.Body)
	var ipJSON ipResponse
	err = decoder.Decode(&ipJSON)
	if err != nil {
		return "", errors.Wrap(err, "decode ipResponse failed")
	}

	return ipJSON.IP, nil
}

// LogPublicIpAddress is a convenience function for getting your public IP and passing
// it directly to a logger with predefined formatting
func LogPublicIpAddress(l *log.Logger, port int) error {
	ip, err := GetPublicIP(http.DefaultClient)
	if err != nil {
		return errors.Wrap(err, "get public IP failed")
	}

	var portString string
	if port != 0 {
		portString = fmt.Sprintf(":%d", port)
	}
	if l != nil {
		l.Printf("Listening on %v%s\n", strings.TrimRight(ip, "\n"), portString)
	} else {
		// This will use the standard logger, which prints to os.Stderr
		log.Printf("Listening on %v%s\n", strings.TrimRight(ip, "\n"), portString)
	}
	return nil
}

// SimpleJSONResponse writes the provided HTTP status code back to the client, formats
// the message into a simple JSON object, and writes that back to the client
//
// 	ExampleSimpleJSONResponse() {
// 		func SomeHandler(w http.ResponseWriter, r *http.Request) {
// 			_, err := returnsErr()
// 			if err != nil {
// 				SimpleJSONResponse(w, http.StatusInternalServerError, "Something went wrong")
// 				return
// 			}
// 			// Output: {"msg":"Something went wrong"}
// 		}
// 	}
func SimpleJSONResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"msg":%q}`, message)
}

// SimpleHttpResponse writes the provided HTTP status code back to the client, formats
// the message into a simple JSON object, and writes that back to the client
// Deprecated: Use SimpleJSONResponse instead
func SimpleHttpResponse(w http.ResponseWriter, status int, msg string) {
	// This just returns a status header and a JSON object with a single field - "status"
	w.WriteHeader(status)
	io.WriteString(w, fmt.Sprintf(`{"status":"%s","msg":"%s"}`, http.StatusText(status), msg))
}
