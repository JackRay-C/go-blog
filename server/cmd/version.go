package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type version struct {
	Version      string
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildTime    string
	GoVersion    string
	Compiler     string
	Platform     string
}

func (v *version) String() string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}

var (
	GoVersion    = ""
	BuildTime    = ""
	GitCommit    = ""
	BuildVersion = ""
	Compiler     = ""
	GitVersion   = ""
	Platform     = ""
	ServerName   = ""
	ClientName = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出go blog ui和server版本信息",
	Long:  "All software has version. This is go-blog's ",
	Run: func(cmd *cobra.Command, args []string) {
		clientVersion := &version{
			Version:    BuildVersion,
			GitCommit:  GitCommit,
			GitVersion: GitVersion,
			BuildTime:  BuildTime,
			GoVersion:  GoVersion,
			Compiler:   Compiler,
			Platform:   Platform,
		}
		serverVersion := &version{
			Version:    BuildVersion,
			GitCommit:  GitCommit,
			GitVersion: GitVersion,
			BuildTime:  BuildTime,
			GoVersion:  GoVersion,
			Compiler:   Compiler,
			Platform:   Platform,
		}

		fmt.Println("Client Version: " + clientVersion.String())
		fmt.Println("Server Version: " + serverVersion.String())
	},
}

func init() {

}
