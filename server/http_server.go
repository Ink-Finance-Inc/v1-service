package http

import (
	"net"
	"net/http"

	"inkfinance.xyz/httprouter"
)

var mux = http.NewServeMux()

var httpServer = http.Server{Handler: mux}

// http run
func Run(addr string, errc chan error) {

	// Register router
	httprouter.RegisterRouter(mux)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- httpServer.Serve(lis)

	
}
