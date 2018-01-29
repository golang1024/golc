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

func TestGoLCClient_GetEmpty(t *testing.T) {
	t.Log("TestGoLCClient_GetEmpty")
	reply, err := golc.Get("debug", "no such key")
	if err == nil {
		t.Errorf("get no err from empty key")
	}
	if len(reply) > 0 {
		t.Errorf("get val from empty key: %v", reply)
	}
}

func TestGoLCClient_SetAndGet(t *testing.T) {
	err := golc.Set("debug", "test key aaa", []byte("test val bbb"))
	if err != nil {
		t.Errorf("set err:%v", err)
		t.FailNow()
	}
	reply, err := golc.Get("debug", "test key aaa")
	if err != nil {
		t.Errorf("get err from key %v", err)
	}
	if len(reply) <= 0 || string(reply) != "test val bbb"{
		t.Errorf("get wrong val from key: %s", reply)
	}

}


func TestGoLCClient_MGET(t *testing.T) {
	golc.Set("debug", "test key aaa2", []byte("test val aaa2"))

	reply, err := golc.MGet("debug", []string{"test key aaa", "test key aaa2", "test key ccc"})
	if err != nil {
		t.Errorf("batch get err:%v", err)
	}
	if len(reply) != 3 || string(reply[0].Result) != "test val bbb" || string(reply[1].Result) != "test val aaa2" ||
		len(reply[2].ErrMsg) <= 0 {
		t.Errorf("batch get val err: %v, reply:%v", err, reply)
	}
}

func TestClose(t *testing.T) {
	err := server.Close()
	if err != nil {
		t.Errorf("run err:%v", err)
		t.FailNow()
	}
}