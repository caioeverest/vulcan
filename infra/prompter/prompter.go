package prompter

// Mock handles prompting user for input
type Prompter interface {
	ChooseWithDefault(string, string, ...string) (string, error)
	Choose(string, ...string) string
	Confirm(string) bool
	StringRequired(string) string
	String(string, string) string
	Password(string) string
}
