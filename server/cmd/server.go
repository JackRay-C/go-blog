package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"syscall"
)

var (
	deamon        bool
	DefaultLock   = "blog.lock"
	serverCommand = &cobra.Command{
		Use: "server [start | stop | restart] [flags]",
	}
	startCommand = &cobra.Command{
		Use:   "start",
		Short: "服务启动",
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
	stopCommand = &cobra.Command{
		Use:   "stop",
		Short: "服务关停",
		Run: func(cmd *cobra.Command, args []string) {
			Stop()
		},
	}
	restartCommand = &cobra.Command{
		Use: "restart",
		Short: "restart",
		Run: func(cmd *cobra.Command, args []string) {
			Restart()
		},
	}
)

func init() {
	serverCommand.Flags().BoolVarP(&deamon, "deamon", "d", false, "If true, server start deamon.")
	serverCommand.AddCommand(startCommand)
	serverCommand.AddCommand(stopCommand)
	serverCommand.AddCommand(restartCommand)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Start() {
	binDir := filepath.Dir(os.Args[0])

	if _, err := os.Stat(path.Join(binDir, DefaultLock)); os.IsNotExist(err) {
		// 判断是否有lock文件，没有的话启动服务
		command := exec.Command(path.Join(binDir, ServerName), "-c", cfgFile)

		if err := command.Start(); err != nil {
			fmt.Println(err)
			return
		}

		if err = ioutil.WriteFile(path.Join(binDir, DefaultLock), []byte(strconv.Itoa(command.Process.Pid)), os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("start server: %v\n", command.Process.Pid)
	} else {
		file, err := ioutil.ReadFile(path.Join(binDir, DefaultLock))
		if err != nil {
			fmt.Println("read lock file failed.")
			return
		}

		pid, err := strconv.Atoi(string(file))
		if err != nil {
			fmt.Println(err)
			return
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("this server has already start, process: %d", process.Pid)
	}

}

func Stop() {
	binDir := filepath.Dir(os.Args[0])
	lockFile := path.Join(binDir, DefaultLock)
	file, err := ioutil.ReadFile(lockFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	pid, err := strconv.Atoi(string(file))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = syscall.Kill(pid, syscall.SIGTERM)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.Remove(lockFile); err != nil {
		fmt.Println(err)
		return
	}
}

func Restart()  {
	binDir := filepath.Dir(os.Args[0])
	lockFile := path.Join(binDir, DefaultLock)
	file, err := ioutil.ReadFile(lockFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	pid, err := strconv.Atoi(string(file))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = syscall.Kill(pid, syscall.SIGUSR2)
	if err != nil {
		fmt.Println(err)
		return
	}
}