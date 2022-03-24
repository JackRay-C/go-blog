package cmd

import (
	"github.com/spf13/cobra"
)


var (
	logCommand   = &cobra.Command{
		Use:   "logging [option] [flags]",
		Short: "log configuration",
		Long:  "configuration server log directory、filename、flush time duration...",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

type loggingT struct {
	level string
	format string
	directory string
	filename string
	showLine bool
	encodeLevel string
	stacktraceKey string
	logInConsole bool
	logMaxSize uint64
	logMaxAge int
	logMaxBackups int
}

var logging loggingT

func init() {
	rootCmd.AddCommand(logCommand)

	logging.level  = "info"
	logging.format = "console"
	logging.directory = "logs"
	logging.filename = "latest_log"
	logging.showLine = true
	logging.encodeLevel = "LowercaseColorLevelEncoder"
	logging.stacktraceKey = "stacktrace"
	logging.logInConsole = true
	logging.logMaxSize = 18000
	logging.logMaxAge = 30
	logging.logMaxBackups = 5


	logCommand.Flags().StringVar(&logging.level, "log-level", "info", "log level [error|info|fatal|debug]")
	logCommand.Flags().StringVar(&logging.directory, "log-dir", "logs", "If non-empty, write log files in this directory")
	logCommand.Flags().StringVar(&logging.filename, "log-filename", "", "If non-empty, use this log file ")
	logCommand.Flags().BoolVar(&logging.showLine, "log-show-line", true, "If non-empty, log show line")
	logCommand.Flags().BoolVar(&logging.logInConsole, "log-in-console", true, "If true, log write in console and file")
	logCommand.Flags().Uint64Var(&logging.logMaxSize, "log-max-size", 18000, "log max size")
	logCommand.Flags().IntVar(&logging.logMaxAge, "log-max-age", 30, "log max age")
	logCommand.Flags().IntVar(&logging.logMaxBackups, "log-max-backups", 5, "log max backups ")

}
