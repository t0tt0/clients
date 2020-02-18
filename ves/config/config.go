package config

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/Myriad-Dreamin/go-ves/lib/core-cfg"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	ConnectionType    string        `json:"connection-type" yaml:"connection-type" toml:"connection-type" xml:"connection-type"`
	Host              string        `json:"host" yaml:"host" toml:"host" xml:"host"`
	Password          string        `json:"password" yaml:"password" toml:"password" xml:"password"`
	Database          int           `json:"database" yaml:"database" toml:"database" xml:"database"`
	MaxIdle           int           `json:"max-idle" yaml:"max-idle" toml:"max-idle" xml:"max-idle"`
	MaxActive         int           `json:"max-active" yaml:"max-active" toml:"max-active" xml:"max-active"`
	ConnectionTimeout time.Duration `json:"connection-timeout" yaml:"connection-timeout" toml:"connection-timeout" xml:"connection-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout" yaml:"write-timeout" toml:"write-timeout" xml:"write-timeout"`
	ReadTimeout       time.Duration `json:"read-timeout" yaml:"read-timeout" toml:"read-timeout" xml:"read-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout" yaml:"idle-timeout" toml:"idle-timeout" xml:"idle-timeout"`
	Wait              bool          `json:"wait" yaml:"wait" toml:"wait" xml:"wait"`
}

type LevelDBConfig struct {
	LocalPath string `json:"path" yaml:"path" toml:"path" xml:"path"`
}

type PathPlaceholder struct {
	User string `json:"user" yaml:"user" toml:"user" xml:"user"`
}

//

type BaseParametersConfig struct {
	PathPlaceholder     PathPlaceholder                `json:"path-placeholder" yaml:"path-placeholder" toml:"path-placeholder" xml:"path-placeholder"`
	ExposeHost          string                         `json:"expose-host" yaml:"expose-host" toml:"expose-host" xml:"expose-host"`
	NSBSignerPrivateKey string                         `json:"signer-key" yaml:"signer-key" toml:"signer-key" xml:"signer-key"`
	NSBSignerChainID    uiptypes.ChainIDUnderlyingType `json:"signer-chain-id" yaml:"signer-chain-id" toml:"signer-chain-id" xml:"signer-chain-id"`
	NSBHost             string                         `json:"nsb-host" yaml:"nsb-host" toml:"nsb-host" xml:"nsb-host"`
}

type Label struct {
	Key   string `json:"key" yaml:"key" toml:"key" xml:"key"`
	Value string `json:"value" yaml:"value" toml:"value" xml:"value"`
}

type DatabaseConfig = core_cfg.DatabaseConfig

type ServerConfig struct {
	LoadType             string               `json:"-" yaml:"-" toml:"-" xml:"-"`
	Name                 xml.Name             `json:"-" yaml:"-" toml:"-" xml:"server-config"`
	Labels               []Label              `json:"label" yaml:"label" toml:"label" xml:"label"`
	DatabaseConfig       DatabaseConfig       `json:"database" yaml:"database" toml:"database" xml:"database"`
	BaseParametersConfig BaseParametersConfig `json:"base-cfg" yaml:"base-cfg" toml:"base-cfg" xml:"base-cfg"`
	RedisConfig          RedisConfig          `json:"redis" yaml:"redis" toml:"redis" xml:"redis"`
	LevelDBConfig        LevelDBConfig        `json:"leveldb" yaml:"leveldb" toml:"leveldb" xml:"leveldb"`
}

func (s ServerConfig) GetDatabaseConfiguration() core_cfg.DatabaseConfig {
	return s.DatabaseConfig
}

func Default() *ServerConfig {
	return &ServerConfig{
		LoadType: "json",
		BaseParametersConfig: BaseParametersConfig{
			PathPlaceholder: PathPlaceholder{
				User: "id",
			},
			ExposeHost:          "127.0.0.1:23452",
			NSBHost:             "39.100.145.91:26657",
			NSBSignerPrivateKey: "2333bfffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff2333bfffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff",
			NSBSignerChainID:    3,
		},
		LevelDBConfig: LevelDBConfig{
			LocalPath: "./level-test",
		},
	}
}

func Load(config *ServerConfig, configPath string) error {
	return LoadStatic(config, configPath)
}
func unmarshal(config interface{}, unmarshaler func(b []byte, i interface{}) error,
	configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(f)
	_ = f.Close()
	if err != nil {
		return err
	}
	err = unmarshaler(b, config)
	if err != nil {
		return err
	}
	return nil
}

func LoadStatic(config interface{}, configPath string) error {

	for _, configX := range []struct {
		Type      string
		Unmarshal func([]byte, interface{}) error
	}{
		{".json", json.Unmarshal}, {".yml", yaml.Unmarshal},
		{".toml", toml.Unmarshal}, {".xml", xml.Unmarshal}} {

		if strings.HasSuffix(configPath, configX.Type) {
			return unmarshal(config, configX.Unmarshal, configPath)
		}

		if _, err := os.Stat(configPath + configX.Type); err == nil {
			return unmarshal(config, configX.Unmarshal, configPath+configX.Type)
		}
	}

	return errors.New("no such file in the root directory")
}

func Save(config *ServerConfig, configPath string) error {
	var b []byte
	var err error
	switch config.LoadType {
	case ".json":
		b, err = json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
	case ".yml":
		b, err = yaml.Marshal(config)
		if err != nil {
			return err
		}
	case ".toml":
		b, err = toml.Marshal(config)
		if err != nil {
			return err
		}
	case ".xml":
		b, err = xml.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(configPath + config.LoadType); err == nil {
		f, err := os.OpenFile(configPath+config.LoadType, os.O_WRONLY|os.O_TRUNC, 0333)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		_ = f.Close()
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("no such file in the root directory")
}
