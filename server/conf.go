package server

import (
	"flag"
	"nats_example/baselib/mysql_client"
	natsclient "nats_example/baselib/nats_client"
)

var (
	confPath string
	Conf     *eventstoreConfig
)

type eventstoreConfig struct {
	MySQL []mysql_client.MySQLConfig
	NATS  natsclient.NatsConf
}

func (c *eventstoreConfig) String() string {
	return ""
}

// prepairConfig func
func (c *eventstoreConfig) prepairConfig() {

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
