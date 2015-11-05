package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Login struct {
	Name string
}

func NewLogin(name string) *Login {
	i := Login{
		Name: name,
	}
	return &i
}

func (self *Login) Run() (err error) {
	locationPathAbs, _ := filepath.Abs(ConvPath(LocationPath))
	var buf bytes.Buffer
	buf.WriteString(locationPathAbs)
	buf.WriteString(ServerListPath)
	list := NewList(buf.String())
	buf.Reset()

	if list.ExistServer(self.Name) {
		server, _ := list.GetServer(self.Name)

		var cmd *exec.Cmd

		if server.UsePassword {
			cmd = exec.Command("ssh", server.Host, "-l", server.LoginUser, "-p", strconv.Itoa(server.Port))
		} else {
			cmd = exec.Command("ssh", server.Host, "-l", server.LoginUser, "-p", strconv.Itoa(server.Port), "-i", server.IdentityFile)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err = cmd.Run()
	} else {
		help := NewHelp()
		help.Run()
	}

	return
}
