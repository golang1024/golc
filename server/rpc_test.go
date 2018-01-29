package server

import (
	"testing"

	"golc/core"
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

func TestGoLC_MGet(t *testing.T) {
	lc := new(GoLC)

	reply := []*core.MGetResponse{}
	lc.Set([]string{"debug", "testkey", "testval"}, nil)
	lc.Set([]string{"debug", "testkey2", "testval2"}, nil)
	err := lc.MGet([]string{"debug", "testkey", "testkey2", "testkey3"}, &reply)
	if err != nil || len(reply) != 3{
		t.Errorf("batch get err: %v, %v", err, reply)
	}
	if string(reply[0].Result) != "testval" || string(reply[1].Result) != "testval2" || reply[2].Err == nil {
		t.Errorf("batch get val is wrong: %v", reply)
	}

}