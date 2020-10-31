package mws

import "net"

//net/http/transport.go:2787
//type tlsHandshakeTimeoutError struct{}
//
//func (tlsHandshakeTimeoutError) Timeout() bool   { return true }
//func (tlsHandshakeTimeoutError) Temporary() bool { return true }
//func (tlsHandshakeTimeoutError) Error() string   { return "net/http: TLS handshake timeout" }
//

func isNetError(err error) (net.Error, bool) {
	netEx, ok := err.(net.Error)
	return netEx, ok
}

func isTimeoutError(err error) bool {
	if netEx, ok := isNetError(err); ok {
		return netEx.Timeout()
	}
	return false
}
