# devludns

Authoritative DNS server built with Go and [miekg/dns](https://github.com/miekg/dns).

## What

`devludns` is a greenfield authoritative DNS server for devlusoft infrastructure. It answers DNS queries for configured zones, persists state to SQLite, and supports multi-tenant operation.

## Build

```bash
go build ./cmd/dvldns
```

## Run

```bash
./dvldns
```

The server listens on port **8053** (UDP and TCP) by default. Override with `DNS_PORT` env var.

## Test

```bash
dig @localhost -p 8053 example.test A
```

Expected response: `127.0.0.1`

## Structure

```
cmd/dvldns/          main entry point
internal/dnsserver/  DNS handler (miekg/dns.Handler)
internal/store/      state management (SQLite)
internal/paths/      default paths and constants
proto/wire/          DNS wire-format types
```

## Status

- [ ] Issue #1 — you are here (bootstrap)
- [ ] Issues #2–#20 — see [github.com/devlusoft/devludns](https://github.com/devlusoft/devludns)
