package cmd

import (
	"encoding/json"
	"os"

	"github.com/caioeverest/vulcan/infra/config"
	"github.com/caioeverest/vulcan/pkg/builder"
	"github.com/caioeverest/vulcan/pkg/project"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCMD)
	createCMD.Flags().StringVarP(&templateAddress, "custom-template", "c", "", "the path of your custom template folder")
	createCMD.Flags().StringVarP(&customPlaceholdersPath, "custom-placeholders", "p", "", "the path of your custom placeholders file [must be on json format]")
}

var (
	templateAddress        string
	customPlaceholdersPath string

	createCMD = &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				conf               = config.Get()
				customPlaceholders = conf.PlaceholderMap
				remote             string
			)

			if templateAddress == "" && len(conf.TemplateMap) == 0 {
				return errors.New("You need to set a template on your configs or use the --custom-template tag to define a template")
			}

			if customPlaceholdersPath != "" {
				customPlaceholders = readCustomMap(conf)
			}

			c.Info("Ready to project a new project?!")

			projectName := c.Prompt.String("Tel me, what is the name of your project", "Something Awesome")
			description := c.Prompt.String("And what did will do?", "Something, definably")

			if c.Prompt.Confirm("Did you create a remote repository already?") {
				remote = c.Prompt.StringRequired("The remote address of your project:")
			}

			if templateAddress == "" {
				lang := c.Prompt.Choose("It will based on which template?", getAvailableTemplates(conf)...)
				templateAddress = conf.TemplateMap[lang]
			}

			p := project.New(projectName, description, conf.Name, conf.Email, project.WithCustomMap(customPlaceholders))

			b := builder.Get(c, p, templateAddress, remote)
			if err := b.Prepare(); err != nil {
				return err
			}
			if err := b.Build(); err != nil {
				return err
			}
			if err := b.Clean(); err != nil {
				return err
			}

			c.Successf("Project %s created with success! Happy Coding :)", p.ProjectName)
			return nil
		},
	}
)

func getAvailableTemplates(conf *config.Config) (a []string) {
	for k := range conf.TemplateMap {
		a = append(a, k)
	}
	return
}

func readCustomMap(conf *config.Config) map[string]interface{} {
	var (
		result       = make(map[string]interface{})
		m            = make(map[string]interface{})
		content, err = os.ReadFile(customPlaceholdersPath)
	)

	if err != nil {
		c.Warnf("Unable to find custom placeholders file on path %s - error %+v", customPlaceholdersPath, err)
	}

	if err = json.Unmarshal(content, &m); err != nil {
		c.Warnf("Error reading custom placeholders file %+v", err)
	}

	for k, v := range conf.PlaceholderMap {
		result[k] = v
	}

	for k, v := range m {
		result[k] = v
	}

	return result
}
