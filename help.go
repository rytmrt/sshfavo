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

	fmt.Printf("Usage: sshfavo [options]\n")
	fmt.Printf("\n")
	fmt.Printf("OPTIONS:\n")
	fmt.Printf("-v, --version\n")
	fmt.Printf("\n")
	fmt.Printf("-h, --help\n")
	fmt.Printf("\n")
	fmt.Printf("-l, --login SERVER_NAME\n")
	fmt.Printf("  ssh login.\n")
	fmt.Printf("\n")
	fmt.Printf("-e, --edit SERVER_NAME\n")
	fmt.Printf("  Edit server config using vi.\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("Favorite server list:\n")
	if len(serverList) <= 0 {
		fmt.Printf("  Favorite sever is empty!\n")
	} else {
		for i := range serverList {
			fmt.Printf("  - %s\n", serverList[i])
		}
	}
	return
}
