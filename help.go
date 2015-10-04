package main

import (
	"bytes"
	"fmt"
	"path/filepath"
)

type Help struct {
}

func NewHelp() *Help {
	i := Help{}
	return &i
}

func (self *Help) Run() (err error) {
	locationPathAbs, _ := filepath.Abs(ConvPath(LocationPath))
	var buf bytes.Buffer
	buf.WriteString(locationPathAbs)
	buf.WriteString(ServerListPath)
	list := NewList(buf.String())
	serverList, _ := list.GetServerList()

	fmt.Printf("usage: sshfavo [-h] [-e server_name] [-l server_name]\n")
	fmt.Printf("\n")
	fmt.Printf("+-------------------------------------------------------------------------+\n")
	fmt.Printf("| Server list                                                             |\n")
	fmt.Printf("+-------------------------------------------------------------------------+\n")
	if len(serverList) <= 0 {
		fmt.Printf("You have not Server.toml!\n")
	} else {
		for i := range serverList {
			fmt.Printf(" + %s\n", serverList[i])
		}
	}
	return
}
