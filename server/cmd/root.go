package cmd

import (
	"github.com/spf13/cobra"
)

var (

	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "go-blog [global options] command [command options] [arguments...]",
		Short: "go blog is an elegant blog based on GoLand and Vue",
		Long: `go blog is an elegant blog based on GoLand and Vue
	
Find more information at: https://kubernetes.io/docs/reference/kubectl/overview`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

const (
	DefaultConfigFile = "conf/default.yaml"
)

func Execute() error {
	_, err := rootCmd.ExecuteC()
	return err
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", DefaultConfigFile, "test configuration file")
	rootCmd.PersistentFlags().String("author", "renhj", "author name for copyright attribution")
	v.SetDefault("app-author", rootCmd.PersistentFlags().Lookup("author").Value)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCommand)
	rootCmd.AddCommand(serverCommand)
	cobra.OnInitialize()
}

