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

func Parse(args []string) (v bool, h bool, e string, l string, err error) {

	var (
		ver   [2]bool
		help  [2]bool
		edit  [2]string
		login [2]string
	)

	// Define option flag parse
	flags := flag.NewFlagSet("option", flag.ContinueOnError)
	flags.Usage = func() {
		NewHelp().Run()
	}
	flags.BoolVar(&ver[0], "v", false, "Show version")
	flags.BoolVar(&ver[1], "version", false, "Show version")
	flags.BoolVar(&help[0], "h", false, "Show help")
	flags.BoolVar(&help[1], "help", false, "Show help")
	flags.StringVar(&edit[0], "e", "", "Edit server list")
	flags.StringVar(&edit[1], "edit", "", "Edit server list")
	flags.StringVar(&login[0], "l", "", "Login server")
	flags.StringVar(&login[1], "login", "", "Login server")

	// Parse commandline flag
	err = flags.Parse(args[0:])
	if err != nil {
		return
	}

	v = ver[0] || ver[1]
	h = help[0] || help[1]
	if edit[1] != "" {
		e = edit[1]
	} else {
		e = edit[0]
	}
	if login[1] != "" {
		l = login[1]
	} else {
		l = login[0]
	}
	return
}

func Run(args []string) int {

	ver, help, edit, login, err := Parse(args)

	if err != nil {
		return -1
	}

	var cmd SubCmd

	switch {
	case ver:
		cmd = NewVersion()

	case help:
		cmd = NewHelp()

	case edit != "" && login != "":
		cmd = NewHelp()

	case edit != "":
		cmd = NewEdit(edit)

	case login != "":
		cmd = NewLogin(login)

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
