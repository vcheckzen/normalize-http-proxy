package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vcheckzen/normalize-http-proxy/server"
)

var (
	baiduAddress string

	baiduCmd = &cobra.Command{
		Use:   "baidu",
		Short: "Using baidu as upstream server",
		Long: `Baidu (np baidu) will start a http proxy server, using baidu server as its upstream.

Examples:
  Normalize baidu server using default config.
	np baidu

  Provide another ip, the default port is 443. Https will use the same address and headers of http.
	np baidu -a 112.80.255.21

Baidu Servers:
  Domains
    cloudnproxy.baidu.com
    cloudnproxy.n.shifen.com

  China Unicom IPs
    Guangzhou 163.177.151.162
    Nanjing   112.80.255.21
    Beijing   123.125.142.40

  China Telecom IPs
    Guangzhou 14.215.177.73
    Nanjing   180.97.104.45
    Beijing   220.181.43.12

  China Mobile IPs
    Guangzhou 183.232.232.223
    Nanjing   36.152.45.80
    Beijing   112.34.116.40, 123.125.142.40
		`,
		Run: func(cmd *cobra.Command, args []string) {
			server.NewBaiduServer(baiduAddress, listen, verbose).Serve()
		},
	}
)

func init() {
	baiduCmd.Flags().StringVarP(&baiduAddress, "address", "a", "cloudnproxy.baidu.com", "custom proxy server address")
}
