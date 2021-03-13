package cmd

import (
	"fmt"

	"github.com/caioeverest/vulcan/infra/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	configCMD.AddCommand(templateCMD)

	templateCMD.Flags().BoolVarP(&list, "list", "l", false, "list all stored templates")
	templateCMD.Flags().BoolVarP(&delTemplate, "delete", "d", false, "delete a template")
	templateCMD.Flags().StringVarP(&templatename, "name", "n", "", "--name Go")
	templateCMD.Flags().StringVarP(&templateaddr, "address", "a", "", "--address git@github.com:something/template.git")
}

var (
	templatename, templateaddr string
	delTemplate                bool
)

var templateCMD = &cobra.Command{
	Use:   "template",
	Short: "menage config template map",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			conf      = config.Get()
			didAddAll bool
		)

		if !list {
			if templatename == "" || templateaddr == "" {
				for !didAddAll {
					templatename = c.Prompt.StringRequired("Name")
					templateaddr = c.Prompt.StringRequired("Address")
					if c.Prompt.Confirm(fmt.Sprintf("Did you want to add the template %s with the address %s", templatename, templateaddr)) {
						conf.TemplateMap[templatename] = templateaddr
					}
					didAddAll = !c.Prompt.Confirm("Do you want to add another template?")
				}
				return save(conf)
			}

			conf.TemplateMap[templatename] = templateaddr
			return save(conf)
		}

		if delTemplate {
			if templatename == "" {
				return errors.New("the flag \"Name\" must be settled to delete a template")
			}
			delete(conf.TemplateMap, templatename)
			return save(conf)
		}

		c.Successf("Stored templates: %s", mapToString(conf.TemplateMap))
		return nil
	},
}
