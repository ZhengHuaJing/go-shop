package util

import (
	"net"
	"net/http"
)

func RemoteIP(req *http.Request) string {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	return host
}
