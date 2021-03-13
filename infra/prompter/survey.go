package prompter

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

// CliPrompter used to prompt for cli input
type CliPrompter struct{}

// ChooseWithDefault given the choice return the option selected with a default
func (cli *CliPrompter) ChooseWithDefault(msg string, defaultValue string, options ...string) (string, error) {
	selected := ""
	prompt := &survey.Select{
		Message: msg,
		Options: options,
		Default: defaultValue,
	}
	_ = survey.AskOne(prompt, &selected, survey.WithValidator(survey.Required))

	// return the selected element index
	for i, option := range options {
		if selected == option {
			return options[i], nil
		}
	}
	return "", errors.New("bad input")
}

// Choose given the choice return the option selected
func (cli *CliPrompter) Choose(msg string, options ...string) string {
	selected := ""
	prompt := &survey.Select{
		Message: msg,
		Options: options,
	}
	_ = survey.AskOne(prompt, &selected, survey.WithValidator(survey.Required))

	// return the selected element index
	for _, option := range options {
		if selected == option {
			return option
		}
	}
	return ""
}

// Confirm Ask the user if they confirm
func (cli *CliPrompter) Confirm(msg string) bool {
	confirm := false
	prompt := &survey.Confirm{
		Message: msg,
	}
	_ = survey.AskOne(prompt, &confirm, survey.WithValidator(survey.Required))
	return confirm
}

// StringRequired prompt for string which is required
func (cli *CliPrompter) String(msg string, defaultValue string) string {
	val := ""
	prompt := &survey.Input{
		Message: msg,
		Default: defaultValue,
	}
	_ = survey.AskOne(prompt, &val)
	return val
}

// StringRequired prompt for string which is required
func (cli *CliPrompter) StringRequired(msg string) string {
	val := ""
	prompt := &survey.Input{
		Message: msg,
	}
	_ = survey.AskOne(prompt, &val, survey.WithValidator(survey.Required))
	return val
}

// Password prompt for password which is required
func (cli *CliPrompter) Password(msg string) string {
	val := ""
	prompt := &survey.Password{
		Message: msg,
	}
	_ = survey.AskOne(prompt, &val)
	return val
}
