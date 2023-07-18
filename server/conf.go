package server

import (
	"flag"
)

var (
	confPath string
	Conf     *cacheConfig
)

type cacheConfig struct {
}

func (c *cacheConfig) String() string {
	return ""
}

// prepairConfig func
func (c *cacheConfig) prepairConfig() {

}

// InitializeConfig func
func InitializeConfig() (err error) {
	// _, err = toml.DecodeFile(confPath, &Conf)

	// if err != nil {
	// 	err = fmt.Errorf("decode file %s error: %+v", confPath, err)
	// 	panic(err)
	// }
	// Conf.prepairConfig()
	return
}
func init() {
	flag.StringVar(&confPath, "conf", "./event_store.toml", "config path")
}
