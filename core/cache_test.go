package core

import (
	"testing"
)

func TestInitCache(t *testing.T) {
	ini := "gdhdhdhdhd"
	err := InitCache(ini)
	if err != nil {
		t.Errorf("TestInitCache:%v", err)
		t.FailNow()
	}
}

func TestInitAgain(t *testing.T) {
	ini := "gdhdhdhdhd"
	err := InitCache(ini)
	if err == nil {
		t.Errorf("TestInitCache:%v", err)
		t.FailNow()
	}
}

func TestGetEmpty(t *testing.T) {
	//InitCache("")
	val, err := Get(DEFALUT_DBNAME, "testKey")
	if len(val) > 0 {
		t.Errorf("get somthing?:%v", val)
	}
	if err == nil {
		t.Errorf("TestGet:no error", err)
	}
}

func TestWrongDbGet (t *testing.T) {
	//InitCache("")
	val, err := Get("no such db", "testKey")
	if len(val) > 0 {
		t.Errorf("get somthing?:%v", val)
	}
	if err == nil {
		t.Errorf("TestGet:no error", err)
	}
}

func TestSetAndGet(t *testing.T) {
	//InitCache("")
	err := Set(DEFALUT_DBNAME, "testKey", []byte("1dfdfdsd"))
	if err != nil {
		t.Errorf("TestSet:error", err)
	}
	val, err := Get(DEFALUT_DBNAME, "testKey")
	if len(val) <= 0 {
		t.Errorf("get nothing?%v", val)
	}
	if err != nil {
		t.Errorf("TestSetAndGetAndDel:%v", err)
	}
	if string(val) != "1dfdfdsd" {
		t.Errorf("get != set %v", val)
	}
}

func TestSetWrongDb(t *testing.T) {
	//InitCache("")
	err := Set("no such db", "testKey", []byte("1dfdfdsd"))
	if err == nil {
		t.Errorf("TestSet:error", err)
	}
}

func TestSetAgain(t *testing.T) {
	err := Set(DEFALUT_DBNAME, "testKey", []byte("1dfdfdsd111"))
	if err != nil {
		t.Errorf("TestSet:error", err)
	}
	val, err := Get(DEFALUT_DBNAME, "testKey")
	if len(val) <= 0 {
		t.Errorf("get nothing?%v", val)
	}
	if err != nil {
		t.Errorf("TestSetAndGetAndDel:%v", err)
	}
	if string(val) != "1dfdfdsd111" {
		t.Errorf("get != set %v", val)
	}
}

func TestDel(t *testing.T) {
	err := Del(DEFALUT_DBNAME, "testKey")
	if err != nil {
		t.Errorf("TestDel:error", err)
	}
	val, err := Get(DEFALUT_DBNAME, "testKey")
	if len(val) > 0 {
		t.Errorf("get somthing?:%v", val)
	}
	if err == nil {
		t.Errorf("TestGet:no error", err)
	}
}