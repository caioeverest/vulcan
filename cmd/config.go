package cmd

import (
	"fmt"
	"os"
	"reflect"

	"github.com/caioeverest/vulcan/infra/git"
	"github.com/caioeverest/vulcan/infra/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCMD)

	configCMD.Flags().BoolVarP(&list, "list", "l", false, "list configs")
	configCMD.Flags().StringVar(&newName, "name", "", "--name Jonh Smith")
	configCMD.Flags().StringVar(&newEmail, "email", "", "--email jonh.smith@something.com")
	configCMD.Flags().StringVar(&newSSHPubKey, "ssh", "", "--ssh /your/favorite/path/to/file")
}

var (
	newName, newEmail, newSSHPubKey string
	list                            bool
)

var configCMD = &cobra.Command{
	Use:   "config",
	Short: "Config user",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			conf = config.Get()
		)

		if !list {
			setConfig(conf)
			return save(conf)
		}

		c.Success(conf.String())
		return nil
	},
}

func setConfig(conf *config.Config) {
	if shouldUseFlags() {
		if newName != "" {
			conf.Name = newName
		}

		if newEmail != "" {
			conf.Email = newEmail
		}

		if newSSHPubKey != "" {
			conf.SSHPubKey = newSSHPubKey
		}
	} else {
		email, name := buildEmailAndNameHelper()
		if conf.Name == "" {
			conf.Name = c.Prompt.String("Name", name)
		}

		if conf.Email == "" {
			conf.Email = c.Prompt.String("Email", email)
		}

		if conf.SSHPubKey == "" {
			conf.SSHPubKey = c.Prompt.String("SSH key path", os.Getenv("HOME")+"/.ssh/id_rsa.pub")
		}
	}
}

func shouldUseFlags() bool {
	return newName != "" || newEmail != "" || newSSHPubKey != ""
}

func save(cfg *config.Config) error {
	if err := cfg.Save(); err != nil {
		return err
	}
	c.Success("Configurations saved with success!")
	return nil
}

func buildEmailAndNameHelper() (email, name string) {
	var (
		err    error
		gitCfg git.Config
	)

	if gitCfg, err = git.OpenGitConfig(); err != nil {
		name = os.Getenv("USER")
		email = fmt.Sprintf("%s@gmail.com", name)
		return
	}

	name = gitCfg.User.Name
	email = gitCfg.User.Email

	if name == "" {
		name = os.Getenv("USER")
	}

	if email == "" {
		email = fmt.Sprintf("%s@gmail.com", name)
	}
	return
}

func mapToString(m interface{}) string {
	var (
		mv     = reflect.ValueOf(m)
		kv     = mv.MapRange()
		result = "["
	)

	for kv.Next() {
		result += fmt.Sprintf("\n %s: %v", kv.Key(), kv.Value())
	}
	result += "\n]"
	return result
}
