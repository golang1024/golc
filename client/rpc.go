package client

import (
	env "github.com/Unknwon/goconfig"

	"golc/core"

	"net/rpc"
	"fmt"
	"time"

	"errors"
)


type GoLCClient struct {
	NetWork	string
	Server	string
}

func NewsClient(ini string) (*GoLCClient){
	c, err := env.LoadConfigFile(ini)
	var network, addr string
	if err == nil {
		network , _ = c.GetValue(core.SERVER_CONF_SECTION, "network")
		addr, _ = c.GetValue(core.SERVER_CONF_SECTION, "address")
	}

	if len(network) <= 0 || (network != "tcp" && network != "unix") {
		network = core.SERVER_NETWORK
	}

	if len(addr) <= 0 {
		network = core.SERVER_ADDRESS
	}
	return &GoLCClient{NetWork:network, Server:addr}
}



func (c *GoLCClient)Get(dbName, key string) ([]byte, error) {
	reply := []byte{}
	err := Call(c.NetWork, c.Server, core.RPC_ACTION_GET, []string{dbName, key}, &reply, core.CLIENT_TIME_OUT)
	return reply, err
}

func (c *GoLCClient)Set(dbName, key string, val []byte) (error) {
	reply := []byte{}
	return Call(c.NetWork, c.Server, core.RPC_ACTION_SET, []string{dbName, key, string(val)}, &reply, core.CLIENT_TIME_OUT)
}

func (c *GoLCClient)Del(dbName, key string) (error) {
	reply := []byte{}
	return Call(c.NetWork, c.Server, core.RPC_ACTION_DEL, []string{dbName, key}, &reply, core.CLIENT_TIME_OUT)
}


func (c *GoLCClient)MGet(dbName string, keys []string) ([]*core.MGetResponse, error) {
	if len(keys) == 0 {
		return nil, errors.New("key is empty")
	}
	reply := make([]*core.MGetResponse, len(keys))
	args := []string{dbName}
	args = append(args, keys...)
	err := Call(c.NetWork, c.Server, core.RPC_ACTION_MGET, args, &reply, core.CLIENT_TIME_OUT)
	return reply, err
}


func Call(network, srv, rpcname string, args []string, reply interface{}, timeOutMS int) error {

	c, errx := rpc.DialHTTP(network, srv)
	if errx != nil {
		return fmt.Errorf("ConnectError: %s", errx.Error())
	}
	ch := make(chan error, 2)
	defer c.Close()
	if timeOutMS <= 0 {
		return c.Call(rpcname, args, reply)
	} else {
		go func() {
			defer close(ch)
			ch <- c.Call(rpcname, args, reply)
			//fmt.Println("rep:",reply)
		}()
		select{
		case <- time.After(time.Millisecond * time.Duration(timeOutMS)):
			return fmt.Errorf("conn time out: %s:%s:%s %v", network, srv, rpcname, args)
		case err := <- ch:
			return err
		}
	}

}


