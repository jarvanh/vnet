package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jarvanh/vnet/common/log"
	"github.com/jarvanh/vnet/utils"
	"github.com/jarvanh/vnet/utils/iox"
)

var (
	cfg        = flag.String("cfg", "", "Config file for Manager.")
	config     *Config
	configFile string
)

// Config is global config
type Config struct {
	Mode               string             `json:"mode"`
	DbConfig           DbConfig           `json:"dbconfig"`
	ShadowsocksOptions ShadowsocksOptions `json:"shadowsocks_options"`
	DNSOptions         DnsOptions         `json:"dns_options"`
}

type DnsOptions struct {
	DNS1       string `json:"dns1"`
	DNS2       string `json:"dns2"`
	IPV4Prefer bool   `json:"ipv4_prefer"`
}

// DbConfig is global database config
type DbConfig struct {
	Host           string  `json:"host"`
	User           string  `json:"user"`
	Passwd         string  `json:"passwd"`
	Port           string  `json:"port"`
	Database       string  `json:"database"`
	Rate           float32 `json:"rate"`
	NodeId         int     `json:"node_id`
	SyncTime       int     `json:"sync_time"`
	OnlineSyncTime int     `json:"online_sync_time"`
	Level          int     `json:"level"`
}

// ShadowsocksOptions is global shadowoscks service config
// the TCPSwitch is set tcp service enable
// the UDPSwitch is set udp service enable
// ConnectTimeout is set Shadowsocks select timeout time. not real connect time
type ShadowsocksOptions struct {
	ConnectTimeout int    `json:"connect_timeout"`
	TCPSwitch      string `json:"tcp_switch"`
	UDPSwitch      string `json:"udp_switch"`
}

func DefaultConfig() *Config {
	return &Config{
		Mode: "db",
		DbConfig: DbConfig{
			Level:          -1,
			Rate:           -1,
			SyncTime:       60000,
			OnlineSyncTime: 60000,
		},
		ShadowsocksOptions: ShadowsocksOptions{
			ConnectTimeout: 3000,
			TCPSwitch:      "true",
			UDPSwitch:      "true",
		},
		DNSOptions: DnsOptions{
			DNS1:       "8.8.8.8:53",
			DNS2:       "8.8.4.4:53",
			IPV4Prefer: true,
		},
	}
}

// CurrentConfig whil return current or default conifg
func CurrentConfig() *Config {
	if config == nil {
		conf, err := LoadDefault()
		if err != nil {
			panic(err)
		}
		config = conf
	}
	return config
}

func LoadDefault() (*Config, error) {
	flag.Parse()
	if *cfg != "" {
		return LoadConfig(*cfg)
	} else {
		return LoadConfig("config.json")
	}
}

func LoadConfig(file string) (*Config, error) {
	utils.RLock(file)
	defer utils.RUnLock(file)
	if !iox.IsFileExist(file) {
		absFile, err := filepath.Abs(file)
		if err != nil {
			log.Err(err)
		} else {
			log.Warn("%s is not exist", absFile)
		}
		configFile = file
		config = DefaultConfig()
		data, _ := json.MarshalIndent(config, "", "    ")
		ioutil.WriteFile(configFile, data, 0644)
		return config, nil
	}
	config = &Config{
		Mode: "bare",
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %v", err)
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("resolve config file failed: %v", err)
	}
	configFile = file
	return config, nil
}

func SaveConfig() error {
	if config == nil {
		return fmt.Errorf("not config loaded!")
	}

	data, err := json.MarshalIndent(config, "", "    ")

	if err != nil {
		return fmt.Errorf("config marshal failed!")
	}

	return ioutil.WriteFile(configFile, data, 0644)
}

func (self Config) String() string {
	data, err := json.MarshalIndent(self, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(data)
}
