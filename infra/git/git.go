package git

import (
	"github.com/go-git/go-git/v5"
)

// CloneOptions are used when cloning a git Repository
type CloneOptions git.CloneOptions

// Clone clones a git Repository with the given options
func Clone(dir string, opts CloneOptions) error {
	o := git.CloneOptions(opts)

	_, err := git.PlainClone(dir, false, &o)
	return err
}

// Initiate starts a new git Repository with the given options
func Initiate(dir string) (Repository, error) {
	repo, err := git.PlainInit(dir, false)
	return Repository{g: repo, dir: dir}, err
}
