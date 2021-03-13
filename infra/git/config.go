package git

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	User struct {
		Name  string `ini:"name"`
		Email string `ini:"email"`
	} `ini:"user"`
}

func OpenGitConfig() (c Config, err error) {
	var (
		content    []byte
		configPath = fmt.Sprintf("%s/.gitconfig", os.Getenv("HOME"))
	)

	if content, err = os.ReadFile(configPath); err != nil {
		return
	}

	if err = ini.MapTo(&c, content); err != nil {
		return
	}

	return c, nil
}
