package main

import (
	"flag"
	"os/user"
	"strings"
)

const (
	LocationPath   = "~/.sshfavo"
	ServerListPath = "/servers"
)

type SubCmd interface {
	Run() (err error)
}

func Run(args []string) int {

	var (
		help  bool
		edit  string
		login string
	)

	// Define option flag parse
	flags := flag.NewFlagSet("option", flag.ContinueOnError)
	flags.Usage = func() {
		NewHelp().Run()
	}
	flags.BoolVar(&help, "h", false, "show help")
	flags.StringVar(&edit, "e", "false", "edit")
	flags.StringVar(&login, "l", "false", "login")

	// Parse commandline flag
	if err := flags.Parse(args[0:]); err != nil {
		return -1
	}

	var (
		cmd SubCmd
	)

	switch {
	case edit != "false" && login != "false":
		cmd = NewHelp()
	case edit != "false":
		cmd = NewEdit(edit)
	case login != "false":
		cmd = NewLogin()
	case help:
		cmd = NewHelp()
	default:
		cmd = NewHelp()
	}
	_ = cmd.Run()

	return 0
}

func ConvPath(srcPath string) string {
	usr, _ := user.Current()
	r := strings.Replace(srcPath, "~", usr.HomeDir, 1)
	return r
}
