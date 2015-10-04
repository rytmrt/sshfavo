package main

import (
	"testing"
)

func TestNewServerFromCmdOpts(t *testing.T) {
	args := []string{
		"test",
		"-H", "host-name.test",
		"-p", "23",
		"-l", "login_user",
		"-i", "inasdf",
	}

	testResult := Server{
		Name:         "test",
		Host:         "host-name.test",
		Port:         23,
		LoginUser:    "login_user",
		UsePassword:  false,
		IdentityFile: "inasdf",
	}

	r, _ := NewServerFromCmdOpts(args)

	if !testResult.equals(r) {
		t.Errorf("res :%#v", *r)
		t.Errorf("res :%#v", testResult)
	}
}

func TestNewServerFromToml(t *testing.T) {
	tomlFile := "/Users/rytmrt/go14/src/github.com/rytmrt/ssh-favorites/test/test_server_info.toml"
	r, _ := NewServerFromToml(tomlFile)

	testResult := Server{
		Name:         "test-server",
		Host:         "hostname",
		LoginUser:    "loginUser",
		IdentityFile: "identityFile",
		UsePassword:  false,
		Port:         22,
	}

	if !testResult.equals(r) {
		t.Errorf("res :%#v", r)
	}
}

func TestServerToml(t *testing.T) {
}

func TestServerCreateSshConfig(t *testing.T) {
}
