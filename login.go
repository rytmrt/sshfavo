package main

import ()

type Login struct {
}

func NewLogin() *Login {
	var i Login
	return &i
}

func (self *Login) Run() (err error) {
	return
}
