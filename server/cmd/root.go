package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	storageType string
	rootCmd = &cobra.Command{
		Use: "go-blog",
		Short: "go blog is an elegant blog based on GoLand and Vue",
		Long: `go blog is an elegant blog based on GoLand and Vue
	
Find more information at: https://kubernetes.io/docs/reference/kubectl/overview`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() error {
	_, err := rootCmd.ExecuteC()
	return err
}

func init()  {
	cobra.OnInitialize(initConfig)

	initCobraCommand()
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c","", "base configuration file (default is $HOME/.go-blog.yaml)")
	rootCmd.PersistentFlags().StringVarP(&storageType, "storage", "s", "local", "storage type [local|aliyun-oss|tencent-ss|qiniu]")

}

func initCobraCommand()  {

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(dbCommand)

	rootCmd.AddCommand(logCommand)

	rootCmd.AddCommand(initCommand)

	rootCmd.AddCommand(storageCommand)

	rootCmd.AddCommand(serverCommand)


	initCommand.AddCommand(dbCommand)
	initCommand.AddCommand(storageCommand)
	initCommand.AddCommand(serverCommand)
	initCommand.AddCommand(logCommand)

}

func initConfig()  {

}

//func initConfig()  {
//	if cfgFile != "" {
//		viper.SetConfigFile(cfgFile)
//	} else {
//		home, err := os.UserHomeDir()
//		cobra.CheckErr(err)
//
//		// 1、判断是否存在配置文件
//		cfgFile = path.Join(home, ".go-blog.yaml")
//		_, err = os.Stat(cfgFile)
//
//		if os.IsNotExist(err) {
//			//log.Fatalf("The config file \".go-blog.yaml\" not found in [%s]\n", )
//		}
//		// 2、读取配置文件
//
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".go-blog")
//		viper.SetConfigType("yaml")
//	}
//	viper.AutomaticEnv()
//
//
//	if err := viper.ReadInConfig(); err != nil {
//		_, _ = fmt.Fprintf(os.Stderr, err.Error())
//	}
//}

