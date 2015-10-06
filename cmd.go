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

type FlagBool struct {
	short bool
	long  bool
}

func (s *FlagBool) Get() bool {
	return s.short || s.long
}

type FlagString struct {
	short string
	long  string
}

func (s *FlagString) Get() string {
	if s.long != "" {
		return s.long
	}
	return s.short
}

func Run(args []string) int {

	var (
		version FlagBool
		help    FlagBool
		edit    FlagString
		login   FlagString
	)

	// Define option flag parse
	flags := flag.NewFlagSet("option", flag.ContinueOnError)
	flags.Usage = func() {
		NewHelp().Run()
	}
	flags.BoolVar(&version.short, "v", false, "show version")
	flags.BoolVar(&version.long, "version", false, "show version")
	flags.BoolVar(&help.short, "h", false, "show help")
	flags.BoolVar(&help.long, "help", false, "show help")
	flags.StringVar(&edit.short, "e", "", "edit")
	flags.StringVar(&edit.long, "edit", "", "edit")
	flags.StringVar(&login.short, "l", "", "login")
	flags.StringVar(&login.long, "login", "", "login")

	// Parse commandline flag
	if err := flags.Parse(args[0:]); err != nil {
		return -1
	}

	var (
		cmd SubCmd
	)

	switch {
	case version.Get():
		cmd = NewVerson()

	case edit.Get() != "" && login.Get() != "":
		cmd = NewHelp()

	case edit.Get() != "":
		cmd = NewEdit(edit.Get())

	case login.Get() != "":
		cmd = NewLogin(login.Get())

	case help.Get():
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
