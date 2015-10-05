package main

import (
	"path/filepath"
	"testing"
)

func TestGetServerList(t *testing.T) {
	serverListPath, _ := filepath.Abs("./test")
	testResult := []string{
		"test_server_info.toml",
	}
	list := NewList(serverListPath)
	res, _ := list.GetServerList()
	if !func() bool {
		var r bool = true
		for i := range testResult {
			r = r && (len(res) > i) && (res[i] == testResult[i])
		}
		return r
	}() {
		t.Errorf("ng")
	}
}

func TestGetServer(t *testing.T) {

	testResult := Server{
		Name:         "test-server",
		Host:         "hostname",
		LoginUser:    "loginUser",
		IdentityFile: "identityFile",
		UsePassword:  false,
		Port:         22,
	}

	serverListPath, _ := filepath.Abs("./test")
	testServer := "test_server_info"
	list := NewList(serverListPath)
	r, _ := list.GetServer(testServer)

	if !testResult.equals(r) {
		t.Errorf("res :%#v", r)
	}
}

func TestExistServer(t *testing.T) {

	serverListPath, _ := filepath.Abs("./test")
	testServer := "test_server_info"
	list := NewList(serverListPath)
	exist := list.ExistServer(testServer)
	if !exist {
		t.Errorf("'%s/%s.toml' is not exist!", serverListPath, testServer)
	}
	exist = list.ExistServer("dummy")
	if exist {
		t.Errorf("'dummy' is exist!")
	}
}
