package client

import (
	"testing"
	"golc/server"
)

var (
	golc *GoLCClient
)


func TestGoLCClientRun(t *testing.T) {
	ini := "client.ini"
	err := server.Run(ini)
	if err != nil {
		t.Errorf("run err:%v", err)
		t.FailNow()
	}
	golc = NewsClient("client.ini")
}

//func TestGoLCClient_GetEmpty(t *testing.T) {
//	t.Log("TestGoLCClient_GetEmpty")
//	reply, err := golc.Get("debug", "no such key")
//	if err == nil {
//		t.Errorf("get no err from empty key")
//	}
//	if len(reply) > 0 {
//		t.Errorf("get val from empty key: %v", reply)
//	}
//}

func TestGoLCClient_SetAndGet(t *testing.T) {
	err := golc.Set("debug", "test key aaa", []byte("test val bbb"))
	if err != nil {
		t.Errorf("set err:%v", err)
		t.FailNow()
	}
	reply, err := golc.Get("debug", "test key aaa")
	if err == nil {
		t.Errorf("get no err from empty key")
	}
	if len(reply) <= 0 || string(reply) != "test val bbb"{
		t.Errorf("get wrong val from key: %s", reply)
	}
}

func TestClose(t *testing.T) {
	err := server.Close()
	if err != nil {
		t.Errorf("run err:%v", err)
		t.FailNow()
	}
}