package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
		l.Printf("Listening on %v%n", strings.TrimRight(string(ipBytes), "\n"))
	} else {
		// This will use the standard logger, which prints to os.Stderr
		log.Printf("Listening on %v%n", strings.TrimRight(string(ipBytes), "\n"))
	}
	return nil
}
