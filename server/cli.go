package main


import (
	"blog/cmd"
	"os"
)

// usage:
// 	Commands:
// 		help  	usage
// 		init  	create app folder and build vue dist and build go main
// 		start 	[ui|server] start front and backend server
// 		stop 	stop sever
//		reload 	reload config
// 		restart restart server
//		version display version information
//
//	Options:
//		--config  	config.yaml file path
//		--logs-dir 	logs file
//		--daemon 	start daemon


func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

