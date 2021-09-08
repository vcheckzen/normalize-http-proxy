package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/vcheckzen/normalize-http-proxy/log"
)

type Headers map[string]string

func (h Headers) String() string {
	str := "{"
	for k, v := range h {
		str = fmt.Sprintf("%s %s=%s,", str, k, v)
	}
	return fmt.Sprintf("%s %s", strings.TrimRight(str, ","), "}")
}

type Upstream struct {
	Address string
	Headers Headers
}

type UpstreamOption func(*Upstream)

func WithAddress(address string) UpstreamOption {
	return func(u *Upstream) {
		u.Address = fmt.Sprintf("http://%s", address)
	}
}

func WithHeaders(headers map[string]string) UpstreamOption {
	return func(u *Upstream) {
		u.Headers = headers
	}
}

func NewUpstream(opts ...UpstreamOption) *Upstream {
	u := new(Upstream)
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func (u *Upstream) setHeaders(r *http.Request) {
	for k, v := range u.Headers {
		r.Header.Set(k, v)
	}
}

type Server struct {
	Listen        string
	HttpUpstream  *Upstream
	HttpsUpstream *Upstream
	Verbose       bool
}

type ServerOption func(*Server)

func WithListen(listen string) ServerOption {
	return func(s *Server) {
		s.Listen = listen
	}
}

func WithHttpUpstream(upstream *Upstream) ServerOption {
	return func(s *Server) {
		s.HttpUpstream = upstream
	}
}

func WithHttpsUpstream(upstream *Upstream) ServerOption {
	return func(s *Server) {
		s.HttpsUpstream = upstream
	}
}

func WithVerbose(verbose bool) ServerOption {
	return func(s *Server) {
		s.Verbose = verbose

	}
}

func specify(s *Server, opts ...ServerOption) *Server {
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func NewServer(opts ...ServerOption) *Server {
	return specify(new(Server), opts...)
}

func (s *Server) dumpInfo() {
	fmt.Printf("Normalize HTTP Proxy: listens at %s.\n", s.Listen)
	fmt.Printf("Http upstream is %s, initial header(s) is(are):\n", s.HttpUpstream.Address)
	for k, v := range s.HttpUpstream.Headers {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Printf("Https upstream is %s, initial header(s) is(are):\n", s.HttpsUpstream.Address)
	for k, v := range s.HttpsUpstream.Headers {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Printf("Verbose output is %s.\n", map[bool]string{true: "on", false: "off"}[s.Verbose])
}

func (s *Server) Serve() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = s.Verbose
	log.SetVerbose(s.Verbose)

	// http
	proxy.Tr.Proxy = func(r *http.Request) (*url.URL, error) {
		s.HttpUpstream.setHeaders(r)
		return url.Parse(s.HttpUpstream.Address)
	}

	// https
	proxy.ConnectDial = proxy.NewConnectDialToProxyWithHandler(
		s.HttpsUpstream.Address, func(r *http.Request) { s.HttpsUpstream.setHeaders(r) })

	s.dumpInfo()
	log.Fatal(http.ListenAndServe(s.Listen, proxy))
}
