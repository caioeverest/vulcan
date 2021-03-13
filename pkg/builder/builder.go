package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	gotemplate "text/template"

	"github.com/caioeverest/vulcan/infra/console"
	"github.com/caioeverest/vulcan/infra/git"
	"github.com/caioeverest/vulcan/pkg/project"
	"github.com/caioeverest/vulcan/pkg/template"
	"github.com/pkg/errors"
)

type Builder interface {
	Prepare() error
	Build() error
	Clean() error
	AbsolutePath() string
	Project() interface{}
	Templates() []template.Template
}

func Get(con console.Console, p *project.Project, templateAddress, remote string) Builder {
	if isGitTemplate(templateAddress) {
		return &Git{project: *p, templateAddress: templateAddress, remote: remote, c: con}
	}
	return &Fsys{project: *p, templateAddress: templateAddress, remote: remote, c: con}
}

func isGitTemplate(templateAddress string) bool {
	r, _ := regexp.Match("^git@git*", []byte(templateAddress))
	return r
}

func generateFromTemplates(b Builder, c console.Console) error {
	for _, t := range b.Templates() {
		if err := formatter(c, b, t); err != nil {
			return err
		}
	}
	return nil
}

func formatter(con console.Console, b Builder, t template.Template) (err error) {
	var tmpl = gotemplate.New(t.OriginalFileName)

	tmpl = tmpl.Funcs(gotemplate.FuncMap{"unescaped": unescaped})
	if tmpl, err = tmpl.ParseFiles(t.OriginalFilePath); err != nil {
		return errors.WithStack(err)
	}

	relateDir := filepath.Dir(t.NewFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(t.NewFilePath))
	distAbsFilePath := filepath.Join(b.AbsolutePath(), distRelFilePath)

	con.Debugf("distRelFilePath: %s", distRelFilePath)
	con.Debugf("distAbsFilePath: %s", distAbsFilePath)

	if err = os.MkdirAll(filepath.Dir(distAbsFilePath), os.ModePerm); err != nil {
		return errors.WithStack(err)
	}

	dist, err := os.Create(distAbsFilePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer dist.Close()

	fmt.Printf("Create %s\n", distRelFilePath)
	return tmpl.Execute(dist, b.Project())
}

func unescaped(x string) interface{} { return gotemplate.New(x) }

func initialCommit(con console.Console, p project.Project, remote string) (err error) {
	var r git.Repository

	if r, err = git.Initiate(p.RepositoryName); err != nil {
		return
	}
	con.Success("Repository initiated")

	if err = r.Commit("Initial Commit", ".", p.Author, p.Email); err != nil {
		return
	}

	if err = r.RenameBranch(); err != nil {
		return
	}

	if remote != "" {
		if err = r.SetRemote(remote); err != nil {
			return
		}
	}

	con.Success("First Commit made")
	return
}
