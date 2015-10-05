package main

import (
	"bytes"
	"path/filepath"
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
	serverListPath, _ := filepath.Abs("./test")
	var buf bytes.Buffer
	buf.WriteString(serverListPath)
	buf.WriteString("/test_server_info.toml")
	tomlFile := buf.String()

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
