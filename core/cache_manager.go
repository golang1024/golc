package core

import (
	theCache "github.com/allegro/bigcache"
	env "github.com/Unknwon/goconfig"

	"sync"
	"errors"
	"time"
	"fmt"
)

var (
	lock = &sync.RWMutex{}
	cahchMap map[string]*theCache.BigCache
)

type MGetResponse struct {
	Key		string
	Result	[]byte
	ErrMsg	string
}

func init() {
	cahchMap = nil
}


func InitCache(ini string) error {
	if cahchMap != nil {
		return errors.New("has init")
	}
	c, err := env.LoadConfigFile(ini)

	if err != nil {
		c, err = env.LoadConfigFile("test.ini")
		if err != nil {
			return err
		}
	}

	dbs := c.MustValueArray("db", "dbs", ",")
	if len(dbs) == 0 {
		dbs = []string{DEFALUT_DBNAME}
		//return errors.New("no dbs in config")
	}
	lock.Lock()
	defer lock.Unlock()
	cahchMap = make(map[string]*theCache.BigCache, len(dbs))
	for _, db := range dbs {
		dc := theCache.DefaultConfig(DEFALUT_CACHE_CONF_LIFEWINDOW * time.Second)
		newCache, err := theCache.NewBigCache(dc)
		if err != nil {
			return err
		}
		cahchMap[db] = newCache
		fmt.Printf("[golc]Init db:%s\n", db)
	}
	return nil
}


 func Get(dbName, key string) ([]byte, error) {
 	lock.RLock()
 	if c, ok := cahchMap[dbName]; ok {
 		lock.RUnlock()
 		return c.Get(key)
	}
	lock.RUnlock()
	return nil, errors.New("No such db")
 }

func Set(dbName, key string, val []byte) (error) {
	lock.RLock()
	if c, ok := cahchMap[dbName]; ok {
		lock.RUnlock()
		return c.Set(key, val)
	}
	lock.RUnlock()
	return errors.New("No such db")
}

func Del(dbName, key string) (error) {
	lock.RLock()
	if c, ok := cahchMap[dbName]; ok {
		lock.RUnlock()
		return c.Delete(key)
	}
	lock.RUnlock()
	return errors.New("No such db")
}
