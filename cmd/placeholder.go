package cmd

import (
	"fmt"

	"github.com/caioeverest/vulcan/infra/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	configCMD.AddCommand(placeholderCMD)

	placeholderCMD.Flags().BoolVarP(&list, "list", "l", false, "list all stored placeholders")
	placeholderCMD.Flags().BoolVarP(&delPlaceholder, "delete", "d", false, "delete a placeholder")
	placeholderCMD.Flags().StringVarP(&placeholdername, "name", "n", "", "--name Team")
	placeholderCMD.Flags().StringVarP(&placeholdervalue, "value", "v", "", "--value @company/awasome-squad")
}

var (
	placeholdername, placeholdervalue string
	delPlaceholder                    bool
)

var placeholderCMD = &cobra.Command{
	Use:   "placeholder",
	Short: "menage custom global placeholder map",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			conf      = config.Get()
			didAddAll bool
		)

		if !list {
			if placeholdername == "" || placeholdervalue == "" {
				for !didAddAll {
					placeholdername = c.Prompt.StringRequired("Name")
					placeholdervalue = c.Prompt.StringRequired("Value")
					if c.Prompt.Confirm(fmt.Sprintf("Did you want to add the placeholder %s with the value %s to the global map?", placeholdername, placeholdervalue)) {
						conf.PlaceholderMap[placeholdername] = placeholdervalue
					}
					didAddAll = !c.Prompt.Confirm("Do you want to add another template?")
				}
				return save(conf)
			}

			conf.PlaceholderMap[placeholdername] = placeholdervalue
			return save(conf)
		}

		if delPlaceholder {
			if placeholdername == "" {
				return errors.New("the flag \"Name\" must be settled to delete a placeholder")
			}
			delete(conf.PlaceholderMap, placeholdername)
			return save(conf)
		}

		c.Successf("Stored placeholders: %s", mapToString(conf.PlaceholderMap))
		return nil
	},
}
