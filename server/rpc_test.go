package server

import (
	"testing"

)

func TestRun(t *testing.T) {
	ini := "server.ini"
	err := Run(ini)
	if err != nil {
		t.Errorf("run err:%v", err)
	}
	err = Run(ini)
	if err == nil {
		t.Errorf("run again no err?")
	}
	err = Close()
	if err != nil {
		t.Errorf("run err:%v", err)
	}
}

func TestGoLC_Set(t *testing.T) {
	lc := new(GoLC)
	reply := []byte{}
	err := lc.Set([]string{"debug", "testkey", "testval"}, &reply)
	if err != nil {
		t.Errorf("set error:%v", err)
	}
	err = lc.Get([]string{"debug", "testkey"}, &reply)
	if err != nil {
		t.Errorf("set error:%v", err)
	}
	if string(reply) !=  "testval" {
		t.Errorf("get wrong:%s", reply)
	}
}