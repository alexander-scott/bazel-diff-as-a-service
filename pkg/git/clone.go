// Package git is used to perform git related operations
package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

// CloneRepo shall clone the specified repo in the specified path
func CloneRepo(cloneURL string, cloneDest string) error {
	_, err := git.PlainClone(cloneDest, false, &git.CloneOptions{
		URL:      cloneURL,
		Progress: os.Stdout,
	})
	return err
}
