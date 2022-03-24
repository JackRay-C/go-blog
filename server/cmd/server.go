package cmd

import (
	"blog/core/storage"
	"github.com/spf13/cobra"
	"time"
)

var (
	serverCommand = &cobra.Command{
		Use: "server [command]",
		Short: "server command",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	startCommand = &cobra.Command{
		Use: "start",
		Short: "start server",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	stopCommand = &cobra.Command{
		Use: "stop",
		Short: "stop server",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	restartCommand = &cobra.Command{
		Use: "restart",
		Short: "restart server",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	serverPort int
	serverReadTimeout time.Duration
	serverWriteTimeout time.Duration
	uploadMaxSize storage.Size
)

func init() {
	rootCmd.AddCommand(serverCommand)

	serverCommand.AddCommand(startCommand)
	serverCommand.AddCommand(stopCommand)
	serverCommand.AddCommand(restartCommand)

	uploadMaxSize = 2* storage.GB

	startCommand.Flags().IntVar(&serverPort, "port", 8000, "server start port")
	startCommand.Flags().DurationVar(&serverReadTimeout, "read-timeout", 60, "server read timeout")
	startCommand.Flags().DurationVar(&serverWriteTimeout, "write-timeout", 60, "server wirte timeout")
	startCommand.Flags().Var(&uploadMaxSize, "upload-max-size", "allow upload file max size")

}
