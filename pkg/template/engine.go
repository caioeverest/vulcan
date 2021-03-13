package template

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	gotemplate "text/template"

	"github.com/caioeverest/vulcan/pkg/project"
	"github.com/pkg/errors"
)

type Engine struct {
	project    project.Project
	Templates  []Template
	projectDir string
}

const goTemplateSuffix = ".tgo"

func (t *Engine) process(referencePath string) func(path string, f os.FileInfo, err error) error {
	var tmpl Template

	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if b, _ := regexp.Match("/*/.git*/", []byte(path)); b {
			return nil
		}

		if ext := filepath.Ext(path); ext == goTemplateSuffix {
			if tmpl, err = t.processGoTemplateFiles(referencePath, path); err != nil {
				return err
			}
		} else if mode := f.Mode(); mode.IsRegular() {
			if tmpl, err = t.processStandardFiles(referencePath, path); err != nil {
				return err
			}
		}

		if reflect.ValueOf(tmpl).IsZero() {
			return nil
		}

		t.Templates = append(t.Templates, tmpl)
		return nil
	}
}

func (t *Engine) processGoTemplateFiles(referencePath, path string) (Template, error) {
	var (
		templateFileName = filepath.Base(path)
		genFileBaseName  = strings.TrimSuffix(templateFileName, goTemplateSuffix) + ".go"
		genFileBasePath  string
		err              error
	)

	if genFileBasePath, err = filepath.Rel(referencePath, filepath.Join(filepath.Dir(path), genFileBaseName)); err != nil {
		return Template{}, errors.WithStack(err)
	}

	return Template{
		OriginalFilePath: path,
		OriginalFileName: templateFileName,
		NewFilePath:      t.generateNewFilePath(genFileBasePath),
	}, nil
}

func (t *Engine) processStandardFiles(basePath, path string) (Template, error) {
	var (
		templateFileName = filepath.Base(path)
		targetPath       = filepath.Join(filepath.Dir(path), templateFileName)
		genFileBasePath  string
		err              error
	)

	if genFileBasePath, err = filepath.Rel(basePath, targetPath); err != nil {
		return Template{}, errors.WithStack(err)
	}

	return Template{
		OriginalFilePath: path,
		OriginalFileName: templateFileName,
		NewFilePath:      t.generateNewFilePath(genFileBasePath),
	}, nil
}

func (t *Engine) generateNewFilePath(basePath string) string {
	buf := bytes.NewBufferString("")
	formatted, _ := gotemplate.New("name").Parse(basePath)
	_ = formatted.Execute(buf, t.project)
	stringBuf := buf.String()
	if stringBuf != "" {
		basePath = stringBuf
	}
	return filepath.Join(t.projectDir, basePath)
}
