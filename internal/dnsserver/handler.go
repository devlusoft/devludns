// Package dnsserver provides the DNS handler for devludns.
package dnsserver

import (
	"github.com/miekg/dns"
)

// Handler wraps a store.State and implements miekg/dns.Handler.
type Handler struct {
	state any // *store.State — filled in #3
}

// NewHandler returns a Handler backed by the given state.
func NewHandler(state any) dns.Handler {
	return &Handler{state: state}
}

// ServeDNS implements miekg/dns.Handler.
func (h *Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = true

	switch r.Question[0].Qtype {
	case dns.TypeA:
		if r.Question[0].Name == "example.test." {
			m.Answer = append(m.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: "example.test.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
				A:   []byte{127, 0, 0, 1},
			})
		}
	}

	w.WriteMsg(m)
}
