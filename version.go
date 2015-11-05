package main

import (
	"fmt"
)

type Version struct {
	Name string
}

func NewVersion() *Version {
	i := Version{
		Name: "0.1.1",
	}
	return &i
}

func (self *Version) Run() (err error) {
	fmt.Printf("sshfavo %s\n", self.Name)
	return
}
