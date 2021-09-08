package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vcheckzen/normalize-http-proxy/server"
)

var (
	httpAddress            string
	httpsAddress           string
	httpHeaders            map[string]string
	httpsHeaders           map[string]string
	useHttpAddressForHttps bool
	useHttpHeadersForHttps bool

	customCmd = &cobra.Command{
		Use:   "custom",
		Short: "Using a custom upstream",
		Long: `Custom (np custom) will start a http proxy server, using custom server as its upstream.

Examples:
  Start a server using custom upstream.
    np custom \
      -l 127.0.0.1:8888 \
      -a 112.80.255.21:80 \
      -x X-T5-Auth=ZjQxNDIh \
      -x other-http-header-name=other-http-header-value \
      -b 112.80.255.21:443 \
      -y X-T5-Auth=ZjQxNDIh \
      -y other-https-header-name=other-https-header-value \
      -v

  Your can add -c flag to reuse http address as https, and -d flag to reuse http headers as https.
    np custom \
      -a 112.80.255.21:80 \
      -x X-T5-Auth=ZjQxNDIh \
      -c \
      -d
		`,
		Run: func(cmd *cobra.Command, args []string) {
			if useHttpAddressForHttps {
				httpsAddress = httpAddress
			}
			if useHttpHeadersForHttps {
				httpsHeaders = httpHeaders
			}
			server.NewServer(
				server.WithListen(listen),
				server.WithVerbose(verbose),
				server.WithHttpUpstream(server.NewUpstream(server.WithAddress(httpAddress), server.WithHeaders(httpHeaders))),
				server.WithHttpsUpstream(server.NewUpstream(server.WithAddress(httpsAddress), server.WithHeaders(httpsHeaders))),
			).Serve()
		},
	}
)

func init() {
	customCmd.Flags().StringVarP(&httpAddress, "http-address", "a", "", "custom proxy server used for http (required)")
	customCmd.Flags().StringVarP(&httpsAddress, "https-address", "b", "", "custom proxy server used for https")
	customCmd.Flags().StringToStringVarP(&httpHeaders, "http-headers", "x", nil, "http authentication headers")
	customCmd.Flags().StringToStringVarP(&httpsHeaders, "https-headers", "y", nil, "https authentication headers")
	customCmd.Flags().BoolVarP(&useHttpAddressForHttps, "http-addr-for-https", "c", false, "use http address for https")
	customCmd.Flags().BoolVarP(&useHttpHeadersForHttps, "http-headers-for-https", "d", false, "use http headers for https")

	customCmd.MarkFlagRequired("http-address")
}
