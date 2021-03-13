package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caioeverest/vulcan/infra/console"
	"github.com/caioeverest/vulcan/infra/git"
	"github.com/caioeverest/vulcan/pkg/project"
	"github.com/caioeverest/vulcan/pkg/template"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/google/uuid"
)

type Git struct {
	c               console.Console
	project         project.Project
	absolutePath    string
	templatePath    string
	templateAddress string
	remote          string
	templates       []template.Template
}

const (
	outputPerm = 0755
	tmpFolder  = "/tmp/vulcan-%s"
)

func (b *Git) AbsolutePath() string           { return b.absolutePath }
func (b *Git) Project() interface{}           { return b.project }
func (b *Git) Templates() []template.Template { return b.templates }

func (b *Git) Prepare() (err error) {
	//Create project folder
	b.c.Info("Creating folder")
	if err = os.Mkdir(b.project.RepositoryName, outputPerm); err != nil {
		return
	}
	b.c.Successf("Folder %s/, created!", b.project.RepositoryName)

	//Get absolute path
	if b.absolutePath, err = filepath.Abs(b.project.RepositoryName); err != nil {
		return
	}

	b.templatePath = genTmplFolderPath()
	b.c.Info("Creating temporary folder")
	if err = os.Mkdir(b.templatePath, outputPerm); err != nil {
		return
	}
	b.c.Successf("Folder %s/, created!", b.templatePath)

	b.c.Info("Downloading base template")
	if err = git.Clone(b.templatePath+"/", b.buildCloneOptions()); err != nil {
		return
	}

	b.templates = template.GetTemplates(b.c, b.project, b.templatePath)
	return
}

func (b *Git) Build() (err error) {
	if err = generateFromTemplates(b, b.c); err != nil {
		return err
	}
	return initialCommit(b.c, b.project, b.remote)
}

func (b *Git) Clean() (err error) {
	if err = os.RemoveAll(b.templatePath); err != nil {
		return
	}
	b.c.Success("Downloaded finished!")
	return
}

func (b *Git) buildCloneOptions() git.CloneOptions {
	addressSlice := strings.Split(b.templateAddress, ";")

	if len(addressSlice) > 1 {
		return git.CloneOptions{
			URL:           addressSlice[0],
			Progress:      os.Stdout,
			ReferenceName: plumbing.NewBranchReferenceName(addressSlice[1]),
		}
	}

	return git.CloneOptions{URL: b.templateAddress, Progress: os.Stdout}
}

func genTmplFolderPath() string {
	return fmt.Sprintf(tmpFolder, uuid.New().String())
}
