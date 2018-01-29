package server

import (
	env "github.com/Unknwon/goconfig"

	"golc/core"


	"errors"
	"net"

	"net/rpc"
	"net/http"
	"sync"
	"fmt"
	"time"
)

var (
	lock = &sync.Mutex{}
	socketListener net.Listener
)
type GoLC struct {

}


func (lc *GoLC) Get(args []string, reply *[]byte) (err error) {
	if len(args) < 2 {
		return errors.New("error args lens")
	}

	*reply, err = core.Get(args[0], args[1])
	return err
}

func (lc *GoLC) MGet(args []string, reply *[]*core.MGetResponse) (err error) {
	if len(args) < 2 {
		return errors.New("error args lens")
	}

	for _, key := range args[1:] {
		tmpReplay, err := core.Get(args[0], key)
		if err == nil {

		}
		res := &core.MGetResponse{Key:key, Result: tmpReplay}

		if err != nil {
			res.ErrMsg = err.Error()
		}
		*reply = append(*reply, res)
	}
	return nil
}

func (lc *GoLC) Set(args []string, reply *[]byte) error {
	if len(args) < 3 {
		return errors.New("error args lens")
	}

	err := core.Set(args[0], args[1], []byte(args[2]))
	return err
}


func (lc *GoLC) Del(args []string, reply *[]byte) error {
	if len(args) <= 2 {
		return errors.New("error args lens")
	}

	err := core.Del(args[0], args[1])
	return err
}

func Run(ini string) error {
	c, err := env.LoadConfigFile(ini)
	var network, addr string
	if err == nil {
		network , _ = c.GetValue(core.SERVER_CONF_SECTION, "network")
		addr, _ = c.GetValue(core.SERVER_CONF_SECTION, "address")
		cacheIniFile, _ := c.GetValue(core.SERVER_CONF_SECTION, "cache_config")
		if len(cacheIniFile) > 0 {
			core.InitCache(cacheIniFile)
		}
	}

	if len(network) <= 0 || (network != "tcp" && network != "unix") {
		network = core.SERVER_NETWORK
	}

	if len(addr) <= 0 {
		network = core.SERVER_ADDRESS
	}
	lock.Lock()
	defer lock.Unlock()
	l, e := net.Listen(network, addr)
	if e != nil {
		return e
	}
	fmt.Printf("[golc]server: Listen on %s, %s\n", network, addr)
	lc := new(GoLC)
	rpc.Register(lc)
	rpc.HandleHTTP()

	go http.Serve(l, nil)
	fmt.Printf("[golc]server: run on %s, %s\n", network, addr)
	time.Sleep(time.Millisecond * 10)
	socketListener = l
	return nil
}

func Close() error {
	lock.Lock()
	defer lock.Unlock()
	fmt.Sprintf("[golc]server: try to close")
	return socketListener.Close()
}