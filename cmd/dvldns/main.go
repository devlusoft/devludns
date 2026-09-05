// Command dvldns is the authoritative DNS server for devludns.
package main

import (
	"log"
	"os"

	"devlusoft/devludns/internal/dnsserver"
	"devlusoft/devludns/internal/paths"
	"devlusoft/devludns/internal/store"

	"github.com/miekg/dns"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("dvldns: %v", err)
	}
}

func run() error {
	state := store.NewState()
	handler := dnsserver.NewHandler(state)

	udpServer := &dns.Server{
		Addr:    ":" + dnsPort(),
		Net:     "udp",
		Handler: handler,
	}
	tcpServer := &dns.Server{
		Addr:    ":" + dnsPort(),
		Net:     "tcp",
		Handler: handler,
	}

	go func() {
		if err := udpServer.ListenAndServe(); err != nil {
			log.Fatalf("UDP server: %v", err)
		}
	}()
	go func() {
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Fatalf("TCP server: %v", err)
		}
	}()

	log.Printf("dvldns listening on port %s (UDP+TCP)", dnsPort())
	<-make(chan struct{})
}

func dnsPort() string {
	if p := os.Getenv("DNS_PORT"); p != "" {
		return p
	}
	return "8053"
}
