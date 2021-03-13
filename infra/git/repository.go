package git

import (
	"os/exec"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Repository struct {
	g      *git.Repository
	remote *git.Remote
	dir    string
}

// RenameBranch Change the default branch on the current Repository to use "main"
func (r *Repository) RenameBranch() error {
	cmd := exec.Command("git", "branch", "-M", "main")
	cmd.Dir = r.dir
	return cmd.Run()
}

// Commit mark a commit on the current Repository
func (r *Repository) Commit(msg, addPattern, author, email string) (err error) {
	var wt *git.Worktree
	if wt, err = r.g.Worktree(); err != nil {
		return err
	}

	if _, err = wt.Add(addPattern); err != nil {
		return err
	}

	_, err = wt.Commit(msg, &git.CommitOptions{All: true, Author: &object.Signature{
		Name:  author,
		Email: email,
		When:  time.Now(),
	}})
	return err
}

// SetRemote defines a new remote for the Repository
func (r *Repository) SetRemote(remoteAddr string) (err error) {
	var remote *git.Remote
	if remote, err = r.g.CreateRemote(&config.RemoteConfig{URLs: []string{remoteAddr}}); err != nil {
		return
	}
	r.remote = remote
	return
}
