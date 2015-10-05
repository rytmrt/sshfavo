package main

import (
	"fmt"
)

type Version struct {
	Name string
}

func NewVerson() *Version {
	i := Version{
		Name: "0.1.0",
	}
	return &i
}

func (self *Version) Run() (err error) {
	fmt.Printf("sshfavo %s\n", self.Name)
	return
}
