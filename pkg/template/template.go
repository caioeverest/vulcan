package template

import (
	"path/filepath"

	"github.com/caioeverest/vulcan/infra/console"
	"github.com/caioeverest/vulcan/pkg/project"
)

type Template struct {
	OriginalFilePath string
	OriginalFileName string
	NewFilePath      string
}

func GetTemplates(con console.Console, project project.Project, templatesFolder string) []Template {
	engine := Engine{project: project}
	con.Debugf("walk: [%s]", templatesFolder)
	if err := filepath.Walk(templatesFolder, engine.process(templatesFolder)); err != nil {
		con.Fatalf("Error looking for folders %+v", err)
	}
	return engine.Templates
}
