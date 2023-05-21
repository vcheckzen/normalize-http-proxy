package server

import (
	"fmt"
	"net"
)

func addrToIp(addr string) string {
	ips, ok := net.LookupIP(addr)
	if ok != nil {
		return addr
	}

	for _, ip := range ips {
		if v4 := ip.To4(); v4 != nil {
			return v4.String()
		}
	}

	return addr
}

func NewBaiduServer(address, listen string, verbose bool) *Server {
	ipPort := fmt.Sprintf("%s:443", addrToIp(address))
	upstream := NewUpstream(
		WithAddress(ipPort),
		WithHeaders(Headers{
			"Host":       ipPort,
			"X-T5-Auth":  "683556433",
			"User-Agent": "okhttp/3.11.0 Dalvik/2.1.0 (Linux; U; Android 11; Build/RP1A.200720.011) baiduboxapp/13.10.0.10 (Baidu; P1 11)",
		}),
	)
	return &Server{
		Listen:        listen,
		HttpUpstream:  upstream,
		HttpsUpstream: upstream,
		Verbose:       verbose,
	}
}

func NewUcServer(listen string, verbose bool) *Server {
	upstream := NewUpstream(
		WithAddress("101.71.140.5:8128"),
		WithHeaders(Headers{"Proxy-Authorization": "Basic dWMxMC4xMDMuMjcuMTgyOjFmNDdkM2VmNTNiMDM1NDQzNDUxYzdlZTc4NzNmZjM4=="}),
	)
	return &Server{
		Listen:        listen,
		HttpUpstream:  upstream,
		HttpsUpstream: upstream,
		Verbose:       verbose,
	}
}
