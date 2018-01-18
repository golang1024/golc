package server

import (
	env "github.com/Unknwon/goconfig"

	"golc/core"


	"errors"
	"net"

	"net/rpc"
	"net/http"
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
	c, _ := env.LoadConfigFile(ini)
	var network, addr string
	network , _ = c.GetValue(core.SERVER_CONF_SECTION, "network")
	addr, _ = c.GetValue(core.SERVER_CONF_SECTION, "address")

	if len(network) <= 0 || (network != "tcp" && network != "unix") {
		network = core.SERVER_NETWORK
	}

	if len(addr) <= 0 {
		network = core.SERVER_ADDRESS
	}

	l, e := net.Listen(network, addr)
	if e != nil {
		return e
	}
	lc := new(GoLC)
	rpc.Register(lc)
	rpc.HandleHTTP()
	http.Serve(l, nil)
	return nil
}