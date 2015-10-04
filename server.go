package main

import (
	"bytes"
	"flag"
	"github.com/BurntSushi/toml"
)

type Server struct {
	Name         string `toml:"name"`
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	UsePassword  bool   `toml:"usePassword"`
	LoginUser    string `toml:"loginUser"`
	IdentityFile string `toml:"identityFile"`
}

func NewServerFromCmdOpts(args []string) (svrInfo *Server, err error) {
	var (
		name         string
		host         string
		loginUser    string
		identityFile string
		usePassword  bool
		port         int
	)

	// Define option flag parse
	flags := flag.NewFlagSet("srv_option", flag.ContinueOnError)
	flags.Usage = func() {}
	flags.StringVar(&host, "H", "localhost", "Host name")
	flags.StringVar(&loginUser, "l", "root", "Login name")
	flags.StringVar(&identityFile, "i", "~/.ssh/id_rsa", "Login name")
	flags.BoolVar(&usePassword, "P", false, "Password authentication")
	flags.IntVar(&port, "p", 22, "Port")

	// Parse commandline flag
	if err := flags.Parse(args[0:]); err != nil {
		return nil, err
	}

	var arguments []string
	for 0 < flags.NArg() {
		arguments = append(arguments, flags.Arg(0))
		flags.Parse(flags.Args()[1:])
	}
	name = arguments[0]

	svrInfo = &Server{
		Name:         name,
		Host:         host,
		LoginUser:    loginUser,
		IdentityFile: identityFile,
		UsePassword:  usePassword,
		Port:         port,
	}

	return
}

func NewServerFromToml(tomlFile string) (svrInfo *Server, err error) {
	var s Server
	svrInfo = &s
	_, err = toml.DecodeFile(tomlFile, &s)
	return
}

func (self *Server) Toml() (string, error) {
	var b bytes.Buffer
	e := toml.NewEncoder(&b)
	err := e.Encode(*self)
	return b.String(), err
}

func (self *Server) equals(obj *Server) bool {
	return (self.Name == obj.Name) &&
		(self.Host == obj.Host) &&
		(self.LoginUser == obj.LoginUser) &&
		(self.IdentityFile == obj.IdentityFile) &&
		(self.UsePassword == obj.UsePassword) &&
		(self.Port == obj.Port)
}
