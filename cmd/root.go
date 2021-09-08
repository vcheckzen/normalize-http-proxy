package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool
	listen  string

	rootCmd = &cobra.Command{
		Use:   "np",
		Short: "Normalize http proxy by adding needed headers.",
		Long: `Normalize HTTP Proxy is a http proxy server, which needs a upstream with several authentication headers.
By doing so, the upstream is converted to a plain proxy, no other parameters needed except for ip and port.

Examples:
  Start a server using baidu, or uc, or kingcard public server as its upstream.
    np baidu
    np uc
    np kingcard

  Start a server using a custom upstream, see
    np custom help
		`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "print every request")
	rootCmd.PersistentFlags().StringVarP(&listen, "listen", "l", "127.0.0.1:8888", "listen address")

	rootCmd.AddCommand(baiduCmd)
	rootCmd.AddCommand(ucCmd)
	rootCmd.AddCommand(kingcardCmd)
	rootCmd.AddCommand(customCmd)
}
