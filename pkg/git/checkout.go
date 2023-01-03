package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// Checkout a git repository to a specific commit
func Checkout(r *git.Repository, commit string) error {
	w, err := r.Worktree()
	if err != nil {
		fmt.Println(err)
	}

	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	})
	return err
}
