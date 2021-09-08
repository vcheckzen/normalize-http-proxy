package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vcheckzen/normalize-http-proxy/server"
)

var (
	ucCmd = &cobra.Command{
		Use:   "uc",
		Short: "Using uc as upstream server",
		Long:  `Uc (np uc) will start a http proxy server, using uc server as its upstream.`,
		Run: func(cmd *cobra.Command, args []string) {
			server.NewUcServer(listen, verbose).Serve()
		},
	}
)
