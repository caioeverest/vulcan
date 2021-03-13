package project

import "strings"

type Project struct {
	RepositoryName string
	ProjectName    string
	Description    string
	Author         string
	Email          string
	Custom         map[string]interface{}
}

func New(name, description, author, email string, opts ...Option) *Project {
	p := &Project{
		RepositoryName: strings.ReplaceAll(strings.ToLower(name), " ", "-"),
		ProjectName:    name,
		Description:    description,
		Author:         author,
		Email:          email,
		Custom:         make(map[string]interface{}),
	}

	for _, o := range opts {
		o(p)
	}

	return p
}
