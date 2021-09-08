package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vcheckzen/normalize-http-proxy/server"
)

var (
	kingcardAddress string
	tokenApi        string

	kingcardCmd = &cobra.Command{
		Use:   "kingcard",
		Short: "Using kingcard as upstream server",
		Long: `Kingcard (np kingcard) will start a http proxy server, using kingcard server as its upstream.

Examples:
  Normalize kingcard server using default config.
    np kingcard

  Provide ip and token api. The default http port is 8090, https port is 8091. Https will use the same ip as http.
    np kingcard -a 157.255.173.182 -t http://kc.iikira.com/kingcard

Kingcard Servers:
  China Unicom IPs
    Shenzhen  157.255.173.182, 157.255.137.185
    Shanghai  116.128.146.115, 140.207.62.38, 210.22.247.193, 210.22.247.196
    Beijing   111.206.25.206, 123.126.122.24
    Tianjing  111.161.111.158, 125.39.52.35
    Chongqing 58.144.152.23, 58.144.152.199

Token APIs:
  http://kc.iikira.com/kingcard
  http://cs.xxzml.cn/k/get_tinyproxy_config.php
  http://api.saoml.com/qg.php
  http://api.saoml.com/qt.php
		`,
		Run: func(cmd *cobra.Command, args []string) {
			server.NewKingCardServerManager(server.WithTokenApi(tokenApi)).NewServer(kingcardAddress, listen, verbose).Serve()
		},
	}
)

func init() {
	kingcardCmd.Flags().StringVarP(&kingcardAddress, "address", "a", "210.22.247.196", "custom proxy server address")
	kingcardCmd.Flags().StringVarP(&tokenApi, "token-api", "t", "http://kc.iikira.com/kingcard", "token api")
}
