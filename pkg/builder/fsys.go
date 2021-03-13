package builder

import (
	"os"
	"path/filepath"

	"github.com/caioeverest/vulcan/infra/console"
	"github.com/caioeverest/vulcan/pkg/project"
	"github.com/caioeverest/vulcan/pkg/template"
)

type Fsys struct {
	c               console.Console
	project         project.Project
	templates       []template.Template
	templateAddress string
	absolutePath    string
	remote          string
}

func (b *Fsys) AbsolutePath() string           { return b.absolutePath }
func (b *Fsys) Project() interface{}           { return b.project }
func (b *Fsys) Templates() []template.Template { return b.templates }

func (b *Fsys) Prepare() (err error) {
	b.c.Info("Creating folder")
	if err = os.Mkdir(b.project.RepositoryName, outputPerm); err != nil {
		return
	}
	b.c.Successf("Folder %s/, created!", b.project.RepositoryName)

	if b.absolutePath, err = filepath.Abs(b.project.RepositoryName); err != nil {
		return
	}

	b.templates = template.GetTemplates(b.c, b.project, b.templateAddress)
	return
}

func (b *Fsys) Build() (err error) {
	if err = generateFromTemplates(b, b.c); err != nil {
		return err
	}
	return initialCommit(b.c, b.project, b.remote)
}

func (b *Fsys) Clean() (err error) {
	return
}
