package cmd

import (
	"github.com/spf13/cobra"
)

var (
	storageCommand = &cobra.Command{
		Use: "storage [command]",
		Short: "storage configuration",
		Run: func(cmd *cobra.Command, args []string) {

		},

	}
	localCommand = &cobra.Command{
		Use: "local",
		Short: "local storage",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	aliyunCommand = &cobra.Command{
		Use: "aliyun-oss",
		Short: "aliyun oss storage",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)




func init() {


	storageCommand.AddCommand(localCommand)
	storageCommand.AddCommand(aliyunCommand)

	localCommand.Flags().String("path", ".", "local storage path")

	aliyunCommand.Flags().String("host", "", "aliyun oss host")
}

