package config

import (
	"log"
	"os"
	"reflect"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Name           string                 `toml:"name"`
	Email          string                 `toml:"email"`
	SSHPubKey      string                 `toml:"ssh_pub_key"`
	TemplateMap    map[string]string      `toml:"template-map"`
	PlaceholderMap map[string]interface{} `toml:"placeholder-map"`
}

const (
	perm                  = 0644
	altConfigPathVarName  = "vulcan_CFG"
	defaultConfigFileName = ".vulcanrc"
)

var (
	_global    *Config
	configPath = os.Getenv("HOME") + "/" + defaultConfigFileName
	Version    string
)

func Get() *Config {
	if _global == nil {
		_ = Open()
	}
	return _global
}

func Open() error {
	var (
		err  error
		file *os.File
	)

	_global = &Config{TemplateMap: make(map[string]string), PlaceholderMap: make(map[string]interface{})}

	if alternativePath := os.Getenv(altConfigPathVarName); alternativePath != "" {
		configPath = alternativePath
	}

	if file, err = os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, perm); err != nil {
		return err
	}
	defer file.Close()

	if err = toml.NewDecoder(file).Decode(_global); err != nil {
		return err
	}
	return nil
}

func (c *Config) IsEmpty() bool {
	return reflect.ValueOf(c.Name).IsZero()
}

func (c *Config) Save() error {
	var (
		err  error
		file *os.File
	)

	if file, err = os.OpenFile(configPath, os.O_WRONLY, perm); err != nil {
		return err
	}
	defer file.Close()

	if err = toml.NewEncoder(file).Encode(*c); err != nil {
		return err
	}
	return nil
}

func (c *Config) String() string {
	out, err := toml.Marshal(*c)
	if err != nil {
		log.Panicf("Unreadable config: %+v", err)
	}
	return string(out)
}
