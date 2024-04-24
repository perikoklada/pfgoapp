package util

import (
	"log"
	"net/http"
	"net/url"
)

const (
	OriginHeaderName = "Origin"
)

func IsSameOrigin(r *http.Request) bool {

	var originHeaderSlice []string
	var ok bool
	if originHeaderSlice, ok = r.Header[OriginHeaderName]; !ok || len(originHeaderSlice) == 0 {
		log.Printf("origin header not present in the request")
		return false
	}

	originHeader := originHeaderSlice[0]
	var originURL *url.URL
	var err error
	if originURL, err = url.Parse(originHeader); err != nil {
		log.Printf("could not parse origin header into a valid url: %s\n", originHeader)
		return false
	}

	if originURL.Host != r.Host {
		log.Printf("CheckOrigin mismatch found. origin: %s, host: %s", originURL.Host, r.Host)
		return false
	}

	return true

}
