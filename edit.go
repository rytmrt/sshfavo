package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Edit struct {
	Name string
}

func NewEdit(name string) *Edit {
	i := Edit{
		Name: name,
	}
	return &i
}

func (self *Edit) Run() (err error) {
	locationPathAbs, _ := filepath.Abs(ConvPath(LocationPath))
	var (
		buf bytes.Buffer
	)

	buf.WriteString(locationPathAbs)
	buf.WriteString(ServerListPath)

	_, e := os.Stat(buf.String())
	if e != nil {
		os.MkdirAll(buf.String(), 0755)
	}

	buf.WriteString("/")
	buf.WriteString(self.Name)
	buf.WriteString(".toml")

	_, e = os.Stat(buf.String())
	if e != nil {
		server, _ := NewServerFromCmdOpts([]string{self.Name})
		tomlStr, _ := server.Toml()
		e2 := ioutil.WriteFile(buf.String(), []byte(tomlStr), os.ModePerm)
		fmt.Printf("%#v/n", e2)
	}

	cmd := exec.Command("vi", buf.String())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	return
}
