package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Unfortunately, maps are not compile-time constant
// So, this is a bit of overkill, but it's convenient
var statusMap map[int]string = map[int]string{
	http.StatusContinue:                      "100 Continue",
	http.StatusSwitchingProtocols:            "101 Switching Protocols",
	http.StatusProcessing:                    "102 Processing",
	http.StatusOK:                            "200 OK",
	http.StatusCreated:                       "201 Created",
	http.StatusAccepted:                      "202 Accepted",
	http.StatusNonAuthoritativeInfo:          "203 Non-Authoritative Info",
	http.StatusNoContent:                     "204 No Content",
	http.StatusResetContent:                  "205 Reset Content",
	http.StatusPartialContent:                "206 Partial Content",
	http.StatusMultiStatus:                   "207 Multiple Statuses",
	http.StatusAlreadyReported:               "208 Already Reported",
	http.StatusIMUsed:                        "226 IM Used",
	http.StatusMultipleChoices:               "300 Multiple Choices",
	http.StatusMovedPermanently:              "301 Moved Permanently",
	http.StatusFound:                         "302 Found",
	http.StatusSeeOther:                      "303 Found",
	http.StatusNotModified:                   "304 Not Modified",
	http.StatusUseProxy:                      "305 Use Proxy",
	http.StatusTemporaryRedirect:             "307 Temporary Redirect",
	http.StatusPermanentRedirect:             "308 Permanent Redirect",
	http.StatusBadRequest:                    "400 Bad Request",
	http.StatusUnauthorized:                  "401 Unauthorized",
	http.StatusPaymentRequired:               "402 Payment Required",
	http.StatusForbidden:                     "403 Forbidden",
	http.StatusNotFound:                      "404 Not Found",
	http.StatusMethodNotAllowed:              "405 Method Not Allowed",
	http.StatusNotAcceptable:                 "406 Not Acceptable",
	http.StatusProxyAuthRequired:             "407 Proxy Authorization Required",
	http.StatusRequestTimeout:                "408 Request Timeout",
	http.StatusConflict:                      "409 Conflict",
	http.StatusGone:                          "410 Gone",
	http.StatusLengthRequired:                "411 Length Required",
	http.StatusPreconditionFailed:            "412 Precondition Failed",
	http.StatusRequestEntityTooLarge:         "413 Request Entity Too Large",
	http.StatusRequestURITooLong:             "414 Request URI Too Long",
	http.StatusUnsupportedMediaType:          "415 Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable:  "416 Request Range Not Satisfiable",
	http.StatusExpectationFailed:             "417 Expectation Failed",
	http.StatusTeapot:                        "418 Short and Stout",
	http.StatusUnprocessableEntity:           "422 Unprocessaable Entity",
	http.StatusLocked:                        "423 Locked",
	http.StatusFailedDependency:              "424 Failed Dependency",
	http.StatusUpgradeRequired:               "426 Upgrade Required",
	http.StatusPreconditionRequired:          "428 Precondition Required",
	http.StatusTooManyRequests:               "429 Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge:   "431 Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:    "451 Unavailable For Legal Reasons",
	http.StatusInternalServerError:           "500 Internal Server Error",
	http.StatusNotImplemented:                "501 Not Implemented",
	http.StatusBadGateway:                    "502 Bad Gateway",
	http.StatusServiceUnavailable:            "503 Service Unavailable",
	http.StatusGatewayTimeout:                "504 Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "505 HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "506 Variant Also Negotiates",
	http.StatusInsufficientStorage:           "507 Insufficient Storage",
	http.StatusLoopDetected:                  "508 Loop Detected",
	http.StatusNotExtended:                   "510 Not Extended",
	http.StatusNetworkAuthenticationRequired: "511 Network Authentication Required",
}

func LogPublicIpAddress(l *log.Logger) error {
	// NOTE: This function is intended to print the public ip of a server, so an internet connection is assumed
	res, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		return err
	}

	ipBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if l != nil {
		l.Printf("Listening on %v\n", strings.TrimRight(string(ipBytes), "\n"))
	} else {
		// This will use the standard logger, which prints to os.Stderr
		log.Printf("Listening on %v\n", strings.TrimRight(string(ipBytes), "\n"))
	}
	return nil
}

func SimpleHttpResponse(w http.ResponseWriter, status int, msg string) {
	// This just returns a status header and a JSON object with a single field - "status"
	w.WriteHeader(status)
	io.WriteString(w, fmt.Sprintf(`{"status":"%s","msg":"%s"}`, statusMap[status], msg))
}
