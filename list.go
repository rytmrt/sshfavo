package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

type List struct {
	ListPath string
}

func NewList(listPath string) *List {
	i := List{
		ListPath: listPath,
	}
	return &i
}

func (self *List) GetServerList() (servers []string, err error) {
	files, _ := ioutil.ReadDir(self.ListPath)
	for _, f := range files {
		servers = append(servers, f.Name())
	}
	return
}

func (self *List) GetServer(name string) (server *Server, err error) {
	var buf bytes.Buffer
	buf.WriteString(self.ListPath)
	buf.WriteString("/")
	buf.WriteString(name)
	buf.WriteString(".toml")
	server, _ = NewServerFromToml(buf.String())
	return
}

func (self *List) ExistServer(name string) bool {
	var buf bytes.Buffer
	buf.WriteString(self.ListPath)
	buf.WriteString("/")
	buf.WriteString(name)
	buf.WriteString(".toml")
	serverFilePath := buf.String()
	_, e := os.Stat(serverFilePath)
	return e == nil
}
