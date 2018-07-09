package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var port string

func excludeLast(arr []string) []string {
	return arr[:len(arr)-1]
}

func formatAddr(addr string) string {
	return strings.Join(excludeLast(strings.Split(addr, ":")), ":")
}

func getIP(w http.ResponseWriter, r *http.Request) {
        addr := ""
        if r.Header.Get("x-real-ip") != "" {
            addr = r.Header.Get("x-real-ip")
	} else {
	    addr = formatAddr(r.RemoteAddr)
	}

	log.Printf("%s request from %s", r.Method, addr)

	w.Write([]byte(addr))
}

func main() {
	flag.StringVar(&port, "port", "8080", "port number on which to serve")
	flag.Parse()

	add := fmt.Sprintf("0.0.0.0:%s", port)
	srv := http.NewServeMux()

	srv.HandleFunc("/", getIP)

	log.Printf("serving on %s", add)
	if err := http.ListenAndServe(add, srv); err != nil {
		log.Fatalf("server stopped: %v\n", err)
	}
}
