package conf

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var Conf *Config

type Config struct {
	ListenAddr string       `toml:"listen_addr"`
	MysqlConf  *MySQLConfig `toml:"mysql_conf"`
}

type MySQLConfig struct {
	Dsn     string
	MaxConn int
	MaxIdle int
}

func LoadConfig(path string) error {
	content, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	if _, err := toml.Decode(string(content), &Conf); err != nil {
		return err
	}
	return nil
}
