package iniconfig

import (
	"io/ioutil"
	"testing"
)

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"password"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

func TestIniConfig(t *testing.T) {
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file failed")
	}

	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal failed, err:%v", err)
		return
	}
	t.Logf("unmarshal success, config: %#v", conf)

	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("marshal failed, err:%v", err)
	}
	t.Logf("marshal success, config:%s", string(confData))
}

func TestIniConfigFile(t *testing.T) {
	filename := "C:/logs/test.conf"
	var conf Config
	conf.ServerConf.Ip = "localhost"
	conf.ServerConf.Port = 88888
	err := MarshalToFile(filename, conf)
	if err != nil {
		t.Errorf("marshal failed, err:%v", err)
		return
	}

	var conf2 Config
	err = UnMarshalFile(filename, &conf)
	if err != nil {
		t.Errorf("unmarshal failed, err:%v", err)
	}
	t.Logf("unmarshal success, conf: %#v", conf2)
}
